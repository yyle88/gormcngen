package models

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        "id",
		Name:      "name",
		Create:    "create",
		Select:    "select",
		Update:    "update",
		Delete:    "delete",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type ExampleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Create    gormcnm.ColumnName[string]
	Select    gormcnm.ColumnName[string]
	Update    gormcnm.ColumnName[string]
	Delete    gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
