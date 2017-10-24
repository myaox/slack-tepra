package main

import (
	"strings"

	"github.com/yutaro/slack-cmd-go"
)

var (
	templates = map[string]string{
		"text":    "text.tpe",
		"text-qr": "text-qr.tpe",
	}
)

var exePath = "C:/Program Files (x86)/KING JIM/TEPRA SPC10/SPC10.exe"

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)
	tepra := bot.NewCmdGroup("tepra")

	tepra.Cmd("print", []string{"print message"},
		func(c *scmd.Context) {
			args := c.GetArgs()
			options := c.GetOptions()
			mes := strings.Join(args, " ")

			//tpe := templates["text"]

			prints := []string{mes}

			if qr, ok := options["qr"]; ok {
				url := urlConv(qr)
				prints = append(prints, url)
			}

			//csvpath := writeCsv(prints)
			writeCsv(prints)
			//cmd := exec.Command("sleep", "5s")
			//cmd.Start()
			//cmd.Wait()

			c.SendMessage(mes)
		})

	tepra.Cmd("qrcode", []string{"print only qrcode"},
		func(c *scmd.Context) {
			args := c.GetArgs()
			if len(args) < 1 {
				return
			}

			qr := args[0]
			url := urlConv(qr)

			c.SendMessage(url)
		})

	tepra.Cmd("image", []string{"print image"},
		func(c *scmd.Context) {

		})

	tepra.Cmd("status", []string{"check status"},
		func(c *scmd.Context) {
			c.SendMessage("ok")
		})

	bot.Start()
}
