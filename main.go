package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/nlopes/slack"
)

type Config struct {
	TOKEN string
}

func LoadToml(path string) *Config {
	var conf Config
	_, err := toml.DecodeFile(path, &conf)
	if err == nil {
		return &conf
	}

	basepath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fullpath := filepath.Join(basepath, path)

	_, err = toml.DecodeFile(fullpath, &conf)
	if err != nil {
		panic(err)
	}
	return &conf
}

func main() {
	conf := LoadToml("config.toml")

	api := slack.New(conf.TOKEN)
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
