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
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	Name gormcnm.ColumnName[string]
	Type gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}
