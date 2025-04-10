FROM golang:1.23-alpine

# 必要なパッケージをインストール
RUN apk add --no-cache git

# ワーキングディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# 依存関係をダウンロード
RUN go mod download

# 本番用にビルド（CGOを無効にし、Linux用にビルド）
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

EXPOSE 8080

# 実行
CMD ["/app/main"]
