// Package gormcngen_test validates CodeGenerationConfig with multi-model code generation
// Auto tests intelligent code injection and smart updates
// Demonstrates type-safe column struct and Columns() method generation
//
// gormcngen_test 验证 CodeGenerationConfig 的多模型代码生成
// 自动测试智能代码注入和增量更新
// 演示类型安全的列结构体和 Columns() 方法生成
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
	cfg.WithIsGenPreventEdit(false)
	cfg.Generate()
}

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		Name: "name",
		Type: "type",
		Rank: "rank",
	}
}

type ExampleColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	Name gormcnm.ColumnName[string]
	Type gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}

func (*Demo) Columns() *DemoColumns {
	return &DemoColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		V名称:       "name",
		V类型:       "type",
	}
}

type DemoColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	V名称       gormcnm.ColumnName[string]
	V类型       gormcnm.ColumnName[string]
}
