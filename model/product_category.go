package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	Id string `json:"id,omitempty" gorm:"primaryKey;size:50" msgpack:"id"`
	ProductId string `json:"product_id,omitempty" gorm:"size:50" msgpack:"product_id"`
	CategoryName string `json:"category_name,omitempty" gorm:"size:256" msgpack:"category_name"`
}

func (p *ProductCategory) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewV4()

	p.Id = uuid.String()

	return
}
