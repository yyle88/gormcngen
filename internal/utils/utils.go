package utils

import (
	"encoding/json"
	"os"
	"unicode"

	"github.com/pkg/errors"
)

func Neat(v interface{}) string {
	data, err := NeatBytes(v)
	if err != nil {
		return "" //when the result is empty string, means wrong
	}
	return string(data)
}

func NeatBytes(v interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, errors.WithMessage(err, "marshal object is wrong")
	}
	return data, nil
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func IsFileExist(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && info != nil && !info.IsDir() //这是简化版的就不要考虑其它错误啦
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
