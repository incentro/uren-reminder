package main

import (
	"testing"
	"time"
)

type shouldSendMessageTest struct {
	arg1     time.Time
	expected bool
}

var shouldSendMessageTests = []shouldSendMessageTest{
	{time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), false},   // Monday
	{time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC), true},   // Last day of the month
	{time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC), false}, // Leap year so not last day
	{time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC), true},  // Last day of the month
}

func TestShouldSendMessage(t *testing.T) {

	for _, test := range shouldSendMessageTests {
		if output := shouldSendMessage(test.arg1); output != test.expected {
			t.Errorf("Ouput %t is not equal to expected %t for time %s", output, test.expected, test.arg1)
		}
	}
}
