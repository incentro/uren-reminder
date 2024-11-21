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

func getRandomMessageFromFile(fileName string) (string, error) {
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Failed to read JSON file: %v\n", err)
		return "", err
	}

	var file MessageFile
	err = json.Unmarshal(jsonFile, &file)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		return "", err
	}

	randomIndex := rand.Intn(len(file.Messages))
	return file.Messages[randomIndex], nil
}

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	afasUrl := os.Getenv("AFAS_URL")
	client := slack.New(token)

	msg, err := getRandomMessageFromFile("input/messages.json")
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		return
	}

	blocks := createBlock(msg, afasUrl)
	if shouldSendMessage(time.Now()) {
		_, _, err := client.PostMessage(channelID, slack.MsgOptionBlocks(blocks...))
		if err != nil {
			log.Fatal("Unable to send message to Slack", err)
			return
		}
		println("Reminder successfully sent")
		return
	}

	println("No need to send reminder")
}
