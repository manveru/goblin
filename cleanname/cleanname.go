package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var dflag *string = flag.String("d", "", "unrooted names are prefixed with [pwd]/ before processing.")

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		os.Stderr.WriteString("usage: cleanname [-d pwd] name...\n")
		os.Exit(1)
	}

	if dflag == nil {
		for _, arg := range flag.Args() {
			fmt.Println(filepath.Clean(arg))
		}
	} else {
		for _, arg := range flag.Args() {
			fmt.Println(filepath.Clean(filepath.Join(*dflag, arg)))
		}
	}
}
