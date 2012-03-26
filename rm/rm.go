package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagRecurse = flag.Bool("r", false, "delete recursive")
	flagForce   = flag.Bool("f", false, "ignore errors")
)

func main() {
	flag.Parse()

	var err error
	for _, arg := range flag.Args() {
		if *flagRecurse {
			err = os.RemoveAll(arg)
		} else {
			err = os.Remove(arg)
		}
		if !*flagForce && err != nil {
			fmt.Fprintln(os.Stderr, "rm:", err)
		}
	}
}
