package main

import (
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		duration, err := time.ParseDuration(os.Args[1])
		if err != nil {
			time.Sleep(duration)
		}
	}
}
