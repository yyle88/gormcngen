package example3

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/runpath/runtestpath"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file"
)

func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	utils_file.EXISTS.MustFile(absPath)
	t.Log(absPath)

	cfg := gormcngen.NewGenCfgsXPath([]interface{}{&Person{}, &Example{}}, absPath, true)
	cfg.GenWrite()
}
