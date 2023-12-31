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

type Revision struct {
	ID         int       `gorm:"column:id" json:"id"`
	UserId     int       `gorm:"column:userId" json:"userId"`
	CreateTime time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime" json:"updateTime"`
	DeleteTime time.Time `gorm:"column:deleteTime" json:"deleteTime"`
}

func (Revision) TableName() string {
	return "Revision"
}

type JobIndexer struct {
	ID         int       `gorm:"column:id" json:"id"`
	Model      string    `gorm:"column:model" json:"model"`
	Version    string    `gorm:"column:version" json:"version"`
	UserId     int       `gorm:"column:userId" json:"userId"`
	CreateTime time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime" json:"updateTime"`
	DeleteTime time.Time `gorm:"column:deleteTime" json:"deleteTime"`
}

func (JobIndexer) TableName() string {
	return "JobIndexer"
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
