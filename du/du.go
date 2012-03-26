/*
Du gives the number of Kbytes allocated to data blocks of named files and,
recursively, of files in named directories. It assumes storage is quantized in
units of 1024 bytes (Kbytes) by default. Other values can be set by the –b
option; size is the number of bytes, optionally suffixed k to specify
multiplication by 1024. If file is missing, the current directory is used. The
count for a directory includes the counts of the contained files and
directories.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	allFiles            = flag.Bool("a", false, "print number of blocks for every file in a directory. Normally counts are printed only for contained directories.")
	scientificNotation  = flag.Bool("e", false, "The –e option causes du to print values (sizes, times or QID paths) in scientific notation")
	ignoreErrors        = flag.Bool("f", false, "The –f option suppresses the printing of warning messages.")
	humanSizes          = flag.Bool("h", false, "The –h option causes du to print values (sizes, times or QID paths) in scientific notation, scaled to less than 1024, and with a suitable SI prefix (e.g., G for binary gigabytes).")
	allFilesAndUseBytes = flag.Bool("n", false, "The –n option prints the size in bytes and the name of each file; it sets –a.")
	SISizes             = flag.Bool("p", false, "The –p option causes du to print values (sizes, times or QID paths) in units of SI–prefix. Case is ignored when looking up SI–prefix. An empty SI–prefix corresponds to a scale factor of 1 (e.g., print sizes in bytes).")
	printQIDPath        = flag.Bool("q", false, "The –q option prints, in the format of du –n, the QID path of each file rather than the size.")
	onlyTopLevel        = flag.Bool("s", false, "The –s option causes du to descend the hierarchy as always, but to print only a summary line for each file.")
	printModifiedTime   = flag.Bool("t", false, "The –t option prints, in the format of du –n, the modified time of each file rather than the size. If the options –tu are specified then the accessed time is printed.")
	printAccessedTime   = flag.Bool("u", false, "only in combination with -t")
)

func main() {
	flag.Parse()

	format := "%d\t%s\n"

	if flag.NArg() == 0 {
		du(format, ".")
	} else {
		for _, arg := range flag.Args() {
			du(format, arg)
		}
	}
}

func du(format string, root string) {
	stat, err := os.Lstat(root)
	if err != nil {
		warn(err)
		return
	}

	if stat.IsDir() {
		err = filepath.Walk(root,
			func(path string, info os.FileInfo, e error) (err error) {
				size := (info.Size() + 1024 - 1) / 1024
				fmt.Fprintf(os.Stdout, "%d\t%s\n", size, path)
				return
			})
		fmt.Println("err:", err)
	}
}

func warn(err error) {
	if !*ignoreErrors {
		fmt.Fprintf(os.Stderr, "du: %s\n", err)
	}
	return
}
