package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc < 1 {
		strings(os.Stdin)
		return
	}

	for _, arg := range args {
		if argc > 2 {
			print(arg, ":\n")
		}

		f, err := os.Open(arg)
		if err != nil {
			os.Stderr.WriteString("strings: " + err.Error())
			continue
		}

		strings(f)
	}
}

const (
	minspan = 6
	bufsize = 70
)

func strings(f *os.File) {
	reader := bufio.NewReader(f)
	buf := make([]rune, 0, 70)

	var start, posn int
	for {
		r, size, err := reader.ReadRune()
		if size < 1 {
			break
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				os.Stderr.WriteString("strings: " + err.Error())
			}
		}

		if unicode.IsPrint(r) {
			if start == 0 {
				start = posn
			}

			buf = append(buf, r)

			if len(buf) >= bufsize-1 {
				fmt.Printf("%8d: %s ...\n", start, string(buf))
				start = 0
				buf = buf[:0]
			}
		} else {
			if len(buf) >= minspan {
				fmt.Printf("%8d: %s\n", start, string(buf))
			}
			start = 0
			buf = buf[:0]
		}

		posn += size
	}

	if len(buf) >= minspan {
		fmt.Printf("%8d: %s\n", start, string(buf))
	}
}
