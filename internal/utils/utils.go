package utils

import (
	"os"
	"unicode"
)

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func ConvertToUnexportable(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToLower(runes[0])
	}
	return string(runes)
}

func IsExportable(name string) bool {
	return unicode.IsUpper(([]rune(name))[0])
}

func SwitchToggleExportable(name string) string {
	runes := []rune(name)
	if unicode.IsUpper(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	} else if unicode.IsLower(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}
