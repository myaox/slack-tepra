## slackからtepraを操作

### 進捗発表 2017/11/10

---

### 前回まで

- slackをCLI化するライブラリ |
  - テプラとCLI化する部分を分離した |
  - <https://github.com/yutaro/slack-tepra> |
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
## コマンドラインオプションを実装

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

## 
