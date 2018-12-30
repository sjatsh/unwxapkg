package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	f := flag.String("f", "", "小程序压缩包所在路径")
	out := flag.String("o", ".", "解压后文件路径")

	flag.Parse()

	file, err := os.OpenFile(*f, os.O_RDONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	*out = *out + "/" + file.Name() + "_unpackage"
	_, err = os.Stat(*out)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(*out, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		}
	}

	bom := make([]byte, 1)
	if _, err := file.Read(bom); err != nil {
		log.Fatal(err)
	}
	//determine file types
	if bom[0] != byte(0xBE) {
		log.Fatal("file type error")
	}
	if _, err := file.Seek(0xE, 0); err != nil {
		log.Fatal(err)
	}
	fileCount, err := readInt(file)
	if err != nil {
		log.Fatal(err)
	}

	wxApkgItems := make([]WxApkgItem, 0, fileCount)
	for i := 0; i < fileCount; i++ {

		nameLen, err := readInt(file)
		if err != nil {
			log.Fatal(err)
		}
		nameBuf := make([]byte, nameLen)
		if _, err := file.Read(nameBuf); err != nil {
			log.Fatal(err)
		}
		name := string(nameBuf[:])
		start, err := readInt(file)
		if err != nil {
			log.Fatal(err)
		}
		length, err := readInt(file)
		if err != nil {
			log.Fatal(err)
		}
		item := WxApkgItem{
			Name:   name,
			Start:  start,
			Length: length,
		}
		wxApkgItems = append(wxApkgItems, item)
	}

	for idx := range wxApkgItems {

		item := wxApkgItems[idx]
		fmt.Println(item)
		path := *out + item.Name

		writeFile(file, item, path)
	}
}

func writeFile(file *os.File, item WxApkgItem, path string) {

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, item.Length)
	if _, err := file.Read(buf); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(buf); err != nil {
		log.Fatal(err)
	}
}

func readInt(file *os.File) (int, error) {

	ch := make([]byte, 4)
	if n, err := file.Read(ch); err != nil {
		return n, err
	}
	ch1 := int(ch[0])
	ch2 := int(ch[1])
	ch3 := int(ch[2])
	ch4 := int(ch[3])
	if ch1 < 0 || ch2 < 0 || ch3 < 0 || ch4 < 0 {
		return 0, errors.New("io exception")
	}
	return ch1<<24 + ch2<<16 + ch3<<8 + ch4<<0, nil
}

type WxApkgItem struct {
	Name   string `json:"name"`
	Start  int    `json:"start"`
	Length int    `json:"length"`
}
