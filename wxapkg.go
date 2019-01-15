package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type UnWxapkg struct {
	InPath  string
	OutPath string
	inFile  *os.File
}

func (extract *UnWxapkg) Unwxapkg() error {

	file, err := os.OpenFile(extract.InPath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	extract.inFile = file

	fullName := filepath.Base(file.Name())
	ext := filepath.Ext(file.Name())
	fileName := strings.TrimSuffix(fullName, ext)

	unapkgDir := extract.OutPath + "/" + fileName
	if _, err := os.Stat(unapkgDir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(unapkgDir, os.ModePerm); err != nil {
				return err
			}
		}
	}
	extract.OutPath = unapkgDir

	bom := make([]byte, 1)
	if _, err := file.Read(bom); err != nil {
		return err
	}
	//determine file types
	if bom[0] != byte(0xBE) {
		return errors.New("file type error")
	}
	if _, err := file.Seek(0xE, 0); err != nil {
		return err
	}
	//file count in package
	fileCount, err := ReadInt(file)
	if err != nil {
		return err
	}

	wxApkgItems := make([]WxApkgItem, 0, fileCount)
	for i := 0; i < fileCount; i++ {
		item, err := GetItem(file)
		if err != nil {
			return err
		}
		wxApkgItems = append(wxApkgItems, item)
	}

	for idx := range wxApkgItems {
		item := wxApkgItems[idx]
		fmt.Println(item)
		if err := extract.WriteFile(item); err != nil {
			return err
		}
	}
	return nil
}

func GetItem(file *os.File) (WxApkgItem, error) {

	nameLen, err := ReadInt(file)
	if err != nil {
		return WxApkgItem{}, err
	}
	nameBuf := make([]byte, nameLen)
	if _, err := file.Read(nameBuf); err != nil {
		return WxApkgItem{}, err
	}
	name := string(nameBuf[:])
	start, err := ReadInt(file)
	if err != nil {
		return WxApkgItem{}, err
	}
	length, err := ReadInt(file)
	if err != nil {
		return WxApkgItem{}, err
	}
	item := WxApkgItem{
		Name:   name,
		Start:  start,
		Length: length,
	}
	return item, nil
}

func (extract *UnWxapkg) WriteFile(item WxApkgItem) error {

	dir := filepath.Dir(extract.OutPath + item.Name)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(extract.OutPath + item.Name)
	if err != nil {
		return err
	}
	defer func() {
		f.Sync()
		f.Close()
	}()

	buf := make([]byte, item.Length)
	if _, err := extract.inFile.Read(buf); err != nil {
		return err
	}
	if _, err := f.Write(buf); err != nil {
		return err
	}
	return nil
}

type WxApkgItem struct {
	Name   string `json:"name"`
	Start  int    `json:"start"`
	Length int    `json:"length"`
}
