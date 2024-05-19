package gormcngen

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen/internal/utils"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/runpath"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestGen(t *testing.T) {
	t.Log(Gen(&Person{}, false))
}

type Person struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `gorm:"not null,type:text"`
	DateOfBirth string
	Gender      bool
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index,->"`
}

func TestGenWrite(t *testing.T) {
	absPath := runpath.Path()
	t.Log(absPath)

	require.True(t, utils.IsFileExist(absPath))

	cfg := NewGenCfgsXPath([]interface{}{&Person{}}, absPath, true)
	cfg.GenWrite()
}

func (*Person) Columns() *PersonColumns {
	return &PersonColumns{
		ID:          "id",
		Name:        "name",
		DateOfBirth: "date_of_birth",
		Gender:      "gender",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
	}
}

type PersonColumns struct {
	gormcnm.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID          gormcnm.ColumnName[uuid.UUID]
	Name        gormcnm.ColumnName[string]
	DateOfBirth gormcnm.ColumnName[string]
	Gender      gormcnm.ColumnName[bool]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
	DeletedAt   gormcnm.ColumnName[gorm.DeletedAt]
}
