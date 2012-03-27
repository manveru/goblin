package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

var (
	flagLines = flag.Bool("l", false, "count lines")
	flagWords = flag.Bool("w", false, "count words")
	flagBytes = flag.Bool("c", false, "count bytes")
	flagRunes = flag.Bool("r", false, "count runes")
	flagBrokn = flag.Bool("b", false, "count broken UTF codes")
)

type report struct {
	name                              string
	lines, words, bytes, runes, brokn uint64
}

func (r report) Print() {
	line := []string{}
	if *flagLines {
		line = append(line, fmt.Sprintf("%7d", r.lines))
	}
	if *flagWords {
		line = append(line, fmt.Sprintf("%7d", r.words))
	}
	if *flagRunes {
		line = append(line, fmt.Sprintf("%7d", r.runes))
	}
	if *flagBrokn {
		line = append(line, fmt.Sprintf("%7d", r.brokn))
	}
	if *flagBytes {
		line = append(line, fmt.Sprintf("%7d", r.bytes))
	}
	if r.name != "" {
		line = append(line, fmt.Sprintf("%s", r.name))
	}

	fmt.Println(strings.Join(line, " "))
}

func main() {
	flag.Parse()

	if flag.NFlag() < 1 {
		*flagLines = true
		*flagWords = true
		*flagBytes = true
	}

	exitStatus := 0

	if flag.NArg() < 1 {
		count(os.Stdin, "<wc>").Print()
	} else if flag.NArg() > 1 {
		total := report{name: "total"}

		for _, arg := range flag.Args() {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wc: %s\n", err)
				exitStatus = 1
				continue
			}

			report := count(f, arg)
			f.Close()
			report.Print()

			total.brokn += report.brokn
			total.bytes += report.bytes
			total.lines += report.lines
			total.runes += report.runes
			total.words += report.words
		}
		total.Print()
	} else {
		for _, arg := range flag.Args() {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprint(os.Stderr, "wc: %s\n", err)
				exitStatus = 1
			}

			report := count(f, arg)
			f.Close()
			report.Print()
		}
	}

	os.Exit(exitStatus)
}

func count(reader io.Reader, name string) (total *report) {
	total = &report{name: name}

	inWord := false
	buf := bufio.NewReader(reader)
	for {
		r, s, err := buf.ReadRune()
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "wc: %s", err)
			}
			return
		}

		total.bytes += uint64(s)

		switch r {
		case '\n':
			total.runes++
			total.lines++
			if inWord {
				total.words++
				inWord = false
			}
		case unicode.ReplacementChar:
			total.brokn++
		default:
			if unicode.IsSpace(r) {
				total.runes++
				if inWord {
					total.words++
					inWord = false
				}
			} else {
				total.runes++
				inWord = true
			}
		}
	}

	return
}
