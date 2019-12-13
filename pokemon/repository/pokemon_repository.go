package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
)

type pokemonRepository struct {
	Conn *gorm.DB
}

// NewPokemonRepository mount pokemonRepository
func NewPokemonRepository(db *gorm.DB) PokemonRepository {
	return &pokemonRepository{
		Conn: db,
	}
}

// PokemonRepository interface
type PokemonRepository interface {
	GetByID(id uuid.UUID) (*pokemon.Pokemon, error)
	Create(payload pokemon.Pokemon) error
	List() ([]*pokemon.Pokemon, error)
}

func (p *pokemonRepository) GetByID(id uuid.UUID) (*pokemon.Pokemon, error) {
	pokemon := &pokemon.Pokemon{}
	err := p.Conn.Model(&pokemon).Where("id = ?", id).Find(&pokemon).Error
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (p *pokemonRepository) Create(payload pokemon.Pokemon) error {
	err := p.Conn.Create(&payload).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *pokemonRepository) List() ([]*pokemon.Pokemon, error) {
	list := []*pokemon.Pokemon{}
	err := p.Conn.Model(&list).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
