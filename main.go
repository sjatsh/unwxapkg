package main

import (
	"flag"
)

func main() {

	f := flag.String("f", "", "wechat wxapkg file path")
	out := flag.String("o", ".", "output file path")
	flag.Parse()

	Unwxapkg(*f, *out)
}
