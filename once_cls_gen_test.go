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

	config := gormcngen.NewSchemaConfig(&Example{}, &gormcngen.Options{})
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

	config := gormcngen.NewSchemaConfig(&Example{}, &gormcngen.Options{
		UseTagName: true,
	})
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

	config := gormcngen.NewSchemaConfig(&Example{}, &gormcngen.Options{
		UseTagName:            true,
		ExcludeUntaggedFields: true,
	})
	output := config.Gen()
	t.Log(output.GetMethodCode())
	t.Log(output.GetStructCode())
	t.Log(neatjsons.S(output.GetPkgImports()))
}
