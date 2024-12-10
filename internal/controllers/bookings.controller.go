package controllers

import (
	"car-rental/internal/models"
	"car-rental/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBookings(c *gin.Context) {
	var bookings []models.Bookings
	if err := database.DB.Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

func CreateBooking(c *gin.Context) {
	var booking models.Bookings
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&booking)
	c.JSON(http.StatusCreated, gin.H{"data": booking})
}

func UpdateBooking(c *gin.Context) {
	var booking models.Bookings
	if err := database.DB.Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking not found!"})
		return
	}

	var newBooking models.Bookings
	if err := c.ShouldBindJSON(&newBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&booking).Updates(newBooking)
	c.JSON(http.StatusOK, gin.H{"data": booking})
}
func DeleteBooking(c *gin.Context) {
	var booking models.Bookings
	if err := database.DB.Where("id = ?", c.Param("id")).First(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking not found!"})
		return
	}

	database.DB.Delete(&booking)
	c.JSON(http.StatusOK, gin.H{"data": "Booking deleted"})
}