package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetUserUUID(c *gin.Context) (uuid.UUID, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userUUID, err := uuid.Parse(userData["uuid"].(string))
	if err != nil {
		return uuid.Nil, err
	}
	return userUUID, nil
}
