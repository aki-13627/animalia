# ---------------------------------------------------------------------------------  # 
#                               学習プロセス(実投稿データ)                               #
# ---------------------------------------------------------------------------------  #

# ライブラリのインポート
import os
import pandas as pd
import psycopg2
import json
import torch
from recommend_system.models.mmneumf import MultiModalNeuMFEngine
from recommend_system.utils.data import SampleGenerator
from dotenv import load_dotenv, find_dotenv

_ = load_dotenv(find_dotenv())

# ----------------------------------
# 1. PostgreSQLから実データを取得
# ----------------------------------
# PostgreSQLデータベースへの接続(接続確認済み)
def get_connection():
    return psycopg2.connect(
        user=os.getenv('DB_USER'),
        password=os.getenv('DB_PASSWORD'),
        database=os.getenv('DB_NAME'),
        host=os.getenv('DB_HOST'),
        port=os.getenv('DB_PORT')
    )
conn = get_connection()

# ratingsデータフレームを作成
query = """
        -- 投稿自体のインタラクション(投稿者による投稿)
        SELECT
            UserID AS user_id, PostID AS post_id, 1 AS rating, CreatedAt AS timestamp, 
            ImageFeature AS image_feature, TextFeature AS text_feature
        FROM Post
        WHERE EmbeddedFlg = TRUE

        UNION -- 縦結合＋重複削除

        -- 「いいね」のインタラクション
        SELECT
            L.UserID AS user_id, L.PostID AS post_id, 1 AS rating, L.CreatedAt AS timestamp,
            P.ImageFeature AS image_feature, P.TextFeature AS text_feature
        FROM "Like" L
        JOIN Post P ON L.PostID = P.PostID
        WHERE P.EmbeddedFlg = TRUE

        UNION

        -- コメントのインタラクション
        SELECT
            C.UserID AS user_id, C.PostID AS post_id, 1 AS rating, C.CreatedAt AS timestamp,
            P.ImageFeature AS image_feature, P.TextFeature AS text_feature
        FROM Comment C
        JOIN Post P ON C.PostID = P.PostID
        WHERE P.EmbeddedFlg = TRUE;
        """
prod_df = pd.read_sql(query, conn)
conn.close()

# ----------------------------------
# 2. 特徴量のパース(JSON -> List)
# ----------------------------------
def parse_feature(x):
    if isinstance(x, str):
        return json.loads(x)
    return x

prod_df["image_feature"] = prod_df["image_feature"].map(parse_feature)
prod_df["text_feature"] = prod_df["text_feature"].map(parse_feature)

# 埋め込みベクトルの次元数を取得
print(f"Image feature dimension: {prod_df['image_feature'].iloc[0].shape}")
print(f"Text feature dimension: {prod_df['text_feature'].iloc[0].shape}")

print(f"Production data loaded: {prod_df.shape[0]} records")

# ----------------------------------
# 3. プロダクション用の設定(config)
# ----------------------------------
prod_config = {
    "alias": "multimodal_neumf_prod",
    "num_epoch": 50,
    "batch_size": 512,
    "optimizer": "adam",
    "adam_lr": 1e-3,
    "num_users": int(prod_df["userId"].nunique()), # データから実際のユーザー数・アイテム数を取得
    "num_items": int(prod_df["itemId"].nunique()),
    "latent_dim_mf": 8,
    "latent_dim_mlp": 8,
    "num_negative": 4,
    "layers": [16, 64, 32, 16, 8],
    "l2_regularization": 0.0000001,
    "weight_init_gaussian": True,
    "use_cuda": True,
    "use_bachify_eval": False,
    "device_id": 0,
    "pretrain": False,
    "model_dir": "recommend_system/checkpoints/prod_checkpoints_1/{}_Epoch{}_HR{:.4f}_NDCG{:.4f}.model",
    "image_emb_dim": 16,
    "text_emb_dim": 16,
    "image_feature_dim": 1024,
    "text_feature_dim": 768
}

# ----------------------------------
# 4. サンプル生成器の作成と評価データの準備
# ----------------------------------
sample_generator = SampleGenerator(ratings=prod_df)
evaluate_data = sample_generator.evaluate_data

# ----------------------------------
# 5. Multi-Modal NeuMFモデルの作成と学習
# ----------------------------------
engine = MultiModalNeuMFEngine(config=prod_config)

# エポックごとに学習と評価を実行
for epoch in range(prod_config["num_epoch"]):
    print(f"Epoch {epoch}/{prod_config['num_epoch']}")
    print("-" * 80)
    train_loader = sample_generator.instance_a_train_loader(prod_config["num_negative"], prod_config["batch_size"])
    engine.train_an_epoch(train_loader, epoch_id=epoch)
    hit_ratio, ndcg = engine.evaluate(evaluate_data, epoch_id=epoch)
    engine.save(prod_config["alias"], epoch, hit_ratio, ndcg)
