package usecase

import (
	"github.com/gofrs/uuid"
	"pokemon-tool-api/pokemon"
)

func (p *pokemonUsecase) GetByID(id uuid.UUID) (*pokemon.Pokemon, error) {
	pokemon, err := p.pokemonRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

func (p *pokemonUsecase) List() ([]*pokemon.Data, error) {
	list, err := p.pokemonRepo.List()
	if err != nil {
		return nil, err
	}
	newList := []*pokemon.Data{}
	for _, item := range list {
		weaks := []string{}
		pws, err := p.pokemonWeakRepo.ListByID(item.ID)
		if err != nil {
			return nil, err
		}
		for _, pw := range pws {
			weaks = append(weaks, pw.Attribute)
		}
		data := &pokemon.Data{
			ID:    item.ID,
			Name:  item.Name,
			Type1: item.Type1,
			Type2: item.Type2,
			Weaks: weaks,
		}
		newList = append(newList, data)
	}
	return newList, nil
}
