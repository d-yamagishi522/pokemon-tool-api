package repository

import (
	"github.com/jinzhu/gorm"
)

type pokemonAttributeRepository struct {
	Conn *gorm.DB
}

// NewPokemonAttributeRepository mount PokemonAttributeRepository
func NewPokemonAttributeRepository(db *gorm.DB) PokemonAttributeRepository {
	return &pokemonAttributeRepository{
		Conn: db,
	}
}

// PokemonAttributeRepository interface
type PokemonAttributeRepository interface {
}
