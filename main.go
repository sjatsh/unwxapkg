package main

import (
	"flag"
	"unwxapkg/lib"
)

func main() {

	f := flag.String("f", "", "wechat wxapkg file path")
	out := flag.String("o", ".", "output file path")
	flag.Parse()

	lib.Unwxapkg(*f, *out)
}
