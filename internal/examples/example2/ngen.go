package example2

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (*Person) Columns() *PersonColumns {
	return &PersonColumns{
		ID:        "id",
		Name:      "name",
		BirthDate: "birth_date",
		Gender:    "gender",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type PersonColumns struct {
	gormcnm.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	//模型各个列名和类型:
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	BirthDate gormcnm.ColumnName[string]
	Gender    gormcnm.ColumnName[bool]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        "id",
		Name:      "name",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type ExampleColumns struct {
	gormcnm.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	//模型各个列名和类型:
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
