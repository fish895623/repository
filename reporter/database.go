package reporter

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID         int       `gorm:"column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	CreateTime time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime" json:"updateTime"`
}

func (User) TableName() string {
	return "User"
}

func ConnDataBase() (db *gorm.DB) {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	return
}
