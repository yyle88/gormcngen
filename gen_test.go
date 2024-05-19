package gormcngen

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/yyle88/runpath"
	"gitlab.yyle.com/golang/uvxlan.git/utils_gorm/utils_gorm/utils_gorm_cname"
	"gitlab.yyle.com/golang/uvyyle.git/utils_file"
	"gorm.io/gorm"
)

func TestGenWrite(t *testing.T) {
	absPath := runpath.Path()
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
	utils_gorm_cname.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID          utils_gorm_cname.ColumnName[uuid.UUID]
	Name        utils_gorm_cname.ColumnName[string]
	DateOfBirth utils_gorm_cname.ColumnName[string]
	Gender      utils_gorm_cname.ColumnName[bool]
	CreatedAt   utils_gorm_cname.ColumnName[time.Time]
	UpdatedAt   utils_gorm_cname.ColumnName[time.Time]
	DeletedAt   utils_gorm_cname.ColumnName[gorm.DeletedAt]
}
