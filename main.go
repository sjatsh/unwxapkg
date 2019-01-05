package main

import (
	"flag"
	"log"
)

var f = flag.String("f", "", "wechat wxapkg file path")
var out = flag.String("o", ".", "output file path")

func main() {

	flag.Parse()

	if err := Unwxapkg(*f, *out); err != nil {
		log.Fatal(err)
	}
}
