package gormcngen

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// GetGenPosFuncMark gets position information for code generation tracing
// Returns a formatted string showing the source file and function that triggered generation
// Uses runtime.Caller to walk up the stack using the specified frame count
// Returns empty string if position information cannot be determined
//
// GetGenPosFuncMark 获取调用者位置信息，用于代码生成的可追溯性
// 返回格式化字符串，显示触发生成的源文件和函数
// 使用 runtime.Caller 按指定帧数向上遍历调用栈
// 如果无法确定调用者信息则返回空字符串
func GetGenPosFuncMark(skip int) string {
	// Get caller position info using runtime package
	// 使用 runtime 包获取调用者位置信息
	pc, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}

	// Extract just the filename, not the full path for clean output
	// 为了可读性，只提取文件名而不是完整路径
	filename := filepath.Base(file)

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return filename
	}

	// Extract just the function name, not the full package path
	// 只提取函数名，不包含完整的包路径
	funcName := filepath.Base(fn.Name())
	return fmt.Sprintf("%s:%d -> %s", filename, line, funcName)
}
