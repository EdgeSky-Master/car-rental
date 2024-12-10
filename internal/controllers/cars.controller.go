package controllers

import (
	"car-rental/internal/models"
	"car-rental/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	var cars []models.Cars
    if err := database.DB.Find(&cars).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": cars})
}

func CreateCar(c *gin.Context) {
	var car models.Cars
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&car)
	c.JSON(http.StatusCreated, gin.H{"data": car})
}

func UpdateCar(c *gin.Context) {
	var car models.Cars
	if err := database.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Car not found!"})
		return
	}

	var newCar models.Cars
	if err := c.ShouldBindJSON(&newCar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&car).Updates(newCar)
	c.JSON(http.StatusOK, gin.H{"data": car})
}

func DeleteCar(c *gin.Context) {
	var car models.Cars
	if err := database.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Car not found!"})
		return
	}

	database.DB.Delete(&car)
	c.JSON(http.StatusOK, gin.H{"data": "Car deleted"})
}