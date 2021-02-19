package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string  `json:"id,omitempty" gorm:"primaryKey;size:50" msgpack:"id"`
	Name        string  `json:"name,omitempty" gorm:"size:256;notNull" msgpack:"name"`
	Price       float32 `json:"price,omitempty" gorm:"scale:18,precision:5"  msgpack:"price"`
	Description string  `json:"description,omitempty" gorm:"size:5000"  msgpack:"desscription"`
	CreatedAt   int64   `json:"created_at,omitempty" gorm:"autoCreateTime"  msgpack:"created_at"`
	UpdatedAt   int64   `json:"updated_at,omitempty" gorm:"autoUpdateTime"  msgpack:"updated_at"`
	ProductImage ProductImage `json:"product_image,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" msgpack:"product_image,as_array,omitempty"`
	ProductCategories []ProductCategory `json:"product_categories,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"  msgpack:"product_categories,as_array,omitempty"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewV4()

	p.ID = uuid.String()

	return
}
