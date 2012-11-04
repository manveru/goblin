// date and time
//
// SYNOPSIS
//   date [ option ] [ seconds ]
//
// DESCRIPTION
//   Print the date, in the format
//
//        Tue Aug 16 17:03:52 CDT 1977
//
//   The options are
//
//   -u   Report Greenwich Mean Time (GMT) rather than local
//        time.
//
//   -n   Report the date as the number of seconds since the
//        epoch, 00:00:00 GMT, January 1, 1970.
//
//   The conversion from Greenwich Mean Time to local time
//   depends on the $timezone environment variable; see ctime(3).
//
//   If the optional argument seconds is present, it is used as
//   the time to convert rather than the real time.
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
