## slackからtepraを操作

### 進捗発表 2017/11/10

---

### 前回まで

- slackをCLI化するライブラリ |
  - テプラとCLI化する部分を分離した |
  <https://github.com/yutaro/slack-tepra>
- 結果 : TepraBotを比較的綺麗に書きなおせた! |

---

## TepraBotの機能を増やしたい!

シンプルなユースケースの  
使い勝手はそのままにしたい

```sh
tepra print Hello World
```

---

## CLI化ライブラリに
### コマンドラインオプションを実装

---

## flag

```sh
tepra print -ta --flag --something
```

[t, a, flag, something] 
フラグがtrueにで実行される  
(デフォルトはfalse)

+++

### 使い方

```go
func(c *scmd.Context) {
	flags := c.GetFlags()

  if flags["t"] {
          c.SendMessage("flag t : true")
  }

  if flags["something"] && flags["a"]{
          c.SendMessage("flag a and something : true")
  }
}
```

---

## option

```sh
tepra print -n=10 --qr=google.com
```

n = "10", qr = "google.com" の状態で実行される
(string型で渡される)

+++

### 使い方

```go
func(c *scmd.Context){
  options := c.GetOptions()

  var n int
  if num, ok := options["n"]; ok{
      n , _ = strconv.Atoi(num)
  }else{
      n = 1
  }

  c.SendMessage(fmt.Sprintf("n = %d で実行されました", n))
}
```

---

## TepraBotのアップデート内容

- -n=10 10枚印刷 |
- -t オプションで印刷せずにテスト画像を表示 |
- --qr=url qrコードを左側に印刷 |
- tepra qr url で qrコードのみを印刷 |

+++

### サンプル画像
![cable](./imgs/cable.bmp)
![wellcome](./imgs/wellcome.bmp)
![qr](./imgs/qr.bmp)


---

# これから

---
## LabelWriter 450 Turbo

- 買っていただいたので色々試したい
- Linux環境で動くらしい |
- UbuntuノートPC環境があるので、それで色々できるかテスト|
- <https://github.com/dsandor/dymojs> |
- 良さげなライブラリ |

+++

### 見た目

![dymo](https://s3-ap-southeast-2.amazonaws.com/wc-prod-pim/JPEG_1000x1000/SALM450TUR_dymo_dymo_labelwriter_450_turbo_label_printer_silver.jpg)

---

## Slash Commandsを試す

```sh
/command aaaa aaa aaa
```

- というような形式のコマンド
- Webサーバーにリクエストを送るようなものらしい
- Linux環境ならそれでWebサーバー立てたほうがいいかも？

+++

## いいところ

- 補完が効く
  - /co ぐらいで候補が表示される
- 便利なライブラリとかがありそう... |
<https://botkit.ai/> |
<http://starzero.hatenablog.com/entry/2016/02/07/000128> |
- そんなに難しくなさそう |

---

## 来週まで
nodejsで実装できそうなので試してみる

---

# 課題

- 簡単に文字が印刷できて便利以上の価値が見出せていない
- ラベルを印刷する機会ってそんなに多くないんじゃないか

- こんなラベルが印刷できたら便利かも
- ラベルの意味が大きい場面について考える必要


