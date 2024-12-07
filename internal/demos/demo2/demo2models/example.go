package demo2models

import "github.com/yyle88/gormcnm"

type User struct {
	ID   uint
	Name string
}

func (*User) TableName() string {
	return "users"
}

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

type Order struct {
	ID     uint
	UserID uint
	Amount float64
}

func (*Order) TableName() string {
	return "orders"
}

func (*Order) Columns() *OrderColumns {
	return &OrderColumns{
		ID:     "id",
		UserID: "user_id",
		Amount: "amount",
	}
}

type OrderColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID     gormcnm.ColumnName[uint]
	UserID gormcnm.ColumnName[uint]
	Amount gormcnm.ColumnName[float64]
}
