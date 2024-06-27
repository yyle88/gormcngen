package example3

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
