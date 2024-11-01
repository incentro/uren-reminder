package main

import "time"

func shouldSendMessage(time time.Time) bool {
	if isWeekDay(time.Weekday()) && isLastDayOfTheMonth(time) {
		return true
	}

	if isFriday(time.Weekday()) {
		return true
	}

	return false
}
