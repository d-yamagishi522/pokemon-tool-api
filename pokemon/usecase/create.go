package usecase

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonattribute"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonweak"
)

func (pu *pokemonUsecase) Create(payload []*pokemon.Request) error {
	for _, item := range payload {
		id, _ := uuid.NewV4()
		p := pokemon.Pokemon{
			ID:   id,
			Name: item.Name,
		}
		_ = pu.pokemonRepo.Create(p)

		for _, a := range item.Attributes {
			pa := pokemonattribute.PokemonAttribute{
				PokemonID:   id,
				AttributeID: a.ID,
			}
			_ = pu.pokemonAttributeRepo.Create(pa)
		}

		for _, w := range item.Weaks {
			pw := pokemonweak.PokemonWeak{
				PokemonID:   id,
				AttributeID: w.ID,
			}
			_ = pu.pokemonWeakRepo.Create(pw)
		}
	}
	return nil
}
