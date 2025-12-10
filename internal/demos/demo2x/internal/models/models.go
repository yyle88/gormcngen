// Package models: Advanced GORM model examples for demo2 application
// Contains GORM models demonstrating custom table naming and column generation
// Shows integration with gormcnm for type-safe database operations in clean architecture
//
// models: demo2 应用程序的高级 GORM 模型示例
// 包含演示自定义表命名和列生成的 GORM 模型
// 在清晰架构中展示与 gormcnm 的集成以进行类型安全的数据库操作
package models

import "github.com/yyle88/gormcnm"

type Account struct {
	ID   uint
	Name string
}

func (*Account) TableName() string {
	return "accounts"
}

func (*Account) Columns() *AccountColumns {
	return &AccountColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:   "id",
		Name: "name",
	}
}

type AccountColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

type Purchase struct {
	ID          uint
	AccountID   uint
	ProductName string
	Amount      float64
}

func (*Purchase) TableName() string {
	return "purchases"
}

func (*Purchase) Columns() *PurchaseColumns {
	return &PurchaseColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:          "id",
		AccountID:   "account_id",
		ProductName: "product_name",
		Amount:      "amount",
	}
}

type PurchaseColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID          gormcnm.ColumnName[uint]
	AccountID   gormcnm.ColumnName[uint]
	ProductName gormcnm.ColumnName[string]
	Amount      gormcnm.ColumnName[float64]
}
