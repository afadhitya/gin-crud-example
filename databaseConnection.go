package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	// open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:@/go_prj_a?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&todoModel{})

}
