package usecase

import (
	"github.com/gofrs/uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
)

func (p *pokemonUsecase) GetByID(id uuid.UUID) (*pokemon.Pokemon, error) {
	pokemon, err := p.pokemonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
