package repository

import "github.com/jinzhu/gorm"

type attributeRepository struct {
	Conn *gorm.DB
}

// NewAttributeRepository mount AttributeRepository
func NewAttributeRepository(db *gorm.DB) AttributeRepository {
	return &attributeRepository{
		Conn: db,
	}
}

// AttributeRepository interface
type AttributeRepository interface {
}
