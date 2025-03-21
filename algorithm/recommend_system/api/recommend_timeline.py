# ---------------------------------------------------------------------------------  # 
#                      レコメンドタイムラインを生成するオンライン推論API                  　 #
# ---------------------------------------------------------------------------------  #

# ライブラリのインポート
import os
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
import torch
import uvicorn
import psycopg2
import json
from recommend_system.models.mmneumf import MultiModalNeuMF
from dotenv import load_dotenv, find_dotenv

_ = load_dotenv(find_dotenv())

# PostgreSQLデータベースへの接続
def get_connection():
    return psycopg2.connect(
        user=os.getenv('DB_USER'),
        password=os.getenv('DB_PASSWORD'),
        database=os.getenv('DB_NAME'),
        host=os.getenv('DB_HOST'),
        port=os.getenv('DB_PORT')
    )

# ----------------------------------
# APIリクエストとレスポンスのデータモデル
# ----------------------------------
class TimelineRequest(BaseModel):
    user_id: int
    score_threshold: float = 0.5 # 予測スコアの閾値

class Post(BaseModel):
    id: int
    timestamp: str
    image_feature: list
    text_feature: list
    score: float

class TimelineResponse(BaseModel):
    posts: list[Post]

# ----------------------------------
# PostgreSQLから候補投稿を取得する関数
# ----------------------------------
def get_candidate_posts():
    conn = get_connection()
    cur = conn.cursor()

    # feature_computed = true で特徴量が抽出済みの投稿を対象とする
    query = """
            SELECT 
                ID AS post_id,
                CreatedAt AS timestamp,
                ImageFeature AS image_feature, -- JSON文字列
                TextFeature AS text_feature -- JSON文字列
            FROM Post
            WHERE EmbeddedFlg = TRUE
            """
    cur.execute(query)
    rows = cur.fetchall()

    # カラム名のリスト
    columns = [desc[0] for desc in cur.description]
    candidates = []
    for row in rows:
        candidate = dict(zip(columns, row))
        candidate["image_feature"] = json.loads(candidate["image_feature"]) if isinstance(candidate["image_feature"], str) else candidate["image_feature"]
        candidate["text_feature"] = json.loads(candidate["text_feature"]) if isinstance(candidate["text_feature"], str) else candidate["text_feature"]
        candidates.append(candidate)
    cur.close()
    conn.close()
    return candidates

# ----------------------------------
# 候補投稿からレコメンドタイムラインを生成する関数
# ----------------------------------
def get_recommended_timeline(user_id, candidates, model, device, score_threshold=0.5):
    recommended = []
    user_tensor = torch.tensor([user_id], dtype=torch.long).to(device)

    for candidate in candidates:
        item_tensor = torch.tensor([candidate["post_id"]], dtype=torch.long).to(device)
        image_features = torch.tensor([candidate["image_feature"]], dtype=torch.float).unsqueeze(0).to(device)
        text_features = torch.tensor([candidate["text_feature"]], dtype=torch.float).unsqueeze(0).to(device)

        with torch.no_grad():
            score = model(user_tensor, item_tensor, image_features, text_features)
        score = score.item()

        if score > score_threshold:
            candidate["score"] = score
            recommended.append(candidate)
    
    # タイムスタンプで降順ソート
    sorted_recommend = sorted(recommended, key=lambda x: x["score"], reverse=True)
    return sorted_recommend

# ----------------------------------
# FastAPIアプリの構築
# ----------------------------------
app = FastAPI()

# デバイス設定とモデルのロード
device = "cuda" if torch.cuda.is_available() else "cpu"
model = MultiModalNeuMF().to(device)
model.eval()

@app.post("/recommend/timeline", response_model=TimelineResponse)
def recommend_timeline(request: TimelineRequest):
    try:
        # PostgreSQLから候補投稿画像を取得
        candidates = get_candidate_posts()
        recommended = get_recommended_timeline(request.user_id, candidates, model, device, request.score_threshold)
        posts = [Post(
            id=rc["post_id"],
            timestamp=rc["timestamp"],
            image_feature=rc["image_feature"],
            text_feature=rc["text_feature"],
            score=rc["score"]
        ) for rc in recommended]
        return TimelineResponse(posts=posts)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
