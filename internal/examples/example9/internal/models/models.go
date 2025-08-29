package models

type User struct {
	ID      uint
	Name    string
	Profile *Profile `gorm:"foreignKey:UserID"`
}

type Profile struct {
	ID     uint
	Bio    string
	UserID uint // 外键
}
