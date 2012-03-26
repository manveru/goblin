package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	flagGMT   = flag.Bool("u", false, "Report Greenwich Mean Time (GMT) rather than local time.")
	flagEpoch = flag.Bool("n", false, "Report the date as the number of seconds since the epoch, 00:00:00 GMT, January 1, 1970.")
)

func main() {
	flag.Parse()

	var now time.Time
	if flag.NArg() == 1 {
		seconds, err := strconv.ParseInt(flag.Arg(0), 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "date: %s\n", err)
			os.Exit(0)
		}
		now = time.Unix(seconds, 0)
	} else {
		now = time.Now()
	}

	if *flagEpoch {
		fmt.Fprintf(os.Stdout, "%d\n", now.Unix())
	} else if *flagGMT {
		fmt.Fprintf(os.Stdout, "%s\n", now.UTC().Format(time.UnixDate))
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", now.Format(time.UnixDate))
	}
}
