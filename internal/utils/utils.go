package utils

import (
	"os"
	"unicode"
)

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
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
