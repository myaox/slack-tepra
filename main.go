package main

import (
	"os/exec"
	"strings"

	"github.com/yutaro/slack-cmd-go"
)

var (
	imgpath = imgPath("result")
	csvpath = csvPath("value.csv")
)

// TODO : 大体ここにあると思うけど別のケースも想定する挙動を書いたほうがいいかも
var exepath = "C:/Program Files (x86)/KING JIM/TEPRA SPC10/SPC10.exe"

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)
	tepra := bot.NewCmdGroup("tepra")

	tepra.Cmd("print",
		[]string{
			"print message",
			"-n=x  x枚印刷します",
			"--qr=URL  URLに接続するQRコードを左側に印刷",
			"-t   印刷せずテスト画像を表示"},
		func(c *scmd.Context) {
			args := c.GetArgs()
			options := c.GetOptions()
			flags := c.GetFlags()

			mes := strings.Join(args, " ")

			tpe := "text"

			prints := []string{mes}

			if qr, ok := options["qr"]; ok {
				url := urlConv(qr)
				prints = append(prints, url)
				tpe = "text_qr"

				c.SendMessage("qrcode : " + url)
			}

			writeCsv(prints)
			print(c, tpePath(tpe))

			c.SendMessage("text     : " + mes)

			reply := ""
			if !flags["t"] {
				reply += "印刷しています..."
			} else {
				reply += "テスト画像を表示します..."
			}

			c.SendMessage(reply)
		})

	tepra.Cmd("qrcode", []string{"print only qrcode"},
		func(c *scmd.Context) {
			args := c.GetArgs()
			if len(args) < 1 {
				return
			}

			qr := args[0]
			url := urlConv(qr)

			prints := []string{url}

			csvpath := writeCsv(prints)
			tpepath := tpePath("qr")

			n := "1"
			if num, ok := c.GetOptions()["n"]; ok {
				n = num
			}

			cmd := exec.Command(exepath, "-p", tpepath+","+csvpath+","+n)
			cmd.Run()

			c.SendMessage(url)
		})

	tepra.Cmd("image", []string{"print image", "don't work now..."},
		func(c *scmd.Context) {
			c.SendMessage("preparing now...")
		})

	tepra.Cmd("status", []string{"check status"},
		func(c *scmd.Context) {
			c.SendMessage("ok")
		})

	bot.Start()
}
