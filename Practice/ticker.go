package main

import (
	"fmt"
	"time"
)

func ticker() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	var cnt int = 0
	for {
		select {
		case now := <-t.C:
			fmt.Println(now.Format(time.RFC3339))
			fmt.Println(cnt)
			cnt = 0
		default:
			cnt++
		}
		// other processing
	}
}

func main() {
	ticker()
}
