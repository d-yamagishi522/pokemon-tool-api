package usecase

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
)

func (p *pokemonUsecase) GetByID(id uuid.UUID) (*pokemon.Pokemon, error) {
	pokemon, err := p.pokemonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
