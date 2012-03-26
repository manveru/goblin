package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	octalBase      = flag.Bool("o", false, "octal base")
	decimalBase    = flag.Bool("d", false, "decimal base")
	hexBase        = flag.Bool("x", false, "hexadecimal base")
	customBase     = flag.Int("b", -1, "[n] base (2..36)")
	forceNumeric   = flag.Bool("n", false, "Force numeric output")
	forceCharacter = flag.Bool("c", false, "Force character output")
	runningText    = flag.Bool("t", false, "Convert from numbers to running text; do not interpret control characters or insert newlines")
	isoExtensions  = flag.Bool("8", false, "Include ISO Latin-1 extensions (codes 0200-0277)")

	nchars = 128
	base   = 16
	ncol   = 0
	text   = 1
	strip  = false

	charTable = [256]string{
		"nul", "soh", "stx", "etx", "eot", "enq", "ack", "bel",
		"bs ", "ht ", "nl ", "vt ", "np ", "cr ", "so ", "si ",
		"dle", "dc1", "dc2", "dc3", "dc4", "nak", "syn", "etb",
		"can", "em ", "sub", "esc", "fs ", "gs ", "rs ", "us ",
		"sp ", " ! ", ` " `, " # ", " $ ", " % ", " & ", " ' ",
		" ( ", " ) ", " * ", " + ", " , ", " - ", " . ", " / ",
		" 0 ", " 1 ", " 2 ", " 3 ", " 4 ", " 5 ", " 6 ", " 7 ",
		" 8 ", " 9 ", " : ", " ; ", " < ", " = ", " > ", " ? ",
		" @ ", " A ", " B ", " C ", " D ", " E ", " F ", " G ",
		" H ", " I ", " J ", " K ", " L ", " M ", " N ", " O ",
		" P ", " Q ", " R ", " S ", " T ", " U ", " V ", " W ",
		" X ", " Y ", " Z ", " [ ", ` \ `, " ] ", " ^ ", " _ ",
		" ` ", " a ", " b ", " c ", " d ", " e ", " f ", " g ",
		" h ", " i ", " j ", " k ", " l ", " m ", " n ", " o ",
		" p ", " q ", " r ", " s ", " t ", " u ", " v ", " w ",
		" x ", " y ", " z ", " { ", " | ", " } ", " ~ ", "del",
		"x80", "x81", "x82", "x83", "x84", "x85", "x86", "x87",
		"x88", "x89", "x8a", "x8b", "x8c", "x8d", "x8e", "x8f",
		"x90", "x91", "x92", "x93", "x94", "x95", "x96", "x97",
		"x98", "x99", "x9a", "x9b", "x9c", "x9d", "x9e", "x9f",
		"xa0", " ¡ ", " ¢ ", " £ ", " ¤ ", " ¥ ", " ¦ ", " § ",
		" ¨ ", " © ", " ª ", " « ", " ¬ ", " ­ ", " ® ", " ¯ ",
		" ° ", " ± ", " ² ", " ³ ", " ´ ", " µ ", " ¶ ", " · ",
		" ¸ ", " ¹ ", " º ", " » ", " ¼ ", " ½ ", " ¾ ", " ¿ ",
		" À ", " Á ", " Â ", " Ã ", " Ä ", " Å ", " Æ ", " Ç ",
		" È ", " É ", " Ê ", " Ë ", " Ì ", " Í ", " Î ", " Ï ",
		" Ð ", " Ñ ", " Ò ", " Ó ", " Ô ", " Õ ", " Ö ", " × ",
		" Ø ", " Ù ", " Ú ", " Û ", " Ü ", " Ý ", " Þ ", " ß ",
		" à ", " á ", " â ", " ã ", " ä ", " å ", " æ ", " ç ",
		" è ", " é ", " ê ", " ë ", " ì ", " í ", " î ", " ï ",
		" ð ", " ñ ", " ò ", " ó ", " ô ", " õ ", " ö ", " ÷ ",
		" ø ", " ù ", " ú ", " û ", " ü ", " ý ", " þ ", " ÿ ",
	}
)

func usage() {
	fmt.Fprintln(os.Stdout, "usage: ascii [-8] [-xod | -b8] [-ncst] [text]")
}

func main() {
	flag.Parse()

	if *isoExtensions {
		nchars = 256
	}
	if *hexBase {
		base = 16
	} else if *octalBase {
		base = 8
	} else if *decimalBase {
		base = 10
	} else if *customBase >= 2 && *customBase <= 36 {
		base = *customBase
	} else if *customBase != -1 {
		flag.Usage()
		os.Exit(1)
	}

	if *forceNumeric {
		text = 0
	}
	if *runningText {
		strip = true
	}
	if *runningText || *forceCharacter {
		text = 2
	}

	switch base {
	case 2:
		ncol = 7
	case 3:
		ncol = 5
	case 4, 5:
		ncol = 4
	case 6, 7, 8, 9, 10, 11:
		ncol = 3
	default:
		ncol = 2
	}

	if flag.NArg() == 0 {
		printTable(os.Stdout, ncol)
	} else {
		if text == 1 {
			_, err := strconv.ParseInt(flag.Arg(0), base, 16)
			if err != nil {
				text = 0
			}
		}
		if text == 1 {
			for _, arg := range flag.Args() {
				printText(os.Stdout, base, arg)
			}
		} else {
			for _, arg := range flag.Args() {
				printNum(os.Stdout, base, arg)
			}
		}
	}
}

func printTable(w io.Writer, ncol int) {
	format := fmt.Sprintf("|%%0%ds %%s", ncol)
	for i := 0; i < nchars; i++ {
		based := strconv.FormatInt(int64(i), base)
		fmt.Fprintf(w, format, based, charTable[i])
		if (i & 7) == 7 {
			fmt.Fprintln(w, "|")
		}
	}
}

func printText(w io.Writer, base int, arg string) {
	num, err := strconv.ParseInt(arg, base, 16)
	if err != nil {
		panic(err)
	}
	if strip {
		fmt.Fprintf(w, "%c\n", num&0377)
	} else {
		fmt.Fprintf(w, "%s\n", charTable[num&0377])
	}
}

func printNum(w io.Writer, base int, arg string) {
	for _, r := range arg {
		fmt.Fprintln(w, strconv.FormatInt(int64(r), base))
	}
}
