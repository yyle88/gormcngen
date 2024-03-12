package example

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (*Person) Columns() *PersonColumns {
	return &PersonColumns{
		ID:          "id",
		Name:        "name",
		DateOfBirth: "date_of_birth",
		Gender:      "gender",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	}
}

type PersonColumns struct {
	gormcnm.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID          gormcnm.ColumnName[int32]
	Name        gormcnm.ColumnName[string]
	DateOfBirth gormcnm.ColumnName[string]
	Gender      gormcnm.ColumnName[bool]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
}
