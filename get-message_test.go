package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

type getMessageTest struct {
	arg1     MessageFile
	arg2     time.Time
	expected string
}

var messageFile = MessageFile{
	Weekly:  []string{"Weekly message"},
	Monthly: []string{"Monthly message"},
}
var getMessageTests = []getMessageTest{
	{messageFile, time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), ""},                       // Monday
	{messageFile, time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC), messageFile.Monthly[0]},  // Last day of the month
	{messageFile, time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC), ""},                     // Leap year so not last day
	{messageFile, time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC), messageFile.Monthly[0]}, // Last day of the month
	{messageFile, time.Date(2024, time.December, 20, 0, 0, 0, 0, time.UTC), messageFile.Weekly[0]},  // Last day of the week
}

func TestGetMessage(t *testing.T) {
	for _, test := range getMessageTests {
		if output := getMessage(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Ouput %s is not equal to expected %s for time %s", output, test.expected, test.arg1)
		}
	}
}

type getRandomMessageTest struct {
	arg1     []string
	expected string
}

var getRandomMessageTests = []getRandomMessageTest{
	{[]string{}, ""},                   // Empty slice
	{[]string{"Message1"}, "Message1"}, // Single message
}

func TestGetRandomMessage(t *testing.T) {
	// Just testing here that a string is returned. Not necessarily that it's random.
	for _, test := range getRandomMessageTests {
		if output := getRandomMessage(test.arg1); output != test.expected {
			t.Errorf("Output %s is not in expected %v for input %v", output, test.expected, test.arg1)
		}
	}
}

func TestGetMessagesFromFile(t *testing.T) {
	validJSON := `{"weekly": ["Weekly message"], "monthly": ["Monthly message"]}`
	validFileName := "valid_messages.json"
	err := os.WriteFile(validFileName, []byte(validJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create valid JSON file: %v", err)
	}
	defer os.Remove(validFileName)

	invalidJSON := `{"weekly": ["Weekly message"], "monthly": ["Monthly message"`
	invalidFileName := "invalid_messages.json"
	err = os.WriteFile(invalidFileName, []byte(invalidJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create invalid JSON file: %v", err)
	}
	defer os.Remove(invalidFileName)

	tests := []struct {
		fileName string
		expected MessageFile
		hasError bool
	}{
		{validFileName, MessageFile{Weekly: []string{"Weekly message"}, Monthly: []string{"Monthly message"}}, false},
		{"non_existent_file.json", MessageFile{}, true},
		{invalidFileName, MessageFile{}, true},
	}

	for _, test := range tests {
		output, err := getMessagesFromFile(test.fileName)
		if (err != nil) != test.hasError {
			t.Errorf("Expected error: %v, got: %v", test.hasError, err)
		}
		if fmt.Sprint(output) != fmt.Sprint(test.expected) {
			t.Errorf("Expected output: %v, got: %v", test.expected, output)
		}
	}
}
