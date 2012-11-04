// cleanname cleans a path name
//
// SYNOPSIS
//   cleanname [ -d pwd ] names ...
//
// DESCRIPTION
//   For each file name argument, cleanname, by lexical processing
//   only, prints the shortest equivalent string that names the
//   same (possibly hypothetical) file.
//   It eliminates multiple and trailing slashes, and it lexically
//   interprets . and .. directory components in the name. If the -d
//   option is present, unrooted names are prefixed with pwd/ before
//   processing.
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
