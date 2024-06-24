package example2usage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example2"
	"github.com/yyle88/gormcngen/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(&example2.Person{}))
	done.Done(db.Save(&example2.Person{
		ID:        0,
		Name:      "abc",
		BirthDate: "1970-01-01",
		Gender:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}).Error)
	done.Done(db.Save(&example2.Person{
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

func TestSelect(t *testing.T) {
	var one example2.Person
	c := one.Columns()
	require.NoError(t, caseDB.Where(c.Name.Eq("abc")).Where(c.Gender.IsFALSE()).First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(utils.SoftNeatString(one))
}

func TestSelect_2(t *testing.T) {
	var res []*example2.Person
	c := (&example2.Person{}).Columns()
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
	t.Log(utils.SoftNeatString(res))
}

func TestSelect_3(t *testing.T) {
	var one example2.Person
	c := one.Columns()

	qsx := c.Name.Qx("= ?", "abc").
		AND(
			c.Gender.Qc("IS FALSE").Qx(),
			c.BirthDate.Qx("=?", "1970-01-01"),
		)

	require.NoError(t, caseDB.Where(qsx.Qx2()).First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(utils.SoftNeatString(one))
}

func TestSelect_4(t *testing.T) {
	var one example2.Person
	c := one.Columns()

	qsx := c.Qx(
		c.Name.Eq("abc"),
	).AND(
		c.Qx(c.Gender.IsFALSE()),
		c.Qx(c.BirthDate.Eq("1970-01-01")),
	)

	require.NoError(t, caseDB.Where(qsx.Qx2()).First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(utils.SoftNeatString(one))
}
