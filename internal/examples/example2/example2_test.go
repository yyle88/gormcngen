package example2

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example2/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global database instance for tests
// 测试用的全局数据库实例
var caseDB *gorm.DB

// TestMain sets up database with test data before running tests
// 在运行测试前设置带有测试数据的数据库
func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := done.VCE(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(&models.Person{}))
	done.Done(db.Save(&models.Person{
		ID:        0,
		Name:      "abc",
		BirthDate: "1970-01-01",
		Gender:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error)
	done.Done(db.Save(&models.Person{
		ID:        0,
		Name:      "aaa",
		BirthDate: "2023-12-28",
		Gender:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error)

	caseDB = db
	m.Run()
}

// TestSelect demonstrates basic column-based query operations
// Uses type-safe column references for WHERE conditions
//
// TestSelect 演示基本的基于列的查询操作
// 使用类型安全的列引用进行 WHERE 条件查询
func TestSelect(t *testing.T) {
	var one models.Person
	c := one.Columns()
	require.NoError(t, caseDB.Where(c.Name.Eq("abc")).Where(c.Gender.IsFALSE()).First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(neatjsons.S(one))
}

// TestSelect_2 demonstrates complex query with OR/AND combinations
// Shows advanced column operation chaining and query building
//
// TestSelect_2 演示带有 OR/AND 组合的复杂查询
// 展示高级列操作链式调用和查询构建
func TestSelect_2(t *testing.T) {
	var res []*models.Person
	c := (&models.Person{}).Columns()
	require.NoError(t, caseDB.Where(c.Name.Qx("=?", "abc").
		OR(
			c.Name.Qx("=?", "aaa"),
		).
		AND(
			c.CreatedAt.Qx(">=?", time.Unix(0, 0).In(time.UTC)),
			c.UpdatedAt.Qx(">=?", time.Unix(0, 0).In(time.UTC)),
		).Qx4()).Where(c.Gender.IsFALSE()).Find(&res).Error)
	require.Contains(t, []string{"abc", "aaa"}, res[0].Name)
	require.Contains(t, []string{"abc", "aaa"}, res[1].Name)
	t.Log(neatjsons.S(res))
}

// TestSelect_3 demonstrates query expression building with custom conditions
// Uses column query expressions with AND logic combinations
//
// TestSelect_3 演示使用自定义条件构建查询表达式
// 使用列查询表达式与 AND 逻辑组合
func TestSelect_3(t *testing.T) {
	var one models.Person
	c := one.Columns()

	qsx := c.Name.Qx("= ?", "abc").
		AND(
			c.Gender.Qc("IS FALSE").Qx(),
			c.BirthDate.Qx("=?", "1970-01-01"),
		)

	require.NoError(t, caseDB.Where(qsx.Qx2()).First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(neatjsons.S(one))
}

// TestSelect_4 demonstrates advanced query building with nested expressions
// Shows column-level query expression composition and nesting
//
// TestSelect_4 演示使用嵌套表达式的高级查询构建
// 展示列级查询表达式组合和嵌套
func TestSelect_4(t *testing.T) {
	var one models.Person
	c := one.Columns()

	qsx := c.Qx(
		c.Name.Eq("abc"),
	).AND(
		c.Qx(c.Gender.IsFALSE()),
		c.Qx(c.BirthDate.Eq("1970-01-01")),
	)

	require.NoError(t, caseDB.Where(qsx.Qx2()).First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(neatjsons.S(one))
}
