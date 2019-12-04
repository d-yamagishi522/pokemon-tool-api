package usecase

import (
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute"
	attributeR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute/repository"
)

type attributeUsecase struct {
	attributeRepo attributeR.AttributeRepository
}

// NewAttributeUsecase mount attribute usecase
func NewAttributeUsecase(
	attribute attributeR.AttributeRepository,
) AttributeUsecase {
	return &attributeUsecase{
		attributeRepo: attribute,
	}
}

// AttributeUsecase interface
type AttributeUsecase interface {
	List() ([]*attribute.Attribute, error)
}

func (a *attributeUsecase) List() ([]*attribute.Attribute, error) {
	attribute, err := a.attributeRepo.List()
	if err != nil {
		return nil, err
	}
	return attribute, nil
}
