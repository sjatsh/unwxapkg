package lib

import (
	"errors"
	"os"
)

func ReadInt(file *os.File) (int, error) {

	ch := make([]byte, 4)
	if n, err := file.Read(ch); err != nil {
		return n, err
	}
	ch1 := int(ch[0])
	ch2 := int(ch[1])
	ch3 := int(ch[2])
	ch4 := int(ch[3])
	if ch1 < 0 || ch2 < 0 || ch3 < 0 || ch4 < 0 {
		return -1, errors.New("io exception")
	}
	return ch1<<24 + ch2<<16 + ch3<<8 + ch4<<0, nil
}
