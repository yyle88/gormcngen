package example

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo"
)

func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	t.Log(absPath)
	require.True(t, utils.IsFileExist(absPath))

	ptx := utils.NewPTX()
	ptx.Println("package", syntaxgo.GetPkgName(absPath))
	ptx.Println(gormcngen.Gen(&Person{}, true))

	newSource := done.VAE(formatgo.FormatBytes(ptx.Bytes())).Nice()
	done.Done(utils.WriteFile(absPath, newSource))
}
