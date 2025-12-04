// Package models demonstrates column generation with SQL reserved words as field names
// Auto contains Example struct with Create, Select, Update, Delete fields
// Used to showcase SafeCnm handling for SQL-conflicting column names
//
// models 包演示使用 SQL 保留字作为字段名的列生成
// 自动包含带有 Create、Select、Update、Delete 字段的 Example 结构体
// 用于展示 SafeCnm 对 SQL 冲突列名的处理
package models

import (
	"time"
)

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"not null,type:text"`
	Create    string    `gorm:"not null,type:text"` // 测测列名是关键字的情况
	Select    string    `gorm:"not null,type:text"` // 测测列名是关键字的情况
	Update    string    `gorm:"not null,type:text"` // 测测列名是关键字的情况
	Delete    string    `gorm:"not null,type:text"` // 测测列名是关键字的情况
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
