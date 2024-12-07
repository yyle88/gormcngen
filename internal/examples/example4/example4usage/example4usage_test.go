package example4usage

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example4"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(
		&example4.Student{},
		&example4.Class{},
	))
	done.Done(db.Save(&example4.Class{
		V班级编码: "LOVE",
		V班级名称: "情侣恋爱班级",
		V班主任名: "月老",
	}).Error)
	done.Done(db.Save(&example4.Class{
		V班级编码: "HANDSOME",
		V班级名称: "型男明星班级",
		V班主任名: "玉帝",
	}).Error)

	done.Done(db.Save(&example4.Student{
		ClassCode: "LOVE",
		Num:       0,
		Name:      "杨亦乐",
		Sex:       example4.Male,
		BornDate:  "1990-08-08",
	}).Error)
	done.Done(db.Save(&example4.Student{
		ClassCode: "LOVE",
		Num:       0,
		Name:      "刘亦菲",
		Sex:       example4.Female,
		BornDate:  "1987-08-25",
	}).Error)
	done.Done(db.Save(&example4.Student{
		ClassCode: "HANDSOME",
		Num:       0,
		Name:      "古天乐",
		Sex:       example4.Male,
		BornDate:  "1970-10-21",
	}).Error)

	caseDB = db
	m.Run()
}

func TestSelect(t *testing.T) {
	var one example4.Student
	c := one.Columns()
	require.NoError(t, caseDB.Where(c.V名字.Eq("杨亦乐")).First(&one).Error)
	require.Equal(t, "杨亦乐", one.Name)
	t.Log(neatjsons.S(one))
}

func TestSelect_x2x(t *testing.T) {
	var one example4.Student
	c := one.Columns()
	require.NoError(t, caseDB.Where(c.V名字.Eq("刘亦菲")).Where(c.V性别.Eq(example4.Female)).First(&one).Error)
	require.Equal(t, "刘亦菲", one.Name)
	t.Log(neatjsons.S(one))
}

func TestSelect_x3x(t *testing.T) {
	var one example4.Student
	c := one.Columns()
	err := caseDB.Where(c.V名字.Eq("古天乐")).Where(c.V性别.Eq(example4.Female)).First(&one).Error
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func TestSelect_x4x(t *testing.T) {
	var classLove example4.Class
	if c := classLove.Columns(); c.OK() {
		require.NoError(t, caseDB.Where(c.V班级编码.Eq("LOVE")).First(&classLove).Error)
	}
	t.Log(neatjsons.S(classLove))

	var students []*example4.Student
	if c := new(example4.Student).Columns(); c.OK() {
		require.NoError(t, caseDB.Where(c.V班级编码.Eq(classLove.V班级编码)).Find(&students).Error)
		require.Len(t, students, 2)
	}
	t.Log(neatjsons.S(students))
}
