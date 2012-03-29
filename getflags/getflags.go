package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	flagfmt := os.Getenv("flagfmt")
	defer func() {
		recover()
		fmt.Println("status=usage")
	}()

	if flagfmt == "" {
		fmt.Fprintln(os.Stderr, "$flagfmt is not set")
		os.Exit(0)
	}

	flagBool := map[uint8]*bool{}
	flagArg := map[uint8]*string{}

	flags := strings.Split(flagfmt, ",")
	for _, flag := range flags {
		parts := strings.SplitN(flag, " ", 2)
		if len(parts) == 1 {
			flagBool[parts[0][0]] = nil
		} else {
			flagArg[parts[0][0]] = nil
		}
	}

	t := true
	var passArgs []string

	for n := 1; n < len(os.Args); n++ {
		arg := os.Args[n]
		if arg[0] == '-' {
			flagName := arg[1]
			_, ok := flagBool[flagName]
			if ok {
				flagBool[flagName] = &t
			} else {
				_, ok := flagArg[flagName]
				if ok {
					flagArg[flagName] = &os.Args[n+1]
					n++
				}
			}
		} else {
			passArgs = os.Args[n:]
			break
		}
	}

	for name, value := range flagBool {
		if value == nil || !*value {
			fmt.Printf("flag%c=()\n", name)
		} else {
			fmt.Printf("flag%c=1\n", name)
		}
	}

	for name, value := range flagArg {
		if value == nil {
			fmt.Printf("flag%c=()\n", name)
		} else {
			fmt.Printf("flag%c=(%s)\n", name, escape(*value))
		}
	}

	passArgsString := []string{}
	for _, arg := range passArgs {
		passArgsString = append(passArgsString, escape(arg))
	}
	fmt.Printf("*=(%s)\n", strings.Join(passArgsString, " "))
	fmt.Println("status=''")
}

func escape(s string) string {
	out := make([]rune, len(s))

	gotSpaces := false
	for _, r := range s {
		switch r {
		case '\'':
			out = append(out, '\\', '\'')
		case ' ':
			gotSpaces = true
			out = append(out, ' ')
		default:
			out = append(out, r)
		}
	}

	if gotSpaces {
		return "'" + string(out) + "'"
	}
	return string(out)
}
