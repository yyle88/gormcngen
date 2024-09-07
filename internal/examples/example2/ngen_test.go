package example2

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	t.Log(absPath)
	utils.MustFile(absPath)

	options := &gormcngen.Options{IsSubClassExportable: true}
	cfg := gormcngen.NewConfigs([]interface{}{&Person{}, &Example{}}, options, absPath)
	cfg.Gen()
}
