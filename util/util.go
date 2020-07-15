// Author: SunJun <i@sjis.me>
package util

import (
	"math"
	"os"
	"regexp"
	"strings"
)

func ExistFile(dir, fileName string) bool {
	_, err := os.Stat(dir + "/" + fileName)
	return err == nil
}

func CommonDir(pathA, pathB string) string {
	if pathA[0] == '.' {
		pathA = pathA[1:]
	}
	if pathB[0] == '.' {
		pathB = pathB[1:]
	}
	r := regexp.MustCompile("/\\/g")
	pathA = string(r.ReplaceAll([]byte(pathA), []byte("/")))
	pathB = string(r.ReplaceAll([]byte(pathB), []byte("/")))

	min := int(math.Min(float64(len(pathA)), float64(len(pathB))))
	for i, m := 1, min; i <= m; i++ {
		if !strings.HasPrefix(pathA, pathB[0:i]) {
			min = i - 1
			break
		}
	}

	pub := pathB[0:min]
	return pathA[0 : strings.LastIndex(pub, "/")+1]
}
