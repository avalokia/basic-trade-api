package helpers

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// Function to hash the password string. Returns the hashed password and error occured
func HashPass(p string) (string, error) {
	// Retrieve the salt from .env
	fmt.Println("Retrieving salt...")
	salt := os.Getenv("BCRYPT_SALT")
	// Convert to int
	fmt.Println("Converting salt...")
	saltInt, err := strconv.Atoi(salt)
	if err != nil {
		return "", err
	}
	fmt.Println("saltInt:", saltInt)
	// Convert password to []byte
	password := []byte(p)
	// Hash the password
	fmt.Println("Generating hash...")
	hash, err := bcrypt.GenerateFromPassword(password, saltInt)
	if err != nil {
		return "", err
	}

	fmt.Println("Hash generated:", string(hash))
	return string(hash), nil
}

// Function to compare the input and hashed password
func ComparePass(hashByte, passByte []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashByte, passByte)
	if err != nil {
		return false
	} else {
		return true
	}
}
