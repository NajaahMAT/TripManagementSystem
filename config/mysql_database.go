package config

import (
	"TripManagementSystem/helper"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "root"       // change as per your MySQL user
	password = "Najaah@123" // change as per your MySQL password
	dbName   = "tags"
	host     = "localhost"
	port     = 3306
)

// func MySQLDatabaseConnection() *sql.DB {
// 	db, err := sql.Open("mysql", "root:Najaah@123@tcp(localhost:3306)/tags")
// 	if err != nil {
// 		fmt.Println("error validating sql.Open arguments")
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	db.SetConnMaxLifetime(time.Minute * 3)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)

// 	return db
// }

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
