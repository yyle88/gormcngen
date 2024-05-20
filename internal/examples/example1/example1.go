package example1 // Package example1 我试试把注释写到这里行不行

import (
	"time"

	"github.com/yyle88/gormcnm"
)

type Person struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string `gorm:"not null,type:text"`
	BirthDate string
	Gender    bool
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (*Person) Columns() *personColumns {
	return &personColumns{
		ID:        "id",
		Name:      "name",
		BirthDate: "birth_date",
		Gender:    "gender",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type personColumns struct {
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	BirthDate gormcnm.ColumnName[string]
	Gender    gormcnm.ColumnName[bool]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"not null,type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (*Example) Columns() *exampleColumns {
	return &exampleColumns{
		ID:        "id",
		Name:      "name",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type exampleColumns struct {
	gormcnm.ColumnOperationClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
