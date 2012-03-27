package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: mtime file...")
	}

	exitStatus := 0
	for _, arg := range os.Args[1:] {
		stat, err := os.Stat(arg)
		if err != nil {
			exitStatus = 1
			fmt.Fprintf(os.Stderr, "stat: %s\n", err)
			continue
		}
		fmt.Printf("%11d %s\n", stat.ModTime().Unix(), arg)
	}

	os.Exit(exitStatus)
}
