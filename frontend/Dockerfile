# ベースイメージを指定（Node.jsの公式イメージを使用）
FROM node:latest

# 作業ディレクトリを設定
WORKDIR /frontend

# 依存関係をインストール
COPY package.json package-lock.json ./
RUN npm install

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN npm run build

# アプリケーションを起動
CMD ["npm", "start"]
