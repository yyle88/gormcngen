// Package example1 demonstrates basic model column generation and usage
// Shows simple column mapping creation and fake data testing with gofakeit
// Used as entry point example for understanding gormcngen basics
//
// example1 演示基本的模型列生成和使用
// 展示简单列映射创建和使用 gofakeit 的虚假数据测试
// 作为理解 gormcngen 基础的入门示例
package example1

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen/internal/examples/example1/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
)

// TestExample1 demonstrates basic model column generation and usage
// Shows how to use gofakeit to generate test data and validate column mappings
//
// TestExample1 演示基本的模型列生成和使用
// 展示如何使用 gofakeit 生成测试数据并验证列映射
func TestExample1(t *testing.T) {
	// Create new person instance and populate with fake data
	// 创建新的 person 实例并填充虚假数据
	person := &models.Person{}
	require.NoError(t, gofakeit.Struct(&person))
	t.Log(neatjsons.S(person))

	// Get column mappings and log them
	// 获取列映射并记录它们
	cls := person.Columns()
	t.Log(neatjsons.S(cls))
}
