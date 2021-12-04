package main

import (
	"bot-slack/db"
	"bot-slack/env"
	"bot-slack/message"
	"bot-slack/querys"
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	//load envs
	env.LoadEnv()

	//data

	data := querys.GetMessage(db.Connect())

	// Create a new api client
	api := slack.New(os.Getenv("BOT_SLACK_TOKEN"))

	for _, v := range data {
		// Send a message to #general channel on the slack workspace
		channelID, timestamp, err := api.PostMessage(
			os.Getenv("SLACK_CHANNEL"),
			slack.MsgOptionText(

				message.SendMessage(v.Name, v.Description),

				false),
		)
		if err != nil {
			log.Printf("%s\n", err)
		}
		fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	}
}
