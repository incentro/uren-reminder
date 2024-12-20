package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type MessageFile struct {
	Weekly  []string `json:"weekly"`
	Monthly []string `json:"monthly"`
}

func getMessagesFromFile(fileName string) (MessageFile, error) {
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Failed to read JSON file: %v\n", err)
		return MessageFile{}, err
	}

	var file MessageFile
	err = json.Unmarshal(jsonFile, &file)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		return MessageFile{}, err
	}
	return file, nil
}

func getMessage(messageFile MessageFile, time time.Time) string {
	if isWeekDay(time.Weekday()) && isLastDayOfTheMonth(time) {
		return getRandomMessage(messageFile.Monthly)
	}

	if isFriday(time.Weekday()) {
		return getRandomMessage(messageFile.Weekly)
	}

	return ""
}

func getRandomMessage(messages []string) string {
	randomIndex := rand.Intn(len(messages))
	return messages[randomIndex]
}
