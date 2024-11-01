package main

import "time"

func main() {

	if shouldSendMessage(time.Now()) {
		println("Send last day of the month message")
	}
}
