package pkg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type UnWxapkg struct {
	InPath         string       // 压缩包文件路径
	OutPath        string       // 输出目录
	compressedFile *os.File     // 压缩包文件
	offset         int64        // 文件读取当前偏移量
	headerLength   int64        // 文件头长度
	infoListLength uint32       // header头长度
	dataLength     uint32       // 数据长度
	unZipFileList  []WxApkgItem // 解压文件列表
}

type WxApkgItem struct {
	Name   string `json:"name"`
	Start  int64  `json:"start"`
	Length int64  `json:"length"`
}

func (extract *UnWxapkg) Unwxapkg() error {
	file, err := os.OpenFile(extract.InPath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 压缩文件
	extract.compressedFile = file

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

	// 解析压缩包文件头
	if err := extract.ReadHeader(14); err != nil {
		return err
	}
	// 解析文件列表
	if err := extract.GenFileList(); err != nil {
		return err
	}
	return nil
}

// 获取压缩包header信息
func (extract *UnWxapkg) ReadHeader(headerLength int64) error {
	header := make([]byte, headerLength)
	if _, err := extract.compressedFile.ReadAt(header, 0); err != nil {
		return err
	}
	firstMark := header[0]
	lastMark := header[headerLength-1]
	if firstMark != 0xbe || lastMark != 0xed {
		return errors.New("magic number is not correct")
	}
	var infoListLength uint32
	if err := binary.Read(bytes.NewBuffer(header[5:9]), binary.BigEndian, &infoListLength); err != nil {
		return err
	}
	var dataLength uint32
	if err := binary.Read(bytes.NewBuffer(header[9:headerLength-1]), binary.BigEndian, &dataLength); err != nil {
		return err
	}

	extract.offset = headerLength
	extract.headerLength = headerLength
	extract.infoListLength = infoListLength
	extract.dataLength = dataLength

	fmt.Printf("CompressedFile: %s\nHeaderLength: %d\nInfoListLength: %d\nDataLength: %d\n\n", extract.InPath, extract.headerLength, extract.infoListLength, extract.dataLength)
	return nil
}

// 解析文件列表
func (extract *UnWxapkg) GenFileList() error {
	listInfo := make([]byte, extract.infoListLength)
	if _, err := extract.compressedFile.ReadAt(listInfo, extract.headerLength); err != nil {
		return err
	}
	extract.offset += int64(extract.infoListLength)

	var step uint32 = 4
	var off uint32 = 0
	var fileCount uint32

	if err := binary.Read(bytes.NewBuffer(listInfo[off:step]), binary.BigEndian, &fileCount); err != nil {
		return err
	}
	off += step

	fileList := make([]WxApkgItem, 0)
	for i := uint32(0); i < fileCount; i++ {
		fileInfo := WxApkgItem{}
		var nameLen uint32
		if err := binary.Read(bytes.NewBuffer(listInfo[off:off+step]), binary.BigEndian, &nameLen); err != nil {
			return err
		}
		off += step

		fileInfo.Name = string(listInfo[off : off+nameLen])
		off += nameLen

		var fileOff uint32
		if err := binary.Read(bytes.NewBuffer(listInfo[off:off+step]), binary.BigEndian, &fileOff); err != nil {
			return err
		}
		fileInfo.Start = int64(fileOff)
		off += step

		var size uint32
		if err := binary.Read(bytes.NewBuffer(listInfo[off:off+step]), binary.BigEndian, &size); err != nil {
			return err
		}
		fileInfo.Length = int64(size)
		off += step

		fileList = append(fileList, fileInfo)
	}

	extract.unZipFileList = fileList
	return extract.writeFile()
}

// 写出文件
func (extract *UnWxapkg) writeFile() error {
	for _, item := range extract.unZipFileList {

		filePath := extract.OutPath + item.Name
		dir := filepath.Dir(filePath)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}

		if _, err := os.Stat(filePath); err != nil && !os.IsExist(err) {
			f, err := os.Create(filePath)
			if err != nil {
				return err
			}

			buf := make([]byte, item.Length)
			if _, err := extract.compressedFile.ReadAt(buf, item.Start); err != nil {
				f.Close()
				return err
			}
			if _, err := f.Write(buf); err != nil {
				f.Close()
				return err
			}
			f.Close()
			fmt.Printf("%v\n", item)
		}
	}
	return nil
}
