package gormcngen

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Exit(0)
}

func TestGen(t *testing.T) {
	t.Log(Gen(&Person{}, false))
}

type Person struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `gorm:"not null,type:text"`
	DateOfBirth string
	Gender      bool
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index,->"`
}
