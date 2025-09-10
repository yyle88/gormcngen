// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: ngen_test.go:21 -> models.TestGenerate
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yyle88/gormcnm"

func (c *User) Columns() *UserColumns {
	return &UserColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:   gormcnm.Cnm(c.ID, "id"),
		Name: gormcnm.Cnm(c.Name, "name"),
	}
}

type UserColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

func (c *Profile) Columns() *ProfileColumns {
	return &ProfileColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:     gormcnm.Cnm(c.ID, "id"),
		Bio:    gormcnm.Cnm(c.Bio, "bio"),
		UserID: gormcnm.Cnm(c.UserID, "user_id"),
	}
}

type ProfileColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID     gormcnm.ColumnName[uint]
	Bio    gormcnm.ColumnName[string]
	UserID gormcnm.ColumnName[uint]
}
