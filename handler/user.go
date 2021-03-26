package handler

import (
	"net/http"

	"example.com/medium/db"
	"example.com/medium/models"
	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context) {
	var user models.User

	err := db.MySQL.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no user found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func FindUsers(c *gin.Context) {
	var users []models.User

	err := db.MySQL.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there's no users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = db.MySQL.FirstOrCreate(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func DeleteUser(c *gin.Context) {
	var user models.User

	err := db.MySQL.Where("id = ?", c.Param("id")).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no user found"})
		return
	}

	err = db.MySQL.Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
