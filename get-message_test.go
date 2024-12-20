package main

import (
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
