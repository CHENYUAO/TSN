package main

import (
	"fmt"
	"time"
)

var (
	Count     int64
	Counter   *time.Ticker = time.NewTicker(TICK)
	Scheduler *time.Ticker = time.NewTicker(SLOT)
)

func sendMessage(msg string) {
	// sleep
	time.Sleep(time.Duration((len(msg) * 8 / DATA_RATE) * int(time.Millisecond)))
	fmt.Println("[", Count, "] send msg success")
}

func main() {
	for {
		select {
		case <-Scheduler.C:
			sendMessage("test msg")
		case <-Counter.C:
			Count++
		default:
			continue
		}
	}
}
