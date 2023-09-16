package controllers

import (
	"basic-trade-api/database"
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	// Get DB instance
	db := database.GetDB()
	// Get content type from the header
	contentType := helpers.GetContentType(c)
	// Initiate admin variable
	Admin := models.Admin{}

	// Bind request
	if contentType == "application/json" {
		err := c.ShouldBindJSON(&Admin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed to bind JSON",
				"message": err.Error(),
			})
			return
		}
	} else {
		err := c.ShouldBind(&Admin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed to bind JSON",
				"message": err.Error(),
			})
			return
		}
	}

	// Input to DB
	err := db.Create(&Admin).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed creating data to DB",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    Admin,
	})
}

func UserLogin(c *gin.Context) {
	// Get DB instance
	db := database.GetDB()
	// Get content type from the header
	contentType := helpers.GetContentType(c)
	// Initiate admin variable
	Admin := models.Admin{}

	// Bind request
	if contentType == "application/json" {
		err := c.ShouldBindJSON(&Admin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed to bind JSON",
				"message": err.Error(),
			})
			return
		}
	} else {
		fmt.Println("Content is not application/json...")
		err := c.ShouldBind(&Admin)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed to bind request",
				"message": err.Error(),
			})
			return
		}
	}

	passwordInput := Admin.Password

	// Check whether email exists or not
	fmt.Println("Admin email:", Admin.Email)
	err := db.Where("email = ?", Admin.Email).First(&Admin).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Email not found",
		})
		return
	}

	// Check password
	valid := helpers.ComparePass([]byte(Admin.Password), []byte(passwordInput))
	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Password is incorrect",
		})
		return
	}

	// Generate token for login session
	token, err := helpers.GenerateToken(Admin.UUID, Admin.ID, Admin.Name, Admin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Generate token failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
