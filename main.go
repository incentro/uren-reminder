package main

import (
	"github.com/slack-go/slack"
	"log"
	"os"
	"time"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	client := slack.New(token)

	if shouldSendMessage(time.Now()) {
		_, _, err := client.PostMessage(channelID, slack.MsgOptionText("Uren reminder", false))
		if err != nil {
			log.Fatal("Unable to send message to Slack", err)
			return
		}
		println("Reminder successfully sent")
		return
	}

	println("No need to send reminder")
}
