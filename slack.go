package main

import (
	"github.com/nlopes/slack"
)

func reply(rtm *slack.RTM, ev *slack.MessageEvent, mes string) {
	rtm.SendMessage(rtm.NewOutgoingMessage(mes, ev.Channel))
}
