package repository

import (
	"github.com/jinzhu/gorm"
	"pokemon-tool-api/attribute"
)

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
	List() ([]*attribute.Attribute, error)
}

func (a *attributeRepository) List() ([]*attribute.Attribute, error) {
	attributes := []*attribute.Attribute{}
	err := a.Conn.Model(&attributes).Find(&attributes).Error
	if err != nil {
		return nil, err
	}
	return attributes, nil
}
