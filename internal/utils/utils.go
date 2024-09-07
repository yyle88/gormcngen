package utils

import (
	"os"
	"unicode"

	"github.com/yyle88/zaplog"
)

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func MustFile(path string) {
	info, err := os.Stat(path)
	AssertTRUE(!os.IsNotExist(err) && info != nil && !info.IsDir()) //这是简化版的就不要考虑其它错误啦
}

func AssertTRUE(v bool) bool {
	if !v {
		zaplog.ZAPS.P1.LOG.Panic("B IS FALSE")
	}
	return v
}

func GetMapKeys[K comparable, V any](m map[K]V) (ks []K) {
	for k := range m {
		ks = append(ks, k)
	}
	return ks //返回默认值比如0或者空字符串等
}

func C0ToLOWER(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToLower(runes[0])
	}
	return string(runes)
}
