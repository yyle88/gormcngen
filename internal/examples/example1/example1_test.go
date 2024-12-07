package example1

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerate(t *testing.T) {
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	options := &gormcngen.Options{ExportGeneratedStruct: false}
	cfg := gormcngen.NewConfigs([]interface{}{&Person{}, &Example{}}, options, absPath)
	cfg.Gen()
}
