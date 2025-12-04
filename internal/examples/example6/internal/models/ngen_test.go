package models

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

// TestGenerate triggers column code generation using gormcngen
// Configures generation options with TableColumns method support
//
// TestGenerate 使用 gormcngen 触发列代码生成
// 配置生成选项以支持 TableColumns 方法
func TestGenerate(t *testing.T) {
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true).  // Exportable class names like PersonColumns // 可导出的类名如 PersonColumns
		WithColumnsMethodRecvName("one"). // Custom method argument name // 自定义方法接收者名称
		WithColumnsCheckFieldType(true).  // Enable field type checking // 启用字段类型检查
		WithIsGenFuncTableColumns(true)   // Generate TableColumns method // 生成 TableColumns 方法

	cfg := gormcngen.NewConfigs([]interface{}{&Person{}, &Example{}}, options, absPath)
	cfg.Gen()
}
