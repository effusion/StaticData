package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "sfd:SFD_mysql2017@tcp(localhost:3307)/sfd?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db err: ", err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.Set("gorm:auto_preload", true)
	db.LogMode(false)
	//db.LogMode(true)
	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		fmt.Printf("Error closing DB %v", err)
	}
}
