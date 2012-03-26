// prints number and its prime factors.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var wheel = [48]float64{
	2, 10, 2, 4, 2, 4, 6, 2, 6, 4, 2, 4, 6, 6, 2, 6, 4, 2, 6, 4, 6, 8, 4, 2, 4,
	2, 4, 8, 6, 4, 6, 2, 4, 6, 2, 6, 6, 4, 2, 4, 6, 2, 6, 4, 2, 4, 2, 10,
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			n, _ := strconv.ParseFloat(arg, 10)
			factor(os.Stdout, n)
		}
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := readLine(reader)
		n, err := strconv.ParseFloat(line, 10)
		if err != nil {
			break
		}
		factor(os.Stdout, n)
	}
}

func factor(w io.Writer, n float64) {
	fmt.Fprintf(w, "%.0f\n", n)

	if n == 0.0 {
		return
	}

	var integral, fractional float64
	s := math.Sqrt(n) + 1.0

	for modf(n/2.0, &integral, &fractional) {
		fmt.Fprintln(w, "     2")
		n = integral
		s = math.Sqrt(n) + 1.0
	}

	for modf(n/3.0, &integral, &fractional) {
		fmt.Fprintln(w, "     3")
		n = integral
		s = math.Sqrt(n) + 1.0
	}

	for modf(n/5.0, &integral, &fractional) {
		fmt.Fprintln(w, "     5")
		n = integral
		s = math.Sqrt(n) + 1.0
	}

	for modf(n/7.0, &integral, &fractional) {
		fmt.Fprintln(w, "     7")
		n = integral
		s = math.Sqrt(n) + 1.0
	}

	var d float64 = 1
	var i int = 1
	for {
		d += wheel[i]
		for modf(n/d, &integral, &fractional) {
			fmt.Fprintf(w, "     %.0f\n", d)
			n = integral
			s = math.Sqrt(n) + 1.0
		}
		i++
		if i >= len(wheel) {
			i = 0
			if d > s {
				break
			}
		}
	}

	if n > 1 {
		fmt.Fprintf(w, "     %.0f\n", n)
	}
	fmt.Fprintln(w)
}

func modf(n float64, integral, fractional *float64) bool {
	*integral, *fractional = math.Modf(n)
	return *fractional == 0.0
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
