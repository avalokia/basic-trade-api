package controllers

import (
	"basic-trade-api/database"
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllVariant(c *gin.Context) {
	// Get search by name
	keyword := c.Query("search")

	// Get database
	db := database.GetDB()

	// Initialize product variable
	variants := []models.Variant{}

	// Get products and their Variants data
	result := db.Scopes(helpers.Paginate(c)).Where("variant_name ilike ?", "%"+keyword+"%").Find(&variants)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed getting data from DB",
			"message": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    variants,
	})
}

func CreateVariant(c *gin.Context) {
	// Bind request
	request := models.VariantRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed binding request",
			"message": err.Error()})
		return
	}

	// Parse UUID
	parsedUUID, err := uuid.Parse(request.ProductUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed parsing product UUID",
			"message": err.Error(),
		})
		return
	}

	// Get database
	db := database.GetDB()

	// Create new variant
	variant := models.Variant{
		VariantName: request.VariantName,
		ProductUUID: parsedUUID,
		Quantity:    request.Quantity,
	}

	err = db.Create(&variant).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed adding variant to DB",
			"message": err.Error(),
		})
		return
	}
	// Return created variant
	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Data":    variant,
	})
}

func UpdateVariant(c *gin.Context) {
	// Get UUID from param
	variantUUID, err := uuid.Parse(c.Param("variantUUID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed parsing Variant UUID",
			"message": err,
		})
	}
	// Bind request
	variant := models.Variant{}
	err = c.ShouldBind(&variant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed binding request",
			"message": err.Error(),
		})
		return
	}
	// Get database
	db := database.GetDB()
	// Update the variant
	err = db.Model(&variant).Where("uuid = ?", variantUUID).Updates(variant).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to update the variant to DB",
			"message": err.Error(),
		})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    variant,
	})
}

func DeleteVariant(c *gin.Context) {
	// Get UUID from param
	variantUUID, err := uuid.Parse(c.Param("variantUUID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed parsing Variant UUID",
			"message": err,
		})
	}
	// Get database
	db := database.GetDB()
	// Delete variant by ID
	variant := models.Variant{}
	err = db.Where("uuid = ?", variantUUID).Delete(&variant).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed deleting variant in DB",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    variant,
	})
}

func GetVariantByID(c *gin.Context) {
	// Get UUID from param
	variantUUID, err := uuid.Parse(c.Param("variantUUID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed parsing Variant UUID",
			"message": err,
		})
	}

	// Get database
	db := database.GetDB()

	// Retrieve variant
	variant := models.Variant{}
	result := db.First(&variant, "uuid = ?", variantUUID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve variant",
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    variant,
	})
}
