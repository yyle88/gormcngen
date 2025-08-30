package models

import "github.com/yyle88/gormcnm"

func (c *User) Columns() *UserColumns {
	return &UserColumns{
		ID:   gormcnm.Cnm(c.ID, "id"),
		Name: gormcnm.Cnm(c.Name, "name"),
	}
}

type UserColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

func (c *Profile) Columns() *ProfileColumns {
	return &ProfileColumns{
		ID:     gormcnm.Cnm(c.ID, "id"),
		Bio:    gormcnm.Cnm(c.Bio, "bio"),
		UserID: gormcnm.Cnm(c.UserID, "user_id"),
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
