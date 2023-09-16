package models

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `gorm:"autoIncrement" json:"id"`
	UUID      uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"not null;type:varchar" json:"name"`
	ImageURL  string    `gorm:"not null;type:varchar" json:"image_url"`
	AdminUUID uuid.UUID `json:"admin_uuid"`
	Variants  []Variant `gorm:"constraint:OnDelete:CASCADE;foreignKey:ProductUUID;references:UUID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// Assign UUID
	p.UUID = uuid.New()

	// Validate struct
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return err
	}

	// Check whether image URL is generated
	if p.ImageURL == "" {
		err = errors.New("IMAGE URL IS NULL")
	}

	// Check whether the name is proper
	if len(p.Name) < 4 {
		err = errors.New("PRODUCT NAME IS TOO SHORT. MINIMUM LENGTH IS 4 CHARS")
	}

	return err
}
