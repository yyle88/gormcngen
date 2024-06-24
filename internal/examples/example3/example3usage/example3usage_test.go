package example3usage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen/internal/examples/example3"
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

	done.Done(db.AutoMigrate(&example3.Example{}))
	done.Done(db.Save(&example3.Example{
		ID:        0,
		Name:      "abc",
		Order:     "a",
		Desc:      "b",
		Asc:       "c",
		Type:      "d",
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
		Where(c.Order.Safe().Eq("a")).
		Where(c.Desc.Safe().Eq("b")).
		Where(c.Asc.Safe().Eq("c")).
		Where(c.Type.Safe().Eq("d")).
		Where(c.Create.Safe().Eq("e")).
		Where(c.Select.Safe().Eq("f")).
		Where(c.Update.Safe().Eq("g")).
		Where(c.Delete.Safe().Eq("h")).
		First(&one).Error)
	require.Equal(t, "abc", one.Name)
	t.Log(utils.SoftNeatString(one))
}
