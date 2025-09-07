# このアプリについて

結婚したので、両親の顔合わせ用に、しおりアプリを作りました。

ネイティブ側はパートナーが作りました。

# 利用方法
### バージョン情報
- go v1.24.3
- gin v1.10.1
- MySQL v8.0


### サーバの起動方法（ローカル）
1. GitHubからリポジトリをクローンします。
2. お好みの方法で.envファイルを作成します。
``` shell
cd ./cmd/server
nano .env
```
（設定項目）
``` 
DB_USER=your_db_username
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306 
DB_NAME=your_db_name
S3_BUCKET_NAME=your_bucket_name
AWS_REGION=your_aws_region
API_TOKEN=your_authorizing_token
```

3. サーバ起動
./cmd/serverディレクトリで、以下のコマンドを実行します。
``` shell
go run main.go
```

必要に応じて、プロフィールや画像などのデータを登録してください。