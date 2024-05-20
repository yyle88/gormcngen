package models

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t) //根据当前测试文件的路径找到对应的源文件路径
	t.Log(absPath)

	//出于安全起见，得限制这个文件就是想要的源码文件，当然实际上也不必这么小心的，或者刚开始小心点也行
	require.True(t, strings.HasSuffix(absPath, "models/gormcnm.gen.go"))

	//出于安全起见，需要需要判断目标文件是已经存在的，也就是你需要首先创建个空文件，让代码能找到此文件
	require.True(t, utils.IsFileExist(absPath))

	//在这里写下你要生成的 models 的对象列表，只有在列表里的才能被用于生成代码
	objects := []any{
		Example{},
	}

	cfg := gormcngen.NewConfigs(objects, true, absPath)
	cfg.Gen() //将会把生成后的代码写到目标位置
}
