// report repeated lines in a file
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	flagFields    = 0
	flagLetters   = 0
	flagMode      = ' '
	firstRun      bool
	prevOriginal  string
	prevLine      string
	prevLineCount int
	originals     []string
	lines         []string
)

func main() {
	var err error
	in := os.Stdin
	out := os.Stdout

	args := os.Args[1:]
	for n := 0; n < len(args); n++ {
		arg := args[n]

		if arg[0] == '-' {
			num, err := strconv.ParseInt(arg[1:], 10, 32)
			if err == nil {
				flagFields = int(num)
			} else if len(arg) > 1 {
				flagMode = rune(arg[1])
			}
			continue
		}
		if arg[0] == '+' {
			num, err := strconv.ParseInt(arg[1:], 10, 32)
			if err == nil {
				flagLetters = int(num)
			}
			continue
		}

		in, err = os.Open(arg)
		if err != nil {
			fmt.Println(os.Stderr, "uniq:", err)
			os.Exit(1)
		}
		break
	}

	defer in.Close()

	exitStatus := uniq(in, out)
	os.Exit(exitStatus)
}

func uniq(in io.Reader, out io.Writer) (exitStatus int) {
	reader := bufio.NewReader(in)

	switch flagMode {
	case ' ':
		return noDuplicates(reader, out)
	case 'u':
		return onlyUnique(reader, out)
	case 'd':
		return onlyDuplicates(reader, out)
	case 'c':
		return withStats(reader, out)
	}

	fmt.Fprintln(os.Stderr, "usage: uniq [ -udc [ +-num ] ] [file ]")
	return 1
}

func onlyUnique(in *bufio.Reader, out io.Writer) (exitStatus int) {
	cmps := make([]*string, 3)
	lines := make([]*string, 3)

	for {
		line, cmp, err := readLine(in)

		if err != nil {
			if err == io.EOF {
				break
			}
			return 1
		}

		cmps = append(cmps[1:], &cmp)
		lines = append(lines[1:], &line)

		if cmps[1] == nil {
			// got 1 line
			continue
		}
		if cmps[0] == nil {
			// got 2 lines
			if *cmps[1] != *cmps[2] {
				fmt.Fprintln(out, *lines[1])
			}
		} else {
			if *cmps[0] != *cmps[1] && *cmps[1] != *cmps[2] {
				fmt.Fprintln(out, *lines[1])
			}
		}
	}

	if cmps[0] != nil {
		// got three lines
		if *cmps[1] != *cmps[2] {
			fmt.Fprintln(out, *lines[2])
		}
	} else if cmps[1] != nil {
		// got two lines
		if *cmps[1] != *cmps[2] {
			fmt.Fprintln(out, *lines[2])
		}
	} else if cmps[2] != nil {
		// got one line
		fmt.Fprintln(out, *lines[2])
	}

	return 0
}

func onlyDuplicates(in *bufio.Reader, out io.Writer) (exitStatus int) {
	cmps := make([]*string, 3)
	lines := make([]*string, 3)

	for {
		line, cmp, err := readLine(in)

		if err != nil {
			if err == io.EOF {
				break
			}
			return 1
		}

		cmps = append(cmps[1:], &cmp)
		lines = append(lines[1:], &line)

		if cmps[0] == nil {
			// wait for more lines
			continue
		}

		if *cmps[0] == *cmps[1] && *cmps[1] != *cmps[2] {
			fmt.Fprintln(out, *lines[1])
		}
	}

	if cmps[2] == nil {
		// got no lines
	} else if cmps[1] == nil {
		// got one line
	} else if cmps[0] == nil {
		// got two lines
		if *cmps[1] == *cmps[2] {
			fmt.Fprintln(out, *lines[1])
		} else {
			fmt.Fprintln(out, *lines[1])
			fmt.Fprintln(out, *lines[2])
		}
	} else {
		// got three lines
		if *cmps[1] == *cmps[2] {
			fmt.Fprintln(out, *lines[2])
		} else {
			fmt.Fprintln(out, *lines[2])
		}
	}

	return
}

func withStats(in *bufio.Reader, out io.Writer) (exitStatus int) {
	count := 1
	var prevLine string
	var prevCmp *string

	for {
		line, cmp, err := readLine(in)

		if err != nil {
			if err == io.EOF {
				break
			}
			return 1
		}

		if prevCmp != nil {
			if *prevCmp == cmp {
				count++
			} else {
				fmt.Fprintf(out, "%4d %s\n", count, prevLine)
				count = 1
			}
		}

		prevLine, prevCmp = line, &cmp
	}

	if prevCmp != nil {
		fmt.Fprintf(out, "%4d %s\n", count, prevLine)
	}

	return
}

func noDuplicates(in *bufio.Reader, out io.Writer) (exitStatus int) {
	var prevLine string
	var prevCmp *string

	for {
		line, cmp, err := readLine(in)

		if err != nil {
			if err == io.EOF {
				break
			}
			return 1
		}

		if prevCmp != nil {
			if *prevCmp != cmp {
				fmt.Fprintln(out, prevLine)
			}
		}

		prevLine, prevCmp = line, &cmp
	}

	if prevCmp != nil {
		fmt.Fprintln(out, prevLine)
	}

	return
}

var (
	linePart []byte
	lineBuf  = bytes.Buffer{}
)

func readLine(in *bufio.Reader) (line string, cmp string, err error) {
	for isPrefix := true; isPrefix; {
		linePart, isPrefix, err = in.ReadLine()
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Fprintln(os.Stderr, "uniq:", err)
			return
		}
		lineBuf.Write(linePart)
	}

	line = lineBuf.String()
	lineBuf.Reset()

	return line, skip(line, flagFields, flagLetters), nil
}

func skip(line string, fields, letters int) string {
	for n := 0; n < fields; n++ {
		idx := strings.IndexAny(line, " \t")
		if idx > 0 {
			line = line[idx+1:]
		}
	}

	if letters > 0 {
		if len(line) > letters {
			line = line[letters:]
		} else {
			line = ""
		}
	}

	return line
}
