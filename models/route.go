package models

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "admin"
	DB_PASSWORD = "lele123"
	DB_PORT     = 5432
	DB_NAME     = "local1"
	DEBUG_MODE  = true // true/false
)
var (
	DB  *gorm.DB
	err error
)


func CDB() *gorm.DB {
	
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	
	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
			panic(err)
	}
	fmt.Println("Go MySQL Tutorial")
	
	DB.AutoMigrate(&Product{})
	return DB
	 
	
}


