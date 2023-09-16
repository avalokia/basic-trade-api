package helpers

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateToken(uuid uuid.UUID, id uint, name string, email string) (string, error) {
	// Store data in jwt
	claims := jwt.MapClaims{
		"uuid":  uuid,
		"id":    id,
		"name":  name,
		"email": email,
	}

	// Parse token
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign token with secret key
	secretKey := os.Getenv("JWT_SECRET_KEY")
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("SIGN IN TO PROCEED")

	headerToken := GetAuthorization(c)
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	// headerToken[0] is "Bearer", so we take the token after
	stringToken := strings.Split(headerToken, " ")[1]

	// Parse token
	token, err := jwt.Parse(stringToken, convertToken)
	if err != nil {
		return nil, err
	}

	// Retrieve token claims (user data)
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}

func convertToken(t *jwt.Token) (interface{}, error) {
	// Check whether token is valid with the signing method
	_, ok := t.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("SIGN IN TO PROCEED")
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return []byte(secretKey), nil
}
