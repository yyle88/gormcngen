package example

import (
	"fmt"
	"testing"

	"github.com/yyle88/gormcngen"
	"gitlab.yyle.com/golang/uvcode.git/utils_gen"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file/utils_filepath"
	"gitlab.yyle.com/golang/uvyyle.git/utils_print"
	"gitlab.yyle.com/golang/uvyyle.git/utils_runtime/utils_runtestpath"
)

func TestGenerate(t *testing.T) {
	absPath := utils_runtestpath.SrcPath(t)
	utils_file.EXISTS.MustFile(absPath)
	t.Log(absPath)

	ptx := utils_print.NewPTX().Must()
	ptx.Println(fmt.Sprintf("package %s", utils_filepath.PATH.Name(utils_filepath.PARENT.Path(absPath))))
	ptx.Println(gormcngen.Gen(&Person{}, true))

	utils_gen.WriteSource(absPath, ptx.GetString())
}
