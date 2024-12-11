package example4

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
		WithColumnClassExportable(true). //中间类型名称的样式为可导出的 ExampleColumns
		WithUseTagName(true)

	cfg := gormcngen.NewConfigs([]interface{}{
		&Student{},
		&Class{},
	}, options, absPath)
	cfg.Gen()
}
