# Node.jsイメージを使用
FROM node:16

# 作業ディレクトリを設定
WORKDIR /app

# package.jsonとlockファイルをコピー
COPY package.json ./

# 依存関係をインストール
RUN npm install

# アプリケーションコードをコピー
COPY . .

# 開発サーバーを起動
CMD ["npm", "run", "serve"]

# サーバーが利用するポートを指定
EXPOSE 8080
