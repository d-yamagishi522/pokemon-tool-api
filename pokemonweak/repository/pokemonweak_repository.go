package repository

import (
	"github.com/gofrs/uuid"
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
	ListByID(id uuid.UUID) ([]pokemonweak.PokemonWeak, error)
}

func (p *pokemonWeakRepository) Create(payload pokemonweak.PokemonWeak) error {
	err := p.Conn.Create(payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *pokemonWeakRepository) ListByID(id uuid.UUID) ([]pokemonweak.PokemonWeak, error) {
	list := []pokemonweak.PokemonWeak{}
	err := p.Conn.Model(&list).Where("pokemon_id = ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
