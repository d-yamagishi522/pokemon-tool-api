package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonweak"
)

type pokemonWeakRepository struct {
	Conn *gorm.DB
}

// NewPokemonWeakRepository mount PokemonWeakRepository
func NewPokemonWeakRepository(db *gorm.DB) PokemonWeakRepository {
	return &pokemonWeakRepository{
		Conn: db,
	}
}

// PokemonWeakRepository interface
type PokemonWeakRepository interface {
	Create(payload pokemonweak.PokemonWeak) error
}

func (p *pokemonWeakRepository) Create(payload pokemonweak.PokemonWeak) error {
	err := p.Conn.Create(payload).Error
	if err != nil {
		return err
	}
	return nil
}
