package models

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerate(t *testing.T) {
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable column class names like ExampleColumns // 生成可导出的列类名称如 ExampleColumns
		WithColumnsMethodRecvName("c").  // Set receiver name for column methods // 设置列方法的接收器名称
		WithColumnsCheckFieldType(true)  // Enable field type checking for type safe // 启用字段类型检查以获得更好的类型安全

	cfg := gormcngen.NewConfigs([]interface{}{&Person{}, &Example{}}, options, absPath)
	cfg.Gen()
}
