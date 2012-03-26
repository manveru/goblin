package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	flagPermission = flag.String("m", "0777", "permissions created directory")
	flagParent     = flag.Bool("p", false, "create parent directories")
)

func main() {
	flag.Parse()

	umask, err := strconv.ParseInt(*flagPermission, 8, 32)
	if err != nil || umask < 0 || umask > 0777 {
		usage(1)
	}

	if flag.NArg() < 1 {
		usage(1)
	}

	mode := os.ModeDir | os.FileMode(umask)

	for _, arg := range flag.Args() {
		if *flagParent {
			os.MkdirAll(arg, mode)
		} else {
			os.Mkdir(arg, mode)
		}
	}
}

func usage(exitStatus int) {
	fmt.Fprintln(os.Stderr, "usage: mkdir [-p] [-m mode] dir...")
	os.Exit(exitStatus)
}
