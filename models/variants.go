package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint      `gorm:"autoIncrement" json:"id"`
	UUID        uuid.UUID `gorm:"primaryKey"`
	VariantName string    `gorm:"not null;type:varchar" json:"variant_name" form:"variant_name"`
	Quantity    int       `gorm:"not null;" json:"quantity" form:"quantity"`
	ProductUUID uuid.UUID `json:"product_uuid" form:"product_uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	// Assign UUID
	v.UUID = uuid.New()

	// Validate struct
	_, err = govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}

	// Check the quantity
	if v.Quantity < 0 {
		err = errors.New("QUANTITY SHOULD NOT BE LESS THAN 0")
	}

	// Check whether the name is proper
	if len(v.VariantName) < 4 {
		err = errors.New("VARIANT NAME IS TOO SHORT. MINIMUM LENGTH IS 4 CHARACTERS")
	}

	return err
}
