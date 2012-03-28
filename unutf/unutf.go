package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	for {
		char, _, err := buf.ReadRune()
		if err != nil {
			return
		}
		fmt.Printf("0x%x\n", char)
	}
}
