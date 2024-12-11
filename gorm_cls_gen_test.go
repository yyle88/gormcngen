package gormcngen_test

import (
	"testing"
	"time"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/runpath"
	"gorm.io/gorm"
)

type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}

type Demo struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);" cnm:"V名称"`
	Type string `gorm:"type:varchar(100);" cnm:"V类型"`
}

func TestConfigs_Generate(t *testing.T) {
	path := runpath.Path()
	t.Log(path)

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true).
		WithUseTagName(true)

	cfg := gormcngen.NewConfigs([]interface{}{
		&Example{},
		&Demo{},
	}, options, path)
	cfg.Generate()
}

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		Name: "name",
		Type: "type",
		Rank: "rank",
	}
}

type ExampleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	Name gormcnm.ColumnName[string]
	Type gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}

func (*Demo) Columns() *DemoColumns {
	return &DemoColumns{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		V名称:       "name",
		V类型:       "type",
	}
}

type DemoColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	V名称       gormcnm.ColumnName[string]
	V类型       gormcnm.ColumnName[string]
}
