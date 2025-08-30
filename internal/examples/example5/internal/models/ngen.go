package models

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (one *Person) Columns() *PersonColumns {
	return &PersonColumns{
		ID:        gormcnm.Cnm(one.ID, "id"),
		Name:      gormcnm.Cnm(one.Name, "name"),
		BirthDate: gormcnm.Cnm(one.BirthDate, "birth_date"),
		Gender:    gormcnm.Cnm(one.Gender, "gender"),
		CreatedAt: gormcnm.Cnm(one.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(one.UpdatedAt, "updated_at"),
	}
}

type PersonColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	BirthDate gormcnm.ColumnName[string]
	Gender    gormcnm.ColumnName[bool]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}

func (one *Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        gormcnm.Cnm(one.ID, "id"),
		Name:      gormcnm.Cnm(one.Name, "name"),
		CreatedAt: gormcnm.Cnm(one.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(one.UpdatedAt, "updated_at"),
	}
}

type ExampleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
