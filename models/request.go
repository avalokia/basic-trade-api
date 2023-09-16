package models

import (
	"mime/multipart"
)

type ProductRequest struct {
	Name string                `form:"name"`
	File *multipart.FileHeader `form:"file"`
}

type VariantRequest struct {
	VariantName string `form:"variant_name"`
	Quantity    int    `form:"quantity"`
	ProductUUID string `form:"product_uuid"`
}
