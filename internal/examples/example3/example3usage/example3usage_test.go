package example3usage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example3"
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

	done.Done(db.AutoMigrate(&example3.Example{}))
	done.Done(db.Save(&example3.Example{
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

func TestSelect(t *testing.T) {
	var one example3.Example
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
