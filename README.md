# bals4rayban
レイバンスパムに制裁を, タイムラインに平穏を  
## 使用ライブラリ
github.com/ChimeraCoder/anaconda  
## 使い方  
### 削除ツール
conf-example.jsonをconf.jsonにリネームする  
conf.jsonにConsumerKey, ConsumerSelect, AccessToken, AccessTokenSelectを設定する  
balsディレクトリに移動してgo run main.goと呪文を唱える

### 待機ツール
プログラムを待機させることによりレイバンのスパム感染を監視します.  
AccessTokenとかの設定は削除ツールと同じ.  
以下の項目を設定することでレイバンスパムの感染の通知及び投稿を削除します.  
* MailAddress
* MailPassword
* SMTPServer
* SMTPPort
