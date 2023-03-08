package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Realizar conexion
var dsn = "root:admin@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"

var DataBase = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexion", err)
		panic(err)
	} else {
		return db
	}
}
