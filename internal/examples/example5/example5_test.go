// Package example5 demonstrates column generation with custom method configurations
// Shows custom method names and multiple model column generation testing
// Used to showcase configurable column generation options
//
// example5 演示自定义方法配置的列生成
// 展示自定义方法名和多模型列生成测试
// 用于展示可配置的列生成选项
package example5

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen/internal/examples/example5/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
)

// TestExample5 demonstrates column generation with custom receiver names
// Tests both Person and Example models with gofakeit data generation
//
// TestExample5 演示使用自定义接收者名称的列生成
// 测试 Person 和 Example 模型的 gofakeit 数据生成
func TestExample5(t *testing.T) {
	// Create person instance and generate fake data
	// 创建 person 实例并生成虚假数据
	person := &models.Person{}
	require.NoError(t, gofakeit.Struct(&person))
	t.Log(neatjsons.S(person))

	// Get person column mappings
	// 获取 person 列映射
	cls := person.Columns()
	t.Log(neatjsons.S(cls))

	// Create example instance and generate fake data
	// 创建 example 实例并生成虚假数据
	example := &models.Example{}
	require.NoError(t, gofakeit.Struct(&example))
	t.Log(neatjsons.S(example))

	// Get example column mappings
	// 获取 example 列映射
	exampleCls := example.Columns()
	t.Log(neatjsons.S(exampleCls))
}
