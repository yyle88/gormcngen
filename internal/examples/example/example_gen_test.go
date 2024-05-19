package example

import (
	"fmt"
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/runpath/runtestpath"
	"gitlab.yyle.com/golang/uvcode.git/utils_gen"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file/utils_filepath"
)

func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	utils_file.EXISTS.MustFile(absPath)
	t.Log(absPath)

	ptx := utils.NewPTX()
	ptx.Println(fmt.Sprintf("package %s", utils_filepath.PATH.Name(utils_filepath.PARENT.Path(absPath))))
	ptx.Println(gormcngen.Gen(&Person{}, true))

	utils_gen.WriteSource(absPath, ptx.GetString())
}
