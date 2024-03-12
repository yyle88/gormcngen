package example3

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file"
	"gitlab.yyle.com/golang/uvyyle.git/utils_runtime/utils_runtestpath"
)

func TestGenerate(t *testing.T) {
	absPath := utils_runtestpath.SrcPath(t)
	utils_file.EXISTS.MustFile(absPath)
	t.Log(absPath)

	cfg := gormcngen.NewGenCfgsXPath([]interface{}{&Person{}, &Example{}}, absPath, true)
	cfg.GenWrite()
}
