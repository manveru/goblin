package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	flagMulti = flag.Bool("m", false, "continue reading and writing multipe files until end of file")
	flagLines = flag.Int("n", -1, "read no more than n lines")
)

func main() {
	flag.Parse()

	var err error

	if flag.NArg() == 0 {
		err = read(os.Stdin)
	} else {
		var f *os.File
		for _, arg := range flag.Args() {
			f, err = os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, "read:", err)
				os.Exit(1)
			}
			err = read(f)
			f.Close()
		}
	}

	if err != nil {
		os.Exit(1)
	}
}

func read(f *os.File) (err error) {
	reader := bufio.NewReader(f)

	if *flagMulti {
		for {
			err = readCore(reader)
			if err != nil {
				break
			}
		}
	} else if *flagLines > 0 {
		n := *flagLines
		for ; n > 0; n-- {
			err = readCore(reader)
			if err != nil {
				break
			}
		}
		if n > 0 {
			err = errors.New("less than n lines read")
		}
	} else {
		err = readCore(reader)
	}

	return
}

func readCore(reader *bufio.Reader) (err error) {
	line, err := readLine(reader)
	if err != nil {
		if err != io.EOF {
			fmt.Fprintln(os.Stderr, "read:", err)
		}
		return err
	}

	os.Stdout.Write(line)
	return
}

func readLine(reader *bufio.Reader) (line []byte, err error) {
	var lineBuf bytes.Buffer
	var linePart []byte
	isPrefix := true

	for isPrefix {
		linePart, isPrefix, err = reader.ReadLine()
		if err != nil {
			return
		}
		lineBuf.Write(linePart)
	}

	line = append(lineBuf.Bytes(), '\n')
	return
}
