package main

import (
	"strings"

	"github.com/nlopes/slack"
)

var cmds = make(map[string]command, 100)

type tepraEnv struct {
	count int
}

func evalComment(rtm *slack.RTM, ev *slack.MessageEvent) {
	msg := ev.Text
	args := strings.Split(msg, " ")

	if args[0] != "tepra" {
		return
	}

	var cmd command
	if val, ok := cmds[args[1]]; ok {
		cmd = val
	} else {
		// TODO : make error message
		return
	}

	cmd.run(rtm, ev)
}
