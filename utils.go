// Package gormcngen provides position tracking utilities for code generation tracing
// Auto captures source file and function information using runtime package
// Supports URL-encoded file paths and stack frame navigation
//
// gormcngen 提供代码生成追踪的位置追踪工具
// 使用 runtime 包自动捕获源文件和函数信息
// 支持 URL 编码的文件路径和调用栈帧导航
package gormcngen

import (
	"fmt"
	"net/url"
	"path/filepath"
	"runtime"

	"github.com/yyle88/rese"
)

// GetGenPosFuncMark gets position information for code generation tracing
// Returns a formatted string showing the source file and function that triggered generation
// Uses runtime.Caller to walk up the stack using the specified frame count
// Supports URL-encoded file paths
//
// GetGenPosFuncMark 获取调用者位置信息，用于代码生成的可追溯性
// 返回格式化字符串，显示触发生成的源文件和函数
// 使用 runtime.Caller 按指定帧数向上遍历调用栈
// 支持URL编码的文件路径
func GetGenPosFuncMark(skip int) string {
	// Get caller position info using runtime package
	// 使用 runtime 包获取调用者位置信息
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	// Extract just the filename, not the complete path for clean output
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
