package usecase

import (
	"github.com/gofrs/uuid"
	"pokemon-tool-api/pokemon"
	"pokemon-tool-api/pokemonweak"
)

func (pu *pokemonUsecase) Create(payload []*pokemon.Request) error {
	for _, item := range payload {
		id, _ := uuid.NewV4()
		p := pokemon.Pokemon{
			ID:    id,
			Name:  item.Name,
			Type1: item.Type1,
			Type2: item.Type2,
		}
		_ = pu.pokemonRepo.Create(p)

		for _, w := range item.Weaks {
			pw := pokemonweak.PokemonWeak{
				PokemonID: id,
				Attribute: w.Name,
			}
			_ = pu.pokemonWeakRepo.Create(pw)
		}
	}
	return nil
}
