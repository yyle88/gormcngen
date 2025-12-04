package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

// TestGenerate triggers column code generation for demo2 models
// Allows integration with go generate ./... for build automation
//
// TestGenerate 为 demo2 模型触发列代码生成
// 支持与 go generate ./... 集成以实现构建自动化
//
//go:generate go test -v -run TestGenerate
func TestGenerate(t *testing.T) {
	// Get source file path from test file location // 从测试文件位置获取源文件路径
	absPath := runtestpath.SrcPath(t)
	t.Log(absPath)

	// Confirm target file exists // 确认目标文件存在
	require.True(t, osmustexist.IsFile(absPath))

	// Define models to generate columns for // 定义要生成列的模型
	objects := []any{&Account{}, &Purchase{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true) // Exportable class names like AccountColumns // 可导出的类名如 AccountColumns

	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.WithIsGenPreventEdit(false) // Disable prevent-edit headers // 禁用防编辑头
	cfg.Gen()                       // Write generated code to target file // 将生成的代码写入目标文件
}
