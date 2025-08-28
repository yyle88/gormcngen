package models // Package example8 我试试把注释写到这里行不行

import "github.com/yyle88/gormcnm"

type User struct {
	ID   uint   `cnm:"V主键"`
	Name string `cnm:"V名字"`
}

func (*User) TableName() string {
	return "users"
}

type Order struct {
	ID     uint    `cnm:"V订单主键"`
	UserID uint    `cnm:"V用户主键"`
	Amount float64 `cnm:"V订单金额"`
}

func (*Order) TableName() string {
	return "orders"
}

func (T *User) TableColumns(decoration gormcnm.ColumnNameDecoration) *UserColumns {
	return &UserColumns{
		V主键: gormcnm.Cmn(T.ID, "id", decoration),
		V名字: gormcnm.Cmn(T.Name, "name", decoration),
	}
}

type UserColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	V主键 gormcnm.ColumnName[uint]
	V名字 gormcnm.ColumnName[string]
}

func (T *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		V订单主键: gormcnm.Cmn(T.ID, "id", decoration),
		V用户主键: gormcnm.Cmn(T.UserID, "user_id", decoration),
		V订单金额: gormcnm.Cmn(T.Amount, "amount", decoration),
	}
}

type OrderColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	V订单主键 gormcnm.ColumnName[uint]
	V用户主键 gormcnm.ColumnName[uint]
	V订单金额 gormcnm.ColumnName[float64]
}
