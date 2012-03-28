package main

import (
	"os"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		os.Stdout.WriteString("???")
	} else {
		idx := strings.LastIndex(wd, "/")
		os.Stdout.WriteString(wd[idx+1:])
	}
}
