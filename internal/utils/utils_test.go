// Package gormcngen tests validate position tracking utilities
// Auto verifies GetGenPosFuncMark function with runtime invocation info
//
// gormcngen 测试包验证位置追踪工具
// 自动验证 GetGenPosFuncMark 函数与运行时调用者信息
package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGenPosFuncMark(t *testing.T) {
	result := GetGenPosFuncMark(0)
	require.NotEmpty(t, result)
	t.Log("result:", result)
}
