package gormcngen

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/yyle88/gormcnm"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file"
	"gitlab.yyle.com/golang/uvyyle.git/utils_runtime/utils_runtestpath"
	"gorm.io/gorm"
)

func TestGenWrite(t *testing.T) {
	absPath := utils_runtestpath.TestPath(t)
	utils_file.EXISTS.MustFile(absPath)
	t.Log(absPath)

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
