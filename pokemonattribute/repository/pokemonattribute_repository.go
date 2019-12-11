package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonattribute"
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
	Create(payload pokemonattribute.PokemonAttribute) error
}

func (p *pokemonAttributeRepository) Create(payload pokemonattribute.PokemonAttribute) error {
	err := p.Conn.Create(payload).Error
	if err != nil {
		return err
	}
	return nil
}
