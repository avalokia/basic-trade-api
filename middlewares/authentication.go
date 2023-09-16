package middlewares

import (
	"basic-trade-api/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Error":   "Authentication failed",
				"Message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}

}
