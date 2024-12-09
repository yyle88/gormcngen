package example3

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
		WithExportGeneratedStruct(true) //中间类型名称的样式为可导出的 ExampleColumns

	cfg := gormcngen.NewConfigs([]interface{}{&Example{}}, options, absPath)
	cfg.Gen()
}
