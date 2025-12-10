// Package models: GORM model definitions for demo1 application
// Contains example GORM models used to showcase automatic column generation
// Demonstrates various GORM tag configurations and field types in a clean structure
//
// models: demo1 应用程序的 GORM 模型定义
// 包含用于展示自动列生成的 GORM 模型示例
// 在清晰结构中演示各种 GORM 标签配置和字段类型
package models

// Example is a GORM model defining 3 fields (name, type, rank)
// Demonstrates basic GORM field configurations and column mappings
// Used as input for gormcngen to generate type-safe column structures
//
// Example 是一个定义 3 个字段（name、type、rank）的 GORM 模型
// 演示基本的 GORM 字段配置和列映射
// 作为 gormcngen 的输入生成类型安全的列结构体
type Example struct {
	Name string `gorm:"primary_key;type:varchar(100);"`
	Type string `gorm:"column:type;"`
	Rank int    `gorm:"column:rank;"`
}
