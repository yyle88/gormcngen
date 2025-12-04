// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: ngen_test.go:26 -> models.TestGenerate
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

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
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:        gormcnm.Cmn(one.ID, "id", decoration),
		Name:      gormcnm.Cmn(one.Name, "name", decoration),
		BirthDate: gormcnm.Cmn(one.BirthDate, "birth_date", decoration),
		Gender:    gormcnm.Cmn(one.Gender, "gender", decoration),
		CreatedAt: gormcnm.Cmn(one.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(one.UpdatedAt, "updated_at", decoration),
	}
}

type PersonColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
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
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:        gormcnm.Cmn(one.ID, "id", decoration),
		Name:      gormcnm.Cmn(one.Name, "name", decoration),
		CreatedAt: gormcnm.Cmn(one.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(one.UpdatedAt, "updated_at", decoration),
	}
}

type ExampleColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
