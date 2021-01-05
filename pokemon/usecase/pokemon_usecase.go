package usecase

import (
	"github.com/gofrs/uuid"
	"pokemon-tool-api/pokemon"
	pokemonR "pokemon-tool-api/pokemon/repository"
	pokemonWeakR "pokemon-tool-api/pokemonweak/repository"
)

type pokemonUsecase struct {
	pokemonWeakRepo pokemonWeakR.PokemonWeakRepository
	pokemonRepo     pokemonR.PokemonRepository
}

// PokemonUsecase interface
type PokemonUsecase interface {
	GetByID(id uuid.UUID) (*pokemon.Pokemon, error)
	Create(payload []*pokemon.Request) error
	List() ([]*pokemon.Data, error)
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
