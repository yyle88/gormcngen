// Package main: Demo application showcasing gormcngen type-safe column generation
// Demonstrates automatic code generation and usage of generated column structs
// Shows database operations with type-safe queries using generated Columns() methods
//
// main: 展示 gormcngen 类型安全列生成的演示应用程序
// 演示自动代码生成和生成列结构体的使用
// 展示使用生成的 Columns() 方法进行类型安全查询的真实数据库操作
package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/demos/demo1/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Create new db connection // 创建数据库连接
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))

	// Create example data // 创建示例数据
	_ = db.AutoMigrate(&models.Example{})
	_ = db.Save(&models.Example{Name: "abc", Type: "xyz", Rank: 123}).Error
	_ = db.Save(&models.Example{Name: "aaa", Type: "xxx", Rank: 456}).Error

	{ // Native SQL query // 原生 SQL 查询
		var res models.Example
		done.Done(db.Where("name=?", "abc").First(&res).Error)
		/* Executed SQL // 执行的 SQL:
		SELECT * FROM `examples` WHERE name="abc"
		ORDER BY `examples`.`name` LIMIT 1
		*/
		zaplog.SUG.Debug(neatjsons.S(res))
	}
	{ // Type-safe column query // 类型安全的列查询
		var res models.Example
		var cls = res.Columns()
		done.Done(db.Where(cls.Name.Eq("abc")).First(&res).Error)
		/* Executed SQL // 执行的 SQL:
		SELECT * FROM `examples` WHERE name="abc"
		ORDER BY `examples`.`name` LIMIT 1
		*/
		zaplog.SUG.Debug(neatjsons.S(res))
	}
	{ // Native SQL query with multiple conditions // 原生 SQL 多条件查询
		var res models.Example
		done.Done(db.Where("name=?", "abc").
			Where("type=?", "xyz").
			Where("rank>?", 100).
			Where("rank<?", 200).
			First(&res).Error)
		/* Executed SQL // 执行的 SQL:
		SELECT * FROM `examples`
		WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200
		ORDER BY `examples`.`name` LIMIT 1
		*/
		zaplog.SUG.Debug(neatjsons.S(res))
	}
	{ // Type-safe column query with multiple conditions // 类型安全多条件查询
		var res models.Example
		var cls = res.Columns()
		done.Done(db.Where(cls.Name.Eq("abc")).
			Where(cls.Type.Eq("xyz")).
			Where(cls.Rank.Gt(100)).
			Where(cls.Rank.Lt(200)).
			First(&res).Error)
		/* Executed SQL // 执行的 SQL:
		SELECT * FROM `examples`
		WHERE name="abc" AND type="xyz" AND rank>100 AND rank<200
		ORDER BY `examples`.`name` LIMIT 1
		*/
		zaplog.SUG.Debug(neatjsons.S(res))
	}
}
