package models

import "github.com/yyle88/gormcnm"

func (*User) Columns() *UserColumns {
	return &UserColumns{
		ID:   "id",
		Name: "name",
	}
}

type UserColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

func (*Profile) Columns() *ProfileColumns {
	return &ProfileColumns{
		ID:     "id",
		Bio:    "bio",
		UserID: "user_id",
	}
}

type ProfileColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID     gormcnm.ColumnName[uint]
	Bio    gormcnm.ColumnName[string]
	UserID gormcnm.ColumnName[uint]
}
