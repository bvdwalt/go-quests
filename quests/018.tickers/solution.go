package tickers

import (
	"fmt"
	"time"
)

func Ticker() {
	helloTicker := time.NewTicker(500 * time.Millisecond)
	worldTicker := time.NewTicker(1 * time.Second)
	timeout := time.After(3 * time.Second)
	defer helloTicker.Stop()
	defer worldTicker.Stop()

	for {
		select {
		case <-helloTicker.C:
			fmt.Println("hello")
		case <-worldTicker.C:
			fmt.Println("world")
		case <-timeout:
			return
		}
	}
}
