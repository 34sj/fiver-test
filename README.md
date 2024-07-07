# Introduction

[Building a Basic REST API in Go using Fiber](https://tutorialedge.net/golang/basic-rest-api-go-fiber/) を参考にFiberに入門してみる  

# 元のTutorialと異なる点

- このTutorialの書き方とFiber V3の書き方で多少違うところがあるので読み替えながら実装する
- TutorialではRDBサービスの構築を省略するためにSQLiteを採用しているが、より実践に近い形で実装したいのでMySQL環境を構築してくれるdocker-compose.yamlを用意して、DBはそれを使う
- ORMには[Gen](https://gorm.io/ja_JP/gen/)を採用する予定　タイプセーフなのと記述量が少ないのがGood
- 追加で評価理由を保存するカラムをBookテーブルに用意する

# 参考

- [Building a Basic REST API in Go using Fiber](https://tutorialedge.net/golang/basic-rest-api-go-fiber/)
- [Fiber](https://github.com/gofiber/fiber)
- [GoのORM「Gen」を使ってみた](https://zenn.dev/mizuko_dev/articles/dbeb7c93883b57)
- [docker-composeでMySQLに初期テーブルを追加する方法](https://scrapbox.io/gyarasu/docker-compose%E3%81%A7MySQL%E3%81%AB%E5%88%9D%E6%9C%9F%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E3%82%92%E8%BF%BD%E5%8A%A0%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95)

# TODO

Tutorialが完了し次第ブログにまとめる Fiber・Gen・Docker・Docker Composeについて理解・整理したいわね
