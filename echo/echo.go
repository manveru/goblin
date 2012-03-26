// echo prints its arguments separated by blanks and terminated by a newline
// to standard output.
package main

import (
	"bufio"
	"flag"
	"os"
)

var (
	noNewline = flag.Bool("n", false, "surpress newline")
)

func main() {
	flag.Parse()

	writer := bufio.NewWriter(os.Stdout)
	argc := flag.NArg()
	for n, flag := range flag.Args() {
		if n < argc {
			writer.WriteString(flag + " ")
		} else {
			writer.WriteString(flag)
		}
	}

	if !*noNewline {
		writer.WriteString("\n")
	}

	writer.Flush()
}
