package controller

import (
	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/usecase"
)

// PokemonController struct
type PokemonController struct {
	pokemonUsecase usecase.PokemonUsecase
}

// NewPokemonController mount pokemonController
func NewPokemonController(
	e *echo.Echo,
	pokemon usecase.PokemonUsecase,
) {
	handler := &PokemonController{
		pokemonUsecase: pokemon,
	}
	e.GET("/pokemon/:id", handler.GetByID)
}

// GetByID return pokemon response
func (p *PokemonController) GetByID(ctx echo.Context) error {
	id := uuid.FromStringOrNil(ctx.Param("id"))
	// TODO: errの処理追加
	res, _ := p.pokemonUsecase.GetByID(id)
	return ctx.JSON(http.StatusOK, res)
}
