package models

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (one *Person) Columns() *PersonColumns {
	return one.TableColumns(gormcnm.NewPlainDecoration())
}

func (one *Person) TableColumns(decoration gormcnm.ColumnNameDecoration) *PersonColumns {
	return &PersonColumns{
		ID:        gormcnm.Cmn(one.ID, "id", decoration),
		Name:      gormcnm.Cmn(one.Name, "name", decoration),
		BirthDate: gormcnm.Cmn(one.BirthDate, "birth_date", decoration),
		Gender:    gormcnm.Cmn(one.Gender, "gender", decoration),
		CreatedAt: gormcnm.Cmn(one.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(one.UpdatedAt, "updated_at", decoration),
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
	return one.TableColumns(gormcnm.NewPlainDecoration())
}

func (one *Example) TableColumns(decoration gormcnm.ColumnNameDecoration) *ExampleColumns {
	return &ExampleColumns{
		ID:        gormcnm.Cmn(one.ID, "id", decoration),
		Name:      gormcnm.Cmn(one.Name, "name", decoration),
		CreatedAt: gormcnm.Cmn(one.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(one.UpdatedAt, "updated_at", decoration),
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
