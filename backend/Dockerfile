# ベースイメージ
FROM golang:1.23.2

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールファイルをコピー
COPY go.mod ./

# 依存関係をインストール
RUN go mod download

# アプリケーションコードをコピー
COPY . .

# ビルド
RUN go build -o main .

# アプリケーションを起動
CMD ["./main"]

# サーバーが利用するポートを指定
EXPOSE 8000
