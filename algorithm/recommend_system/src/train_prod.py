# ---------------------------------------------------------------------------------  # 
#                               学習プロセス(実投稿データ)                               #
# ---------------------------------------------------------------------------------  #

# ライブラリのインポート
import pandas as pd
import json
import subprocess
import requests
from recommend_system.models.mmneumf import MultiModalNeuMFEngine
from recommend_system.utils.data import SampleGenerator
from recommend_system.utils.database import get_connection

# ----------------------------------
# 1. PostgreSQLから実データを取得
# ----------------------------------
# PostgreSQLデータベースへの接続
conn = get_connection()

# ratingsデータフレームを作成
query = """
        -- 投稿自体のインタラクション(投稿者による投稿)
        SELECT
            UserID AS user_id, ID AS post_id, 1 AS rating, CreatedAt AS timestamp, 
            ImageFeature AS image_feature, TextFeature AS text_feature
        FROM Post
        WHERE TextFeature IS NOT NULL AND ImageFeature IS NOT NULL

        UNION -- 縦結合＋重複削除

        -- 「いいね」のインタラクション
        SELECT
            L.UserID AS user_id, L.PostID AS post_id, 1 AS rating, L.CreatedAt AS timestamp,
            P.ImageFeature AS image_feature, P.TextFeature AS text_feature
        FROM "Like" L
        JOIN Post P ON L.PostID = P.ID
        WHERE P.TextFeature IS NOT NULL AND P.ImageFeature IS NOT NULL

        UNION

        -- コメントのインタラクション
        SELECT
            C.UserID AS user_id, C.PostID AS post_id, 1 AS rating, C.CreatedAt AS timestamp,
            P.ImageFeature AS image_feature, P.TextFeature AS text_feature
        FROM Comment C
        JOIN Post P ON C.PostID = P.ID
        WHERE P.TextFeature IS NOT NULL AND P.ImageFeature IS NOT NULL;
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
    "alias": "prod",
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
    "pretrain": True,
    "model_dir": "recommend_system/models/checkpoints/prod_{}_HR{:.4f}_NDCG{:.4f}.model",
    "pretrain_model_dir": "recommend_system/models/latest.model",
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
    engine.save_prod(hit_ratio, ndcg)

# ----------------------------------
# 6. 最新モデルのアップロードとモデルのリロード
# ----------------------------------
# upload_model.pyを実行
subprocess.run(["python", "recommend_system/src/upload_model.py"], check=True)

# 推論APIの /reload を叩く
try:
    response = requests.post("http://localhost:8000/reload")
    if response.status_code == 200:
        print("Model reloaded successfully", response.json())
    else:
        print("Failed to reload model", response.json())
except Exception as e:
    print("Failed to reload model", str(e))

