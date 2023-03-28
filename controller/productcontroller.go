package controller

import (
	// "gorm_api/models"
	"gorm_api/models"
	// "gorm_api/db"
	"encoding/json"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Melihat Data
func Index(c *gin.Context){
	var Product []models.Product
	models.DB.Find(&Product);
	c.JSON(http.StatusOK, gin.H{"Product": Product})

}
// Melihat Data Menggunakan ID
func Show(c *gin.Context){
	var Product models.Product
	Id := c.Param("Id")	

	if err := models.DB.First(&Product, Id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message":"Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		

		}
	}

	c.JSON(http.StatusOK, gin.H{"Product": Product})

}

// Menambahkan DATA
func Create(c *gin.Context){
	var Product models.Product
	if err := c.ShouldBindJSON(&Product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": err.Error()})
		return
	}
	models.DB.Create(&Product)
	c.JSON(http.StatusOK, gin.H{"Product": Product})


}

//Update Data
func Update(c *gin.Context){
	var Product models.Product
	Id := c.Param("Id")

	if err := c.ShouldBindJSON(&Product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": err.Error()})
		return
	}
	
	if models.DB.Model(&Product).Where("Id = ?", Id).Updates(&Product).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": "Update tidak dapat dilakukan"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Data telah di update"})

}

//Delete data
func Delete(c *gin.Context){

	var Product models.Product
	var input struct{
		Id json.Number

	}
	// input := map[string]string{"Id":"0"}


	//ERROR data
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": err.Error()})
		return
	}
	Id, _ := input.Id.Int64()
	//strconv.ParseInt(input["Id"], 10, 64)
	if models.DB.Delete(&Product, Id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"mesaage": "Tindak Dapat Menghapus"})
		return
	}
	//END ERROR DATA

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})//data Berhasil 

}