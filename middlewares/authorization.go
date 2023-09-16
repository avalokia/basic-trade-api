package middlewares

import (
	"basic-trade-api/database"
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Check whether the user is authorized to update/delete the product
func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get database
		db := database.GetDB()
		// Get UUID from param
		productUUID, err := uuid.Parse(c.Param("productUUID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed parsing Product UUID",
				"message": err,
			})
		}
		// retrieve userData from jwt map claim
		userUUID, err := helpers.GetUserUUID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed parsing User UUID",
				"message": err,
			})
			return
		}

		// Retrieve the product data
		Product := models.Product{}
		err = db.First(&Product, "uuid = ?", productUUID).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Product not found",
			})
			c.Abort()
			return
		}

		// Check whether the logged in user is the product owner
		if Product.AdminUUID != userUUID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unathorized",
				"message": "You are not allowed to do this action",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Check whether the user is authorized to create the variant
func CreateVariantAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the product UUID from request body
		// Bind request
		variant := models.VariantRequest{}
		err := c.ShouldBind(&variant)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed binding request",
				"message": err.Error()})
			return
		}

		// Retrieve userData from jwt map claim
		userUUID, err := helpers.GetUserUUID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed parsing User UUID",
				"message": err,
			})
			return
		}

		// Find the owner of the product UUID
		// Get database
		db := database.GetDB()
		// Retrieve the product data
		Product := models.Product{}
		err = db.First(&Product, "uuid = ?", variant.ProductUUID).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Product not found",
			})
			c.Abort()
			return
		}

		// Check whether the logged in user is the product owner
		if Product.AdminUUID != userUUID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unathorized",
				"message": "You are not allowed to do this action",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Check whether the user is authorized to update/delete the variant
func VariantAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Variant UUID from param
		variantUUID, err := uuid.Parse(c.Param("variantUUID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed parsing Variant UUID",
				"message": err,
			})
		}
		// Retrieve userData from jwt map claim
		userUUID, err := helpers.GetUserUUID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed parsing User UUID",
				"message": err,
			})
			return
		}

		// Get database
		db := database.GetDB()
		// Find the product UUID of variant
		variant := models.Variant{}
		err = db.First(&variant, "uuid = ?", variantUUID).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Varant not found",
			})
			c.Abort()
			return
		}

		// Find the owner of the product UUID
		// Retrieve the product data
		Product := models.Product{}
		err = db.First(&Product, "uuid = ?", variant.ProductUUID).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Product not found",
			})
			c.Abort()
			return
		}

		// Check whether the logged in user is the product owner
		if Product.AdminUUID != userUUID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unathorized",
				"message": "You are not allowed to do this action",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
