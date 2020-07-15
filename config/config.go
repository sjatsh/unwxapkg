// Author: SunJun <i@sjis.me>
package config

import (
	"math"
	"regexp"
	"strings"
)

func toDir(to, from string) string {
	if from[0] == '.' {
		from = from[1:]
	}
	if to[0] == '.' {
		to = to[1:]
	}
	from = string(regexp.MustCompile("/\\/g").ReplaceAll([]byte(from), []byte{'/'}))
	to = string(regexp.MustCompile("/\\/g").ReplaceAll([]byte(to), []byte{'/'}))
	a := int(math.Min(float64(len(from)), float64(len(to))))
	for i, m := 1, a; i <= m; i++ {
		if !strings.HasPrefix(to, from[0:i]) {
			a = i - 1
			break
		}
	}
	pub := from[0:a]
	length := strings.LastIndex(pub, "/") + 1
	k := from[length:]
	ret := ""
	for i := 0; i < len(k); i++ {
		if k[i] == '/' {
			ret += "../"
		}
	}
	return ret + to[length:]
}
