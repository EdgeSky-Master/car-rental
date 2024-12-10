package controllers

import (
	"car-rental/internal/models"
	"car-rental/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customers
	if err := database.DB.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customers
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&customer)
	c.JSON(http.StatusCreated, gin.H{"data": customer})
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customers
	if err := database.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found!"})
		return
	}

	var newCustomer models.Customers
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&customer).Updates(newCustomer)
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func DeleteCustomer(c *gin.Context) {
	var customer models.Customers
	if err := database.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found!"})
		return
	}

	database.DB.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"data": "Customer deleted"})
}