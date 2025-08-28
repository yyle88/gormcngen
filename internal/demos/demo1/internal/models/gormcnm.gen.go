package models

import "github.com/yyle88/gormcnm"

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		Name: "name",
		Type: "type",
		Rank: "rank",
	}
}

type ExampleColumns struct {
	// The column names and types of the model's columns // 模型各列的列名和类型
	Name gormcnm.ColumnName[string]
	Type gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}
