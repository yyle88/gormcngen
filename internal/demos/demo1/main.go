// Package main: Demo application showcasing gormcngen type-safe column generation
// Demonstrates automatic code generation and usage of generated column structs
// Shows real database operations with type-safe queries using generated Columns() methods
//
// main: 展示 gormcngen 类型安全列生成的演示应用程序
// 演示自动代码生成和生成列结构体的使用
// 展示使用生成的 Columns() 方法进行类型安全查询的真实数据库操作
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/demos/demo1/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//new db connection
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := done.VCE(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()

	//create example data
	_ = db.AutoMigrate(&models.Example{})
	_ = db.Save(&models.Example{Name: "abc", Type: "xyz", Rank: 123}).Error
	_ = db.Save(&models.Example{Name: "aaa", Type: "xxx", Rank: 456}).Error

	{
		var res models.Example
		err := db.Where("name=?", "abc").First(&res).Error
		done.Done(err)
		fmt.Println(res)
	}
	{ //select an example data
		var res models.Example
		var cls = res.Columns()
		if err := db.Where(cls.Name.Eq("abc")).
			Where(cls.Type.Eq("xyz")).
			Where(cls.Rank.Gt(100)).
			Where(cls.Rank.Lt(200)).
			First(&res).Error; err != nil {
			panic(errors.WithMessage(err, "wrong"))
		}
		fmt.Println(res)
	}
}
