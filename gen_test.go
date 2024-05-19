package gormcngen

import (
	"testing"
	"time"
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

	cfg := NewConfigXObject(&Example{}, false)
	res := cfg.Gen()
	t.Log(res.clsFuncCode)
	t.Log(res.nmClassCode)
	t.Log(res.moreImports)
}
