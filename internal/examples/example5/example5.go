package example5 // Package example5 我试试把注释写到这里行不行

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

func (one *Person) Columns() *personColumns {
	return &personColumns{
		ID:        gormcnm.Cnm(one.ID, "id"),
		Name:      gormcnm.Cnm(one.Name, "name"),
		BirthDate: gormcnm.Cnm(one.BirthDate, "birth_date"),
		Gender:    gormcnm.Cnm(one.Gender, "gender"),
		CreatedAt: gormcnm.Cnm(one.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(one.UpdatedAt, "updated_at"),
	}
}

type personColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
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

func (one *Example) Columns() *exampleColumns {
	return &exampleColumns{
		ID:        gormcnm.Cnm(one.ID, "id"),
		Name:      gormcnm.Cnm(one.Name, "name"),
		CreatedAt: gormcnm.Cnm(one.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(one.UpdatedAt, "updated_at"),
	}
}

type exampleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
