package main

import (
	"testing"
	"time"
)

type isWeekDayTest struct {
	arg1     time.Weekday
	expected bool
}

var isWeekDayTests = []isWeekDayTest{
	{time.Monday, true},
	{time.Tuesday, true},
	{time.Wednesday, true},
	{time.Thursday, true},
	{time.Friday, true},
	{time.Saturday, false},
	{time.Sunday, false},
}

func TestIsWeekDay(t *testing.T) {
	for _, test := range isWeekDayTests {
		if output := isWeekDay(test.arg1); output != test.expected {
			t.Errorf("Output %t not equal to expected %t for argument %s", output, test.expected, test.arg1)
		}
	}
}

func TestIsFriday(t *testing.T) {
	output := isFriday(time.Friday)
	expected := true
	if output != expected {
		t.Errorf("Output %t not equal to expected %t", output, expected)
	}
}

func TestIsFridayFalse(t *testing.T) {
	output := isFriday(time.Monday)
	expected := false
	if output != expected {
		t.Errorf("Output %t not equal to expected %t", output, expected)
	}
}

func TestDaysInMonth(t *testing.T) {
	output := daysInMonth(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))
	expected := 31
	if output != expected {
		t.Errorf("Output %d not equal to expected %d", output, expected)
	}
}

func TestIsLastDayOfTheMonth(t *testing.T) {
	output := isLastDayOfTheMonth(time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC))
	expected := true
	if output != expected {
		t.Errorf("Output %t not equal to expected %t", output, expected)
	}
}

func TestIsLastDayOfTheMonthFalse(t *testing.T) {
	output := isLastDayOfTheMonth(time.Date(2024, time.January, 30, 0, 0, 0, 0, time.UTC))
	expected := false
	if output != expected {
		t.Errorf("Output %t not equal to expected %t", output, expected)
	}
}
