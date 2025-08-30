package models

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (c *Person) Columns() *PersonColumns {
	return &PersonColumns{
		ID:        gormcnm.Cnm(c.ID, "id"),
		Name:      gormcnm.Cnm(c.Name, "name"),
		BirthDate: gormcnm.Cnm(c.BirthDate, "birth_date"),
		Gender:    gormcnm.Cnm(c.Gender, "gender"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
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

func (c *Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        gormcnm.Cnm(c.ID, "id"),
		Name:      gormcnm.Cnm(c.Name, "name"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
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
