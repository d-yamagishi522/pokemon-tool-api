package usecase

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
	pokemonR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/repository"
	pokemonWeakR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonweak/repository"
)

type pokemonUsecase struct {
	pokemonWeakRepo pokemonWeakR.PokemonWeakRepository
	pokemonRepo     pokemonR.PokemonRepository
}

// PokemonUsecase interface
type PokemonUsecase interface {
	GetByID(id uuid.UUID) (*pokemon.Pokemon, error)
	Create(payload []*pokemon.Request) error
}

// NewPokemonUsecase mount pokemonUsecase
func NewPokemonUsecase(
	pokemonWeak pokemonWeakR.PokemonWeakRepository,
	pokemon pokemonR.PokemonRepository,
) PokemonUsecase {
	return &pokemonUsecase{
		pokemonWeakRepo: pokemonWeak,
		pokemonRepo:     pokemon,
	}
}
