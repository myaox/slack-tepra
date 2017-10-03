package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NetRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				//
			case *slack.TeamJoinEvent:
				//
			case *slack.MessageEvent:
				evalComment(rtm, ev)

			case *slack.ReactionAddedEvent:
				//
			case *slack.ReactionRemovedEvent:
				//
			case *slack.RTMError:
				fmt.Println("Err:%s\n", ev.Error())
				break Loop
			case *slack.InvalidAuthEvent:
				fmt.Println("Invalid credentials")
				break Loop
			default:
			}
		}
	}
}
