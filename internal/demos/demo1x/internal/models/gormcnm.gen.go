// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: gormcnm.gen_test.go:34 -> models.TestGenerate
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yyle88/gormcnm"

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		Name: "name",
		Type: "type",
		Rank: "rank",
	}
}

type ExampleColumns struct {
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	Name gormcnm.ColumnName[string]
	Type gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}
