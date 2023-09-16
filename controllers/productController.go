package controllers

import (
	"basic-trade-api/database"
	"basic-trade-api/helpers"
	"basic-trade-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllProduct(c *gin.Context) {
	// Get search by name
	keyword := c.Query("search")

	// Get database
	db := database.GetDB()

	// Initialize product variable
	products := []models.Product{}

	// Get products and their Variants data with pagination and with search by name
	result := db.Scopes(helpers.Paginate(c)).Where("name ilike ?", "%"+keyword+"%").Preload("Variants").Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed getting data from DB",
			"message": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
	})
}

func CreateProduct(c *gin.Context) {
	// Get user data
	userUUID, err := helpers.GetUserUUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed getting user UUID",
			"message": err.Error(),
		})
	}
	// Bind URL
	var request models.ProductRequest
	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed binding request",
			"message": err.Error(),
		})
		return
	}

	// Upload to cloudinary
	fileName := helpers.RemoveExtension(request.File.Filename)

	uploadResult, err := helpers.UploadFile(request.File, fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed uploading file to Cloudinary",
			"message": err.Error(),
		})
		return
	}

	// New product
	var product models.Product = models.Product{
		Name:      request.Name,
		AdminUUID: userUUID,
		ImageURL:  uploadResult,
	}

	// Get database
	db := database.GetDB()
	// Create new product and check the error
	err = db.Create(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed adding product to DB",
			"message": err.Error(),
		})
		return
	}

	// Return created student
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func UpdateProduct(c *gin.Context) {
	// Retrieve userData
	userUUID, err := helpers.GetUserUUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed getting User UUID",
			"message": err.Error(),
		})
	}

	// Get UUID from param
	productUUID, err := uuid.Parse(c.Param("productUUID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed getting Product UUID",
			"message": err.Error(),
		})
	}

	// Bind URL
	var request models.ProductRequest
	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed binding request",
			"message": err.Error(),
		})
		return
	}

	var product models.Product = models.Product{}
	var updatedProduct models.Product = models.Product{}

	// If there's a new file, upload
	if request.File != nil {
		fileName := helpers.RemoveExtension(request.File.Filename)

		uploadResult, err := helpers.UploadFile(request.File, fileName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Failed uploading file to Cloudinary",
				"message": err.Error(),
			})
			return
		}
		// Updated product with new image URL
		updatedProduct = models.Product{
			Name:      request.Name,
			AdminUUID: userUUID,
			ImageURL:  uploadResult,
		}
	} else {
		updatedProduct = models.Product{
			Name:      request.Name,
			AdminUUID: userUUID,
		}
	}

	// Get database
	db := database.GetDB()

	// Update the product
	err = db.Model(&product).Where("uuid = ?", productUUID).Updates(updatedProduct).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to update the product to DB",
			"message": err.Error(),
		})
		return
	}

	// Return the updated product
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})

}

func DeleteProduct(c *gin.Context) {
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

	// Retrieve existing student & scores
	existingProduct := models.Product{}
	result := db.First(&existingProduct, "uuid = ?", productUUID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve product and variants",
			"message": result.Error,
		})
		return
	}
	// Delete asset in cloudinary
	err = helpers.DeleteFile(existingProduct.ImageURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed deleting file in Cloudinary",
			"message": err.Error(),
		})
		return
	}
	// Delete product by ID
	product := models.Product{}
	err = db.Where("uuid = ?", productUUID).Delete(&product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed deleting product from DB",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func GetProductByID(c *gin.Context) {
	// Get UUID from param
	productUUID, err := uuid.Parse(c.Param("productUUID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed parsing Product UUID",
			"message": err,
		})
	}

	// Get database
	db := database.GetDB()

	// Retrieve existing product and variants
	product := models.Product{}
	result := db.Preload("Variants").First(&product, "uuid = ?", productUUID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve product and variants",
			"message": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}
