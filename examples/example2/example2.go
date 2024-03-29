package example2 // Package example2 我试试把注释写到这里行不行

import (
	"time"

	"github.com/yyle88/gormcnm"
)

type Person struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string `gorm:"not null,type:text"`
	DateOfBirth string
	Gender      bool
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (*Person) Columns() *PersonColumns {
	return &PersonColumns{
		ID:          "id",
		Name:        "name",
		DateOfBirth: "date_of_birth",
		Gender:      "gender",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	}
}

type PersonColumns struct {
	gormcnm.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID          gormcnm.ColumnName[int32]
	Name        gormcnm.ColumnName[string]
	DateOfBirth gormcnm.ColumnName[string]
	Gender      gormcnm.ColumnName[bool]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
}

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"not null,type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        "id",
		Name:      "name",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type ExampleColumns struct {
	gormcnm.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
