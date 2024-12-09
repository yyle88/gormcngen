package example5

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
		WithExportGeneratedStruct(false). //中间类型名称的样式为可导出的 ExampleColumns
		WithColumnsMethodRecvName("one").
		WithColumnsCheckFieldType(true) //这是新特性，非常建议启用

	cfg := gormcngen.NewConfigs([]interface{}{&Person{}, &Example{}}, options, absPath)
	cfg.Gen()
}
