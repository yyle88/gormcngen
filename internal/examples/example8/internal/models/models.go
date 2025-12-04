// Package models demonstrates Chinese cnm tag naming with TableName method
// Auto contains User and Order structs with Chinese field aliases
// Used to showcase namespace pattern with custom table names
//
// models 包演示带有 TableName 方法的中文 cnm 标签命名
// 自动包含带有中文字段别名的 User 和 Order 结构体
// 用于展示带有自定义表名的命名空间模式
package models

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
