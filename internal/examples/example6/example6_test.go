// Package example6 demonstrates enhanced column generation with TableColumns method
// Shows table decoration support and advanced column generation features
// Used to showcase isGenFuncTableColumns option and extended features
//
// example6 演示带有 TableColumns 方法的增强列生成
// 展示表装饰支持和高级列生成特性
// 用于展示 isGenFuncTableColumns 选项和扩展功能
package example6

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen/internal/examples/example6/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
)

// TestExample6 demonstrates enhanced column generation with table decorations
// Tests both Person and Example models with TableColumns method support
//
// TestExample6 演示使用表装饰的增强列生成
// 测试支持 TableColumns 方法的 Person 和 Example 模型
func TestExample6(t *testing.T) {
	// Create person instance and generate fake data
	// 创建 person 实例并生成虚假数据
	person := &models.Person{}
	require.NoError(t, gofakeit.Struct(&person))
	t.Log(neatjsons.S(person))

	// Get person column mappings with table decoration support
	// 获取支持表装饰的 person 列映射
	cls := person.Columns()
	t.Log(neatjsons.S(cls))

	// Create example instance and generate fake data
	// 创建 example 实例并生成虚假数据
	example := &models.Example{}
	require.NoError(t, gofakeit.Struct(&example))
	t.Log(neatjsons.S(example))

	// Get example column mappings with table decoration support
	// 获取支持表装饰的 example 列映射
	exampleCls := example.Columns()
	t.Log(neatjsons.S(exampleCls))
}
