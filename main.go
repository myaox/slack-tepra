package main

import (
	"github.com/yutaro/slack-cmd-go"
)

func main() {
	conf := scmd.LoadToml("config.toml")
	bot := scmd.New(conf.TOKEN)
	tepra := bot.NewCmds("tepra")

	tepra.Cmd("print", "print message",
		func(c *scmd.Context) {
			c.SendMessage("hello")
		})

	bot.Start()
}
