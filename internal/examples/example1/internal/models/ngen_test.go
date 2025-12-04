package models

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

// TestGenerate triggers column code generation using gormcngen
// Configures generation options and outputs type-safe column methods to ngen.go
//
// TestGenerate 使用 gormcngen 触发列代码生成
// 配置生成选项并将类型安全的列方法输出到 ngen.go
func TestGenerate(t *testing.T) {
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable column class names like ExampleColumns // 生成可导出的列类名称如 ExampleColumns
		WithColumnsMethodRecvName("c").  // Set method argument name // 设置方法接收者名称
		WithColumnsCheckFieldType(true)  // Enable field type checking // 启用字段类型检查

	cfg := gormcngen.NewConfigs([]interface{}{&Person{}, &Example{}}, options, absPath)
	cfg.Gen()
}
