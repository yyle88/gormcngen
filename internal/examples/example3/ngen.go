package example3

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        "id",
		Name:      "name",
		Order:     "order",
		Desc:      "desc",
		Asc:       "asc",
		Type:      "type",
		Create:    "create",
		Select:    "select",
		Update:    "update",
		Delete:    "delete",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type ExampleColumns struct {
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Order     gormcnm.ColumnName[string]
	Desc      gormcnm.ColumnName[string]
	Asc       gormcnm.ColumnName[string]
	Type      gormcnm.ColumnName[string]
	Create    gormcnm.ColumnName[string]
	Select    gormcnm.ColumnName[string]
	Update    gormcnm.ColumnName[string]
	Delete    gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
