package example3

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	t.Log(absPath)
	require.True(t, utils.IsFileExist(absPath))

	options := &gormcngen.Options{IsSubClassExportable: true}
	cfg := gormcngen.NewConfigs([]interface{}{&Example{}}, options, absPath)
	cfg.Gen()
}
