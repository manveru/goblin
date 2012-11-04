// basename - strip file name affixes
//
// SYNOPSIS
//    basename [ -d ] string [ suffix ]
//
// DESCRIPTION
//    Basename deletes any prefix ending in slash (/) and the
//    suffix, if present in string, from string, and prints the
//    result on the standard output.
//
//    The -d option instead prints the directory component, that
//    is, string up to but not including the final slash.  If the
//    string contains no slash, a period and newline are printed.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	argc := len(args)

	if argc < 1 || argc > 3 {
		fmt.Println(os.Stderr, "usage: basename [-d] string [suffix]\n")
		os.Exit(1)
	}

	dflag := false

	if argc > 0 && args[0] == "-d" {
		dflag = true
		args = args[1:]
		argc -= 1
	}

	if dflag {
		fmt.Println(dirname(args[0]))
	} else {
		if argc == 1 {
			fmt.Println(basename(args[0], ""))
		} else if argc == 2 {
			fmt.Println(basename(args[0], args[1]))
		}
	}
}

func basename(path, suffix string) string {
	pathIndex := strings.LastIndex(path, "/")
	if pathIndex != 1 {
		path = path[pathIndex+1:]
	}
	if suffix != "" {
		suffixIndex := strings.LastIndex(path, suffix)
		if len(path)-(suffixIndex+len(suffix)) == 0 {
			path = path[:suffixIndex]
		}

	}
	return path
}

func dirname(path string) string {
	pathIndex := strings.LastIndex(path, "/")
	if pathIndex != -1 {
		return path[:pathIndex]
	}
	return "."
}
