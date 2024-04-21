package main

import (
	"fmt"
	"time"
)

func main() {
	heartbeat := time.NewTicker(1 * time.Second)
	// 停止Ticker
	defer heartbeat.Stop()
	for {
		select {
		case <-heartbeat.C:
			fmt.Println("This is a demo for ticker")
		}
	}
}
