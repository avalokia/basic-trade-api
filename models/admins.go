package models

import (
	"basic-trade-api/helpers"
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint      `gorm:"autoIncrement" json:"id"`
	UUID      uuid.UUID `gorm:"primaryKey" json:"uuid"`
	Name      string    `gorm:"type:varchar" json:"name" form:"name"`
	Email     string    `gorm:"not null;type:varchar;unique" json:"email" form:"email"`
	Password  string    `gorm:"not null;type:varchar" json:"password" form:"password"`
	Products  []Product `gorm:"constraint:OnDelete:CASCADE;foreignKey:AdminUUID;references:UUID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate the struct
	fmt.Println("Validating struct..")
	_, err = govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}

	// Check if email is valid or empty
	isEmail := govalidator.IsEmail(a.Email)
	isEmailEmpty := govalidator.IsNull(a.Email)
	if !isEmail || isEmailEmpty {
		return errors.New("EMAIL IS NOT VALID")
	}

	// Check password requirements
	if len(a.Password) < 6 {
		return errors.New("PASSWORD LENGTH IS TOO SHORT. MINIMUM 6 CHARACTERS")
	}

	// Check if name is empty
	isNameEmpty := govalidator.IsNull(a.Name)
	if isNameEmpty {
		return errors.New("NAME CANNOT BE EMPTY")
	}

	// Generate uuid
	a.UUID = uuid.New()

	// Replace the password data into the hashed password
	// fmt.Println("Hashing password..")
	a.Password, err = helpers.HashPass(a.Password)
	if err != nil {
		return err
	}

	fmt.Println("To be submitted:", a)

	return nil
}
