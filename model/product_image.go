package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	ID string `json:"id,omitempty" gorm:"primaryKey;size:50" msgpack:"id"`
	ProductId string `json:"product_id,omitempty" gorm:"size:50" msgpack:"product_id"`
	Filename string `json:"filename,omitempty" gorm:"size:256;notNull" msgpack:"filename"`
}

func (p *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, _ := uuid.NewV4()

	p.ID = uuid.String()

	return
}
