package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Get limit from URL query
		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			limit = -1
		}

		// Get offset from URL query
		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			offset = -1
		}

		return db.Offset(offset).Limit(limit)
	}
}
