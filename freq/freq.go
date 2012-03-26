package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	flagDecimal = flag.Bool("d", true, "decimal")
	flagHex     = flag.Bool("x", true, "hex")
	flagOctal   = flag.Bool("o", true, "octal")
	flagChar    = flag.Bool("c", true, "character")
	flagRune    = flag.Bool("r", false, "count UTF sequences instead of bytes")
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		freq(os.Stdout, os.Stdin)
		return
	}

	for _, arg := range flag.Args() {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		} else {
			freq(os.Stdout, f)
			f.Close()
		}
	}
}

func freq(writer io.Writer, reader io.Reader) {
	wBuf := bufio.NewWriter(writer)
	rBuf := bufio.NewReader(reader)
	var histogram []rune

	if *flagRune {
		histogram = make([]rune, 1<<16)
		for {
			r, _, err := rBuf.ReadRune()
			if err == io.EOF {
				break
			}
			histogram[r] += 1
		}
	} else {
		histogram = make([]rune, 255)
		for {
			c, err := rBuf.ReadByte()
			if err == io.EOF {
				break
			}
			histogram[c] += 1
		}
	}

	for r, count := range histogram {
		if count == 0 {
			continue
		}
		if *flagDecimal {
			fmt.Fprintf(wBuf, "%3d ", r)
		}
		if *flagOctal {
			fmt.Fprintf(wBuf, "%03o ", r)
		}
		if *flagHex {
			fmt.Fprintf(wBuf, "%02x ", r)
		}
		if *flagChar {
			if r <= 0x20 || r >= 0x7f && r < 0xa0 || r > 0xff && !*flagRune {
				fmt.Fprint(wBuf, "- ")
			} else {
				fmt.Fprintf(wBuf, "%c", r)
			}
		}
		fmt.Fprintf(wBuf, "%8d\n", count)
	}

	wBuf.Flush()
}
