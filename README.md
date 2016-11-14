[![wercker status](https://app.wercker.com/status/60547319ed47a9ef330a10bef25bc863/s/master "wercker status")](https://app.wercker.com/project/byKey/60547319ed47a9ef330a10bef25bc863)
# eventory

イベントまとめアプリ。自分にとって本当にほしい新着情報だけを見る。

## 概要
同じようなワードで複数の類似サイトで検索するという面倒を解消。  
興味のあるワードをあらかじめ設定し、それに対して必要な情報とそうじゃない情報を分別します。そして、また新しい情報が入ってくるまでは放置するだけです。  

![イメージ図](https://github.com/tikasan/eventory/blob/master/doc/eventory_plan.png?raw=true)

## 今後の計画
[issue]("https://github.com/tikasan/eventory/issues")に上げています。  
[現状の設計の完成後の予定](https://github.com/tikasan/eventory/issues/52)

## 使用技術

- Swift
- golang
- apache
- MySQL
- Realm
- Linux(Rasbery Pi2)
- DDNS

**今後使用予定**

- Nginx
- https
- ElasticSearch(?)

## 実装済み機能

- イベントデータ収集
- APIサーバー
- 並列処理
- DBマイグレーション
- ログ
- Makefile
- GolangTest

**今後実装予定**

- プッシュ通知
