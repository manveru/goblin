package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	f, err := os.Open("/opt/plan9/lib/fortunes")
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	buf := bufio.NewReader(f)

	var n float64 = 1
	var keep string
	for {
		x := rand.Float64() * n
		if x < 1 {
			keep, err = readLine(buf)
		} else {
			_, err = readLine(buf)
		}

		if err != nil {
			break
		}
		n += 1
	}

	fmt.Println(keep)
}

func readLine(reader *bufio.Reader) (line string, err error) {
	lineBuf := new(bytes.Buffer)
	var linePart []byte
	isPrefix := true

	for isPrefix {
		linePart, isPrefix, err = reader.ReadLine()
		if err != nil {
			return
		}
		lineBuf.Write(linePart)
	}

	line = lineBuf.String()
	return
}
