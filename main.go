package main

import (
	// "gorm_api/db"
	"gorm_api/controller"
	"gorm_api/models"

	"fmt"

	"github.com/gin-gonic/gin"
)



func main(){
	fmt.Println("Starting server.......")

	r := gin.Default()
	models.CDB()
	r.GET("/books", controller.Index)
	r.GET("/books/:Id", controller.Show)
	r.POST("/books", controller.Create)
	r.PUT("/books/:Id", controller.Update)
	r.DELETE("/books", controller.Delete)



	r.Run()
}