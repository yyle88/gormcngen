// Package example3 demonstrates SQL keyword column name handling
// Shows safe escaping of reserved words like CREATE, SELECT, UPDATE, DELETE
// Used to showcase SafeCnm method and SQL keyword protection
//
// example3 演示 SQL 关键字列名处理
// 展示 CREATE、SELECT、UPDATE、DELETE 等保留字的安全转义
// 用于展示 SafeCnm 方法和 SQL 关键字保护
package example3

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example3/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global database instance for SQL keyword handling tests
// 用于 SQL 关键字处理测试的全局数据库实例
var caseDB *gorm.DB

// TestMain sets up database for SQL keyword column testing
// Creates test data with SQL reserved word field names
//
// TestMain 为 SQL 关键字列测试设置数据库
// 创建带有 SQL 保留字字段名的测试数据
func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&models.Example{}))
	done.Done(db.Save(&models.Example{
		ID:        0,
		Name:      "abc",
		Create:    "e",
		Select:    "f",
		Update:    "g",
		Delete:    "h",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error)

	caseDB = db
	m.Run()
}

// TestSelect demonstrates handling SQL keyword column names with escaping
// Uses SafeCnm to escape reserved words like CREATE, SELECT, UPDATE, DELETE
//
// TestSelect 演示安全处理 SQL 关键字列名
// 使用 SafeCnm 正确转义 CREATE、SELECT、UPDATE、DELETE 等保留字
func TestSelect(t *testing.T) {
	var one models.Example
	c := one.Columns()
	require.NoError(t, caseDB.Where(c.Name.Eq("abc")).
		Where(c.Create.SafeCnm("``").Eq("e")).
		Where(c.Select.SafeCnm("``").Eq("f")).
		Where(c.Update.SafeCnm("``").Eq("g")).
		Where(c.Delete.SafeCnm("``").Eq("h")).
		First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(neatjsons.S(one))
}
