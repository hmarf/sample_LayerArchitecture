# sample_LayerArchitecture

## 概要
Golangで書いたレイヤードアーキテクチャのサンプルコードです。  
簡単なユーザー登録、情報取得、削除機能のついたAPIサーバーです。  
Qiita: https://qiita.com/hmarf/items/7f4d39c48775c205b99b  

## How to use
DBはDockerの乗せてあるので、以下のコマンドで起動します。  
(1) docker-compose up -d  
go runすれば動くはずです。  
(2) go run cmd/main.go

## URL
(1) localhost:9000/create/user  
新しいユーザーを作成します。   
[Request]登録したいユーザーの名前をPostします。  
{  
　　"name":"userName"  
}  
[Response]登録したユーザーのIDと名前、アカウントの作成時間をResponse  
{  
　　"userId": "e3dcfbec-3b1d-4aae-927b-6a1d22737d7b",  
　　"name": "userName",  
　　"createdAt": "2019-09-05 01:30:50.477038 +0900 JST m=+7.483626405"  
}  

(2) localhost:9000/get/user  
ユーザー情報を取得します。
[Request]登録時に返ってきたUserIDをPostします。  

{  
　　"userId":"------------------------------"  
}  
[Response]登録したユーザーのIDと名前、アカウントの作成時間をResponse  
{  
　　"userId": "e3dcfbec-3b1d-4aae-927b-6a1d22737d7b",  
　　"name": "userName",  
　　"createdAt": "2019-09-05 01:30:50.477038 +0900 JST m=+7.483626405"  
}  

(3) localhost:9000/delete/user  
ユーザー情報を取得します。  
[Request]登録時に返ってきたUserIDをPostします。  

{  
　　"userId":"------------------------------"  
}  
[Response]削除が成功したかどうかResponse  
{  
　　"results":"success"  
}  
