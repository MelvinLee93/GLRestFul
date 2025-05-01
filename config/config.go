package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "localhost"
	user := "postgres"
	password := "S8rxc34Vi"
	dbName := "db_employee"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
}

func DB() *gorm.DB {
	return database
}