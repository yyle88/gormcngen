package models

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
		WithColumnClassExportable(true). //中间类型名称的样式为可导出的 UserColumns
		WithEmbedColumnOperations(true). //其实没啥用
		WithUseTagName(true).
		WithTagKeyName("cnm").
		WithColumnsMethodRecvName("T").
		WithColumnsCheckFieldType(true). //这是新特性，非常建议启用
		WithIsGenFuncTableColumns(true).
		WithIsGenNewSimpleColumns(false)

	cfg := gormcngen.NewConfigs([]interface{}{&User{}, &Order{}}, options, absPath)
	cfg.Gen()
}
