package main

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/demos/demo1/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//new db connection
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
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
