# Go(Echo) Go Modules × Dockerで開発環境構築

### 環境
Go 1.15 .2
Echo

### 一応、GO言語をインストールすること
https://golang.org/dl/

### 使用方法
ダウンロードして`docker-compose up -d` で開発環境構築完了

### DB接続情報はconf/config.goを確認
### Usersテーブルを作成する
### 参考 https://nomad.office-aship.info/gorm-echo-mysql-json/
create table users (
	id int,
	email varchar(255),
	name varchar(255)
);
insert into users values(1,'test@gmail.com', 'test');

### 起動確認

### ホーム画面
### localhost:5555

### ユーザ一覧画面
### http://localhost:5555/users

### ユーザー詳細画面
### http://localhost:5555/users/1

### JWT TOKEN認証
###　token取得
### jwt token 参考 https://qiita.com/unvavo/items/b344a3ded2df8fa65c58
### curl -X POST -d 'username=test' -d 'password=test' localhost:5555/login

#### token認証確認
#### curl -H "Authorization: Bearer トークン" localhost:5555/restricted/welcome
