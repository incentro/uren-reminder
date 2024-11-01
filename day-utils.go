package main

import "time"

func isWeekDay(day time.Weekday) bool {
	return day >= time.Monday && day <= time.Friday
}

func isFriday(day time.Weekday) bool {
	return day == time.Friday
}

func isLastDayOfTheMonth(day time.Time) bool {
	return day.Day() == daysInMonth(day)
}

func daysInMonth(day time.Time) int {
	return time.Date(day.Year(), day.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
