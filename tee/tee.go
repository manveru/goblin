package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
)

var (
	flagIgnoreInterrupts = flag.Bool("i", false, "ignore interrupts")
	flagAppend           = flag.Bool("a", false, "append output")
)

func main() {
	flag.Parse()

	if *flagIgnoreInterrupts {
		go func() {
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, os.Interrupt)
			<-sigChan
			os.Exit(1)
		}()
	}

	src := os.Stdin
	dsts := make([]io.Writer, 0)

	for _, arg := range flag.Args() {
		var dst *os.File
		var err error

		if *flagAppend {
			dst, err = os.Open(arg)
			if err != nil {
				dst, err = os.Create(arg)
			} else {
				dst.Seek(0, 2)
			}
		} else {
			dst, err = os.Create(arg)
		}

		if err == nil {
			dsts = append(dsts, dst)
		} else {
			fmt.Fprintln(os.Stderr, "tee:", err)
		}
	}

	writer := io.MultiWriter(dsts...)
	io.Copy(writer, src)
}
