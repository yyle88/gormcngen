// Package models provides example GORM model definitions for gormcngen demonstration
// Auto contains Person and Example structs with standard GORM tags
// Used to showcase basic column generation patterns
//
// models 包为 gormcngen 演示提供示例 GORM 模型定义
// 自动包含带有标准 GORM 标签的 Person 和 Example 结构体
// 用于展示基本的列生成模式
package models

import (
	"time"
)

type Person struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string `gorm:"not null,type:text"`
	BirthDate string
	Gender    bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"not null,type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
