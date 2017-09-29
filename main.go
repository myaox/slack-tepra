package main

import (
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
				//
			case *slack.ReactionAddedEvent:
				//
			case *slack.RTMError:
				//

			default:
			}
		}

	}
}
