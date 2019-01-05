package main

import (
	"os"
	"strings"
	"testing"
)

func TestUnwxapkg(t *testing.T) {
	if err := Unwxapkg("dest/102.wxapkg", "."); err != nil {
		t.Fatal(err)
	}
}

func TestReadFileError(t *testing.T) {
	err := Unwxapkg("dest/xx.wxapkg", ".")
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") &&
		!strings.Contains(err.Error(), "The system cannot find the file") {
		t.Fatal(err)
	}
}

func TestErrBom(t *testing.T) {
	f, err := os.Create("dest/test.wxapkg")
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	err = Unwxapkg("dest/test.wxapkg", ".")
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		t.Fatal(err)
	}
}
