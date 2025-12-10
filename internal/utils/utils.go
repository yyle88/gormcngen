// Package utils: Internal utilities for gormcngen code generation
// Provides functions for string manipulation and exportable status conversion
// Supports case conversion and status toggles for generated identifiers
//
// utils: gormcngen 代码生成的内部工具函数
// 提供字符串操作和导出性转换的辅助函数
// 支持生成标识符的大小写转换和可见性切换
package utils

import (
	"fmt"
	"net/url"
	"path/filepath"
	"runtime"
	"unicode"

	"github.com/yyle88/rese"
)

// ConvertToUnexportable converts the first string character to lowercase
// Transforms exported identifiers to unexported ones for internal use
// Returns the string with the first rune converted to lowercase
//
// ConvertToUnexportable 将字符串的第一个字符转换为小写
// 将导出标识符转换为未导出标识符供内部使用
// 返回第一个 rune 转换为小写的字符串
func ConvertToUnexportable(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToLower(runes[0])
	}
	return string(runes)
}

// IsExportable checks if an identifier is exportable (starts with uppercase)
// Returns true if the first character is uppercase, false otherwise
// Used to determine status of generated structs and methods
//
// IsExportable 检查标识符是否可导出（以大写字母开头）
// 如果第一个字符是大写返回 true，否则返回 false
// 用于确定生成的结构体和方法的可见性
func IsExportable(name string) bool {
	return unicode.IsUpper(([]rune(name))[0])
}

// SwitchToggleExportable toggles the exportable status of an identifier
// Converts uppercase first character to lowercase and vice versa
// Used for finding alternative struct names during AST analysis
//
// SwitchToggleExportable 切换标识符的可导出性
// 将大写第一个字符转换为小写，反之亦然
// 用于在 AST 分析过程中查找替代结构体名称
func SwitchToggleExportable(name string) string {
	runes := []rune(name)
	if unicode.IsUpper(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	} else if unicode.IsLower(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

// GetGenPosFuncMark gets position information used in code generation tracing
// Returns a formatted string showing the source file and function that triggered generation
// Uses runtime.Caller to walk up the stack using the specified frame count
// Supports URL-encoded file paths
//
// GetGenPosFuncMark 获取调用者位置信息，用于代码生成的可追溯性
// 返回格式化字符串，显示触发生成的源文件和函数
// 使用 runtime.Caller 按指定帧数向上遍历调用栈
// 支持URL编码的文件路径
func GetGenPosFuncMark(skip int) string {
	// Get invocation position info using runtime package
	// 使用 runtime 包获取调用者位置信息
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	// Extract just the filename, not the complete path, to produce clean output
	// 为了可读性，只提取文件名而不是完整路径
	filename := filepath.Base(file)

	// Decode URL-encoded filename
	// 解码URL编码的文件名
	filename = rese.C1(url.PathUnescape(filename))

	funcInfo := runtime.FuncForPC(pc)
	if funcInfo == nil {
		return filename
	}

	// Extract just the function name, not the complete package path
	// 只提取函数名，不包含完整的包路径
	funcName := filepath.Base(funcInfo.Name())
	return fmt.Sprintf("%s:%d -> %s", filename, line, funcName)
}
