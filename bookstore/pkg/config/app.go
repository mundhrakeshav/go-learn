package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
);

var (
	db * gorm.DB
)

func Connect()  {
 	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
  	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  	if err != nil{
		fmt.Println(err);
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB  {
	return db;
}