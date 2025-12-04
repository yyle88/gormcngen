// Package gormcngen_test validates SchemaConfig with single-model code generation
// Auto tests schema parsing, options configuration, and output generation
// Demonstrates various generation modes including tag handling and field filtering
//
// gormcngen_test 验证 SchemaConfig 的单模型代码生成
// 自动测试 schema 解析、选项配置和输出生成
// 演示包括标签处理和字段过滤在内的各种生成模式
package gormcngen_test

import (
	"testing"
	"time"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGen(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	config := gormcngen.NewSchemaConfig(&Example{}, gormcngen.NewOptions())
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}

func TestGen_UseTagName(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text" form:"name" json:"name" validate:"required" cnm:"V名称"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	options := gormcngen.NewOptions().
		WithUseTagName(true)

	config := gormcngen.NewSchemaConfig(&Example{}, options)
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}

func TestGen_ExcludeUntaggedFields(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text" form:"name" json:"name" validate:"required" cnm:"V名称"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	options := gormcngen.NewOptions().
		WithUseTagName(true).
		WithExcludeUntaggedFields(true)

	config := gormcngen.NewSchemaConfig(&Example{}, options)
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}

func TestGen_ForeignKeyAssociation(t *testing.T) {
	type Profile struct {
		ID     uint
		Bio    string
		UserID uint // 外键
	}

	type Example struct {
		ID      uint
		Name    string
		Profile Profile `gorm:"foreignKey:UserID"`
	}

	options := gormcngen.NewOptions().
		WithUseTagName(true)

	config := gormcngen.NewSchemaConfig(&Example{}, options)
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}

func TestGen_ColumnsMethodRecvName(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text" form:"name" json:"name" validate:"required" cnm:"V名称"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	options := gormcngen.NewOptions().
		WithUseTagName(true).
		WithColumnsMethodRecvName("example")

	config := gormcngen.NewSchemaConfig(&Example{}, options)
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}

func TestGen_ColumnsCheckFieldType(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text" form:"name" json:"name" validate:"required" cnm:"V名称"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	options := gormcngen.NewOptions().
		WithUseTagName(true).
		WithColumnsMethodRecvName("example").
		WithColumnsCheckFieldType(true)

	config := gormcngen.NewSchemaConfig(&Example{}, options)
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}

func TestGen_IsGenFuncTableColumns(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text" form:"name" json:"name" validate:"required" cnm:"V名称"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	options := gormcngen.NewOptions().
		WithUseTagName(true).
		WithColumnsMethodRecvName("example").
		WithColumnsCheckFieldType(true).
		WithIsGenFuncTableColumns(true)

	config := gormcngen.NewSchemaConfig(&Example{}, options)
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetMethodTableColumnsCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}
