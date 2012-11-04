// cat - catenate files
//
// SYNOPSIS
//   cat [ file ... ]
//   read [ -m ] [ -n nline ] [ file ... ]
//   nobs [ file ... ]
//
// DESCRIPTION
//   Cat reads each file in sequence and writes it on the stan-
//   dard output.  Thus
//
//     cat file
//
//   prints a file and
//
//     cat file1 file2 >file3
//
//   concatenates the first two files and places the result on
//   the third.
//
//   If no file is given, cat reads from the standard input.
//   Output is buffered in blocks matching the input.
//
//   Read copies to standard output exactly one line from the
//   named file, default standard input.  It is useful in inter-
//   active rc(1) scripts.
//
//   The -m flag causes it to continue reading and writing multi-
//   ple lines until end of file; -n causes it to read no more
//   than nline lines.
//
// BUGS
//   Beware of `cat a b >a' and `cat a b >b', which destroy input
//   files before reading them.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cat: %s\n", err)
			} else {
				io.Copy(os.Stdout, file)
				file.Close()
			}
		}
	}
}
