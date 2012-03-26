package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc < 1 || argc > 3 {
		os.Stderr.WriteString("usage: basename [-d] string [suffix]\n")
		os.Exit(1)
	}

	dflag := false

	if argc > 0 && args[0] == "-d" {
		dflag = true
		args = args[1:]
		argc -= 1
	}

	path := args[0]
	pathIndex := strings.LastIndex(path, "/")

	if dflag {
		if pathIndex != -1 {
			fmt.Println(path[:pathIndex])
		} else {
			fmt.Println(".")
		}
	} else {
		if pathIndex != 1 {
			path = path[pathIndex+1:]
		}
		switch argc {
		case 1:
			fmt.Println(path)
		case 2:
			suffix := args[1]
			fmt.Println(path, suffix)
			suffixIndex := strings.LastIndex(path, suffix)
			if len(path)-(suffixIndex+len(suffix)) == 0 {
				fmt.Println(path[:suffixIndex])
			} else {
				fmt.Println(path)
			}
		}
	}
}
