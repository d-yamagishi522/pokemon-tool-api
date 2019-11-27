package usecase

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
	pokemonR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/repository"
)

type pokemonUsecase struct {
	pokemonRepo pokemonR.PokemonRepository
}

// PokemonUsecase interface
type PokemonUsecase interface {
	GetByID(id uuid.UUID) (*pokemon.Pokemon, error)
}

// NewPokemonUsecase mount pokemonUsecase
func NewPokemonUsecase(
	pokemon pokemonR.PokemonRepository,
) PokemonUsecase {
	return &pokemonUsecase{
		pokemonRepo: pokemon,
	}
}
