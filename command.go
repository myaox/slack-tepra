package main

import "fmt"

type command interface {
	run()
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

func (p *printCmd) run() {
	fmt.Println("running print")
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
