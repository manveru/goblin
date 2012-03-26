package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	flagNoCreate = flag.Bool("c", false, "don't create files")
	flagTime     = flag.Int64("t", time.Now().Unix(), "set mtime to this epoch")
)

func main() {
	flag.Parse()

	mtime := time.Unix(*flagTime, 0)

	for _, arg := range flag.Args() {
		err := os.Chtimes(arg, mtime, mtime)

		if err == nil {
			continue
		}

		if *flagNoCreate {
			reportPathError(err)
			continue
		}

		f, err := os.Create(arg)
		if err != nil {
			reportPathError(err)
			continue
		}

		err = os.Chtimes(arg, mtime, mtime)
		f.Close()
		if err != nil {
			reportPathError(err)
			continue
		}
	}
}

func reportPathError(err error) {
	p := err.(*os.PathError)
	fmt.Fprintf(os.Stderr, "touch: %s: cannot %s: %s\n", p.Path, p.Op, p.Err)
}
