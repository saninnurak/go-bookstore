package config

import {
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
}

var (
	db * gorm.db
)

func Connect(){
	d, err := gorm.Open("mysql", "sensei:upadaj/simplerest?charser=utf8&parseTime=True&loc=Local")
	if err := nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
