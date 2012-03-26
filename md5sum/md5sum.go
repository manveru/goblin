package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

const (
	BUFFER_SIZE = 16384
)

func main() {
	if len(os.Args) == 1 {
		sum(os.Stdin, "")
	} else {
		for _, arg := range os.Args[1:] {
			if arg == "-h" || arg == "--help" {
				fmt.Println("usage: md5sum [file...]")
				os.Exit(0)
			}

			file, err := os.Open(arg)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("md5sum: %s\n", err))
				continue
			}

			sum(file, arg)
			file.Close()
		}
	}
}

func sum(file *os.File, name string) {
	hash := md5.New()
	buf := make([]byte, BUFFER_SIZE)
	for {
		n, err := file.Read(buf)
		if err == nil {
			hash.Write(buf[0:n])
		} else if n == 0 && err == io.EOF {
			break
		} else {
			os.Stderr.WriteString(fmt.Sprintf("md5sum: %s\n", err))
			os.Exit(1)
		}
	}

	if name == "" {
		fmt.Printf("%x\n", hash.Sum(nil))
	} else {
		fmt.Printf("%x\t%s\n", hash.Sum(nil), name)
	}
}
