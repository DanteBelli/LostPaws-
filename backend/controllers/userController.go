package controllers

import (
	"LostPaws/backend/config"
	"LostPaws/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"mensaje": " Usuario Ingresada", "user": user})
}
func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se puede encontrar el usuario "})
	}
	c.JSON(http.StatusOK, user)
}
func UpdateUser(c *gin.Context) {

}
