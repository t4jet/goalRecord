# ベースイメージを指定（Golangの公式イメージを使用）
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /backend

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main

# アプリケーションのポートを公開
EXPOSE 8080

# コンテナ内でアプリケーションを実行
CMD ["./main"]
