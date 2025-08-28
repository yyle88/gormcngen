package example4

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example4/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global database instance for Chinese column name tests
// 用于中文列名测试的全局数据库实例
var caseDB *gorm.DB

// TestMain sets up database with Chinese column names and test data
// Creates students and classes with Chinese field names for testing
//
// TestMain 为中文列名和测试数据设置数据库
// 创建带有中文字段名的学生和班级数据进行测试
func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := done.VCE(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(
		&models.Student{},
		&models.Class{},
	))
	done.Done(db.Save(&models.Class{
		V班级编码: "LOVE",
		V班级名称: "情侣恋爱班级",
		V班主任名: "月老",
	}).Error)
	done.Done(db.Save(&models.Class{
		V班级编码: "HANDSOME",
		V班级名称: "型男明星班级",
		V班主任名: "玉帝",
	}).Error)

	done.Done(db.Save(&models.Student{
		ClassCode: "LOVE",
		Num:       0,
		Name:      "杨亦乐",
		Sex:       models.Male,
		BornDate:  "1990-08-08",
	}).Error)
	done.Done(db.Save(&models.Student{
		ClassCode: "LOVE",
		Num:       0,
		Name:      "刘亦菲",
		Sex:       models.Female,
		BornDate:  "1987-08-25",
	}).Error)
	done.Done(db.Save(&models.Student{
		ClassCode: "HANDSOME",
		Num:       0,
		Name:      "古天乐",
		Sex:       models.Male,
		BornDate:  "1970-10-21",
	}).Error)

	caseDB = db
	m.Run()
}

// TestSelect demonstrates basic query with Chinese column names
// Uses type-safe Chinese column references for database operations
//
// TestSelect 演示使用中文列名的基本查询
// 使用类型安全的中文列引用进行数据库操作
func TestSelect(t *testing.T) {
	var one models.Student
	c := one.Columns()
	require.NoError(t, caseDB.Where(c.V名字.Eq("杨亦乐")).First(&one).Error)
	require.Equal(t, "杨亦乐", one.Name)
	t.Log(neatjsons.S(one))
}

// TestSelect_x2x demonstrates multi-condition query with Chinese columns
// Uses multiple WHERE conditions with Chinese field names and enum values
//
// TestSelect_x2x 演示使用中文列的多条件查询
// 使用多个 WHERE 条件和中文字段名及枚举值
func TestSelect_x2x(t *testing.T) {
	var one models.Student
	c := one.Columns()
	require.NoError(t, caseDB.Where(c.V名字.Eq("刘亦菲")).Where(c.V性别.Eq(models.Female)).First(&one).Error)
	require.Equal(t, "刘亦菲", one.Name)
	t.Log(neatjsons.S(one))
}

// TestSelect_x3x demonstrates record not found scenario with Chinese columns
// Tests error handling when query conditions don't match any records
//
// TestSelect_x3x 演示使用中文列的记录不存在场景
// 测试当查询条件不匹配任何记录时的错误处理
func TestSelect_x3x(t *testing.T) {
	var one models.Student
	c := one.Columns()
	err := caseDB.Where(c.V名字.Eq("古天乐")).Where(c.V性别.Eq(models.Female)).First(&one).Error
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

// TestSelect_x4x demonstrates complex relational query with Chinese columns
// Shows querying classes and their associated students using Chinese field names
//
// TestSelect_x4x 演示使用中文列的复杂关联查询
// 展示使用中文字段名查询班级及其关联学生
func TestSelect_x4x(t *testing.T) {
	var classLove models.Class
	if c := classLove.Columns(); c.OK() {
		require.NoError(t, caseDB.Where(c.V班级编码.Eq("LOVE")).First(&classLove).Error)
	}
	t.Log(neatjsons.S(classLove))

	var students []*models.Student
	if c := new(models.Student).Columns(); c.OK() {
		require.NoError(t, caseDB.Where(c.V班级编码.Eq(classLove.V班级编码)).Find(&students).Error)
		require.Len(t, students, 2)
	}
	t.Log(neatjsons.S(students))
}
