package controllers

import (
	"manajemen-user/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func GetUsersByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.Where("ID = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": user,
		})
	}
}

func CreateUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var CreateUsersRequest struct {
			Name     string `json:"name" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
			Role     string `json:"role"`
		}

		err := c.ShouldBindBodyWithJSON(&CreateUsersRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(CreateUsersRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		user := models.User{
			Name:     CreateUsersRequest.Name,
			Email:    CreateUsersRequest.Email,
			Password: string(password),
			Role:     CreateUsersRequest.Role,
		}

		if err = db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully",
		})

	}
}

func UpdateUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var UpdateUsersRequest struct {
			Name     string `json:"name" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
			Role     string `json:"role"`
		}

		err := c.ShouldBindBodyWithJSON(&UpdateUsersRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id := c.Param("id")
		var user models.User

		if err := db.Where("ID = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(UpdateUsersRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		user.Name = UpdateUsersRequest.Name
		user.Email = UpdateUsersRequest.Email
		user.Password = string(password)
		user.Role = UpdateUsersRequest.Role

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User updated successfully",
			"user":    user,
		})
	}
}

func DeleteUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User

		if err := db.Where("ID = ?", id).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User delete successfully",
		})
	}
}
