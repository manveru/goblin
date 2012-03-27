package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

// ProbablyPrime is good enough for casual use, i've only verified numbers up
// to 9007199254740847 manually, but it should be good beyond that.
// More checks slow the algorithm down quite a bit.
func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "usage: primes starting [ending]")
		os.Exit(1)
	}

	lower, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "primes: %s\n", err)
		os.Exit(0)
	}
	if lower < 3 {
		lower = 3
	}
	if lower%2 == 0 {
		lower += 1
	}

	var upper int64
	stepOnce := false

	if len(os.Args) > 2 {
		upper, err = strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "primes: %s\n", err)
			os.Exit(1)
		}
		if upper < lower {
			return
		}
	} else {
		stepOnce = true
	}

	// increment current by two, we already know it's an uneven number, so let's
	// only check those.
	bTwo := big.NewInt(2)
	bLower := big.NewInt(lower)
	bUpper := big.NewInt(upper)
	current := bLower

	if stepOnce {
		for {
			if current.ProbablyPrime(4) {
				fmt.Println(current.Int64())
				return
			}
			current = current.Add(current, bTwo)
		}
	}

	for current.Cmp(bUpper) < 0 {
		if current.ProbablyPrime(4) {
			fmt.Println(current.Int64())
		}
		current = current.Add(current, bTwo)
	}
}
