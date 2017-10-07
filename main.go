package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	Env_load()

	token := os.Getenv("TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	fmt.Println("--- start connection ---")

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				//
			case *slack.ConnectedEvent:
				//
			case *slack.TeamJoinEvent:
				//
			case *slack.MessageEvent:
				go evalComment(rtm, ev)

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
