package main

import (
	"StaticData/importer"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db* gorm.DB

func main(){
	db, _ := gorm.Open("mysql", "sfd:SFD_mysql2017@/sfd?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()

	importer.ImportXlsx()

}
