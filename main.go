package main

import (
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"math/rand"
	"os"
	"time"
)

type MessageFile struct {
	Messages []string `json:"messages"`
}

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	afasUrl := os.Getenv("AFAS_URL")
	client := slack.New(token)

	jsonFile, err := os.ReadFile("input/messages.json")
	if err != nil {
		fmt.Printf("Failed to read JSON file: %v\n", err)
		return
	}

	var file MessageFile
	err = json.Unmarshal(jsonFile, &file)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		return
	}

	randomIndex := rand.Intn(len(file.Messages))
	selectedMessage := createBlock(file.Messages[randomIndex], afasUrl)
	if shouldSendMessage(time.Now()) {
		_, _, err := client.PostMessage(channelID, slack.MsgOptionBlocks(selectedMessage...))
		if err != nil {
			log.Fatal("Unable to send message to Slack", err)
			return
		}
		println("Reminder successfully sent")
		return
	}

	println("No need to send reminder")
}
