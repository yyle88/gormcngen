package gormcngen

import (
	"testing"
	"time"

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

	options := &Options{}
	cfg := NewConfig(&Example{}, options)
	res := cfg.Gen()
	t.Log(res.clsFuncCode)
	t.Log(res.nmClassCode)
	t.Log(neatjsons.S(res.moreImports))
}

func TestGen_UseTagName(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text" form:"name" json:"name" validate:"required" cnm:"V名称"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	options := &Options{
		UseTagName: true,
	}
	cfg := NewConfig(&Example{}, options)
	res := cfg.Gen()
	t.Log(res.clsFuncCode)
	t.Log(res.nmClassCode)
	t.Log(neatjsons.S(res.moreImports))
}

func TestGen_SkipNotTag(t *testing.T) {
	type Example struct {
		ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
		Name      string    `gorm:"not null,type:text" form:"name" json:"name" validate:"required" cnm:"V名称"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	options := &Options{
		UseTagName: true,
		SkipNotTag: true,
	}
	cfg := NewConfig(&Example{}, options)
	res := cfg.Gen()
	t.Log(res.clsFuncCode)
	t.Log(res.nmClassCode)
	t.Log(neatjsons.S(res.moreImports))
}
