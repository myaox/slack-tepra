# Slackbot for Tepra

## 概要
slackのチャンネルでコマンドを打つことで気軽にテプラのラベルを印刷できます。  
例えばbotのいるチャンネルで以下のようにコメントをすると、

```
tepra print Hello World
```

```
Hello World
```
というラベルが印刷されます。


### QRコードを印刷

```
tepra qrcode https://google.com
```

## オプション

```-t ```  ラベルを印刷せずにテスト画像をslack上に表示  
```-n=10``` 同じラベルを複数回印刷  
```tepra print --qr=https://google.com hello world```  文字の右側にそのurlのQRコードを配置  



## 使い方

<https://my.slack.com/services/new/bot>
こちらに接続して新しいbotを作ります。
APIKeyは使うのでメモしておいてください。

<https://github.com/yutaro/slack-tepra/releases>
次にこちらからslack-bot本体をインストールします。
config.tomlにAPIKEYを記入し起動すれば利用可能状態になります。
