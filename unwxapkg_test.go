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

func TestMkdirError(t *testing.T) {
	Unwxapkg("dest/102.wxapkg", "./|$")
}

func TestReadFileError(t *testing.T) {
	err := Unwxapkg("dest/xx.wxapkg", ".")
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") &&
		!strings.Contains(err.Error(), "The system cannot find the file") {
		t.Fatal(err)
	}
}

func TestErrBom(t *testing.T) {
	f, err := os.Create("dest/test1.wxapkg")
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	err = Unwxapkg("dest/test1.wxapkg", ".")
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		t.Fatal(err)
	}
}

func TestFileTypeError(t *testing.T) {

	f, err := os.Create("dest/test2.wxapkg")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write([]byte{0XBF}); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	err = Unwxapkg("dest/test2.wxapkg", ".")
	if err != nil && !strings.Contains(err.Error(), "file type error") {
		t.Fatal(err)
	}
}

func TestSeekError(t *testing.T) {

	f, err := os.Create("dest/test3.wxapkg")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write([]byte{0XBE}); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	err = Unwxapkg("dest/test3.wxapkg", ".")
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		t.Fatal(err)
	}
}
