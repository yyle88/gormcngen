// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: ngen_test.go:20 -> models.TestGenerate
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import (
	"time"

	"github.com/yyle88/gormcnm"
)

func (c *Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		Name:      gormcnm.Cnm(c.Name, "name"),
		Create:    gormcnm.Cnm(c.Create, "create"),
		Select:    gormcnm.Cnm(c.Select, "select"),
		Update:    gormcnm.Cnm(c.Update, "update"),
		Delete:    gormcnm.Cnm(c.Delete, "delete"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
	}
}

type ExampleColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Create    gormcnm.ColumnName[string]
	Select    gormcnm.ColumnName[string]
	Update    gormcnm.ColumnName[string]
	Delete    gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
