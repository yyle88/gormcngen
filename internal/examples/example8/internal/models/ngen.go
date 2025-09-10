// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: ngen_test.go:25 -> models.TestGenerate
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yyle88/gormcnm"

func (T *User) TableColumns(decoration gormcnm.ColumnNameDecoration) *UserColumns {
	return &UserColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		V主键: gormcnm.Cmn(T.ID, "id", decoration),
		V名字: gormcnm.Cmn(T.Name, "name", decoration),
	}
}

type UserColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	V主键 gormcnm.ColumnName[uint]
	V名字 gormcnm.ColumnName[string]
}

func (T *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		V订单主键: gormcnm.Cmn(T.ID, "id", decoration),
		V用户主键: gormcnm.Cmn(T.UserID, "user_id", decoration),
		V订单金额: gormcnm.Cmn(T.Amount, "amount", decoration),
	}
}

type OrderColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	V订单主键 gormcnm.ColumnName[uint]
	V用户主键 gormcnm.ColumnName[uint]
	V订单金额 gormcnm.ColumnName[float64]
}
