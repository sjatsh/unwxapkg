package main

import (
	"flag"
	"log"
)

var f = flag.String("f", "", "wechat wxapkg file path")
var out = flag.String("o", ".", "output file path")

func main() {

	flag.Parse()

	extract := new(UnWxapkg)
	extract.InPath = *f
	extract.OutPath = *out
	if err := extract.Unwxapkg(); err != nil {
		log.Fatal(err)
	}
}
