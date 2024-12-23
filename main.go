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
	afasUrl := os.Getenv("AFAS_URL")
	client := slack.New(token)

	messages, err := getMessagesFromFile("input/messages.json")
	if err != nil {
		log.Fatal("Failed to parse JSON", err)
		return
	}

	message := getMessage(messages, time.Now())
	if len(message) == 0 {
		println("No need to send reminder")
		return
	}

	blocks := createBlock(message, afasUrl)

	_, _, e := client.PostMessage(channelID, slack.MsgOptionBlocks(blocks...))
	if e != nil {
		log.Fatal("Unable to send message to Slack", err)
		return
	}
	println("Reminder successfully sent")
}
