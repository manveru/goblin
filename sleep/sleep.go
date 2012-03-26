package main

import (
	"os"
	"time"
)

func main() {
	if len(os.Args) > 0 {
		duration, err := time.ParseDuration(os.Args[1])
		if err != nil {
			duration, err = time.ParseDuration(os.Args[1] + "s")
		}
		if err == nil {
			time.Sleep(duration)
		}
	}
}
