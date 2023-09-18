package configs

import (
	"errors"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

// Initialize cloudinary instance
func InitializeCloudinary() (*cloudinary.Cloudinary, error) {
	// Retrieve from .env
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }
	cloud := os.Getenv("CLOUDINARY_CLOUD_NAME")
	key := os.Getenv("CLOUDINARY_KEY")
	secret := os.Getenv("CLOUDINARY_SECRET")
	if (cloud == "") || (key == "") || (secret == "") {
		return nil, errors.New("CLOUDINARY PARAMETER(S) ARE EMPTY")
	}
	// Add cloudinary instance from provided parameters
	cld, err := cloudinary.NewFromParams(cloud, key, secret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
