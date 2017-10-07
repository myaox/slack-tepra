package main

import (
	"github.com/nlopes/slack"
)

type command interface {
	run(*slack.RTM, *slack.MessageEvent)
}

type cmdInfo struct {
	name      string
	helpMes   string
	detailMes string
}

func (c *cmdInfo) getHelp() string {
	return c.name + " - " + c.helpMes
}

type tepraCmd struct {
	cmdInfo
	options []optionCmd
}

func (t *tepraCmd) getDetail() string {
	str := t.getHelp() + "\n   " + t.detailMes
	for _, opt := range t.options {
		str += "\n     " + opt.getHelp()
	}
	return str
}

type optionCmd struct {
	cmdInfo
}

func (o *optionCmd) getDetail() string {
	return o.getHelp() + "\n   " + o.detailMes
}

type printCmd tepraCmd

func (p *printCmd) run(rtm *slack.RTM, ev *slack.MessageEvent) {
	reply(rtm, ev, p.getHelp())
}

func init() {
	cmds["print"] = &printCmd{
		cmdInfo: cmdInfo{
			name:      "print",
			helpMes:   "基本的な印刷を行うコマンド",
			detailMes: "工事中",
		},
		options: []optionCmd{},
	}
}
