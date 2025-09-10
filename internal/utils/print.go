// Package utils: Print utilities for gormcngen debugging and output
// Provides convenient printing functions for development and testing
// Wraps printgo functions with gormcngen-specific configurations
//
// utils: gormcngen 调试和输出的打印工具
// 为开发和测试提供便利的打印函数
// 封装 printgo 功能，带有 gormcngen 特定配置
package utils

import (
	"github.com/yyle88/printgo"
)

// NewPTX creates a new print context for gormcngen debugging output
// Returns a configured printgo.PTX instance for structured printing
// Used throughout gormcngen for consistent debug and trace output
//
// NewPTX 为 gormcngen 调试输出创建新的打印上下文
// 返回一个配置好的 printgo.PTX 实例用于结构化打印
// 在 gormcngen 中用于一致的调试和跟踪输出
func NewPTX() *printgo.PTX {
	return printgo.NewPTX()
}
