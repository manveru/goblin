package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	flagFormat = flag.String("f", "%g", "format")
	flagWidth  = flag.Bool("w", false, "equalize width")
)

func main() {
	flag.Parse()

	var first, incr, last float64

	switch flag.NArg() {
	case 0:
		fmt.Fprintln(os.Stderr, "usage: seq [-fformat] [-w] [first [incr]] last")
	case 1:
		last = parseFloat(0)
		seq(1, 1, last)
	case 2:
		first = parseFloat(0)
		last = parseFloat(1)
		seq(first, 1, last)
	case 3:
		incr = parseFloat(1)
		if incr == 0 {
			fmt.Fprintln(os.Stderr, "seq: zero increment")
		} else {
			first = parseFloat(0)
			last = parseFloat(2)
			seq(first, incr, last)
		}
	}
}

func seq(first, incr, last float64) {
	format := *flagFormat + "\n"
	if *flagWidth {
		maxW, maxP := 0, 0
		for n := first; n <= last; n += incr {
			wp := strings.SplitN(fmt.Sprintf("%g", n), ".", 2)
			if len(wp[0]) > maxW {
				maxW = len(wp[0])
			}
			if len(wp) > 1 {
				if len(wp[1]) > maxP {
					maxP = len(wp[1])
				}
			}
		}
		if maxP > 0 {
			maxW += (maxP + 1)
		}
		format = fmt.Sprintf("%%%d.%df\n", maxW, maxP)
	}

	for n := first; n <= last; n += incr {
		fmt.Printf(format, n)
	}
}

func parseFloat(n int) float64 {
	f, err := strconv.ParseFloat(flag.Arg(n), 64)
	if err != nil {
		return 1
	}
	return f
}
