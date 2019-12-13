package controller

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
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
	e.GET("/pokemon", handler.List)
	e.POST("/pokemon", handler.Create)
}

// GetByID return pokemon response
func (p *PokemonController) GetByID(ctx echo.Context) error {
	id := uuid.FromStringOrNil(ctx.Param("id"))
	// TODO: errの処理追加
	res, _ := p.pokemonUsecase.GetByID(id)
	return ctx.JSON(http.StatusOK, res)
}

// Create create pokemon
func (p *PokemonController) Create(ctx echo.Context) error {
	payload := []*pokemon.Request{}
	// TODO: errの処理追加
	_ = ctx.Bind(&payload)
	_ = p.pokemonUsecase.Create(payload)

	return ctx.NoContent(http.StatusNoContent)
}

// List return pokemon data list
func (p *PokemonController) List(ctx echo.Context) error {
	res, _ := p.pokemonUsecase.List()
	return ctx.JSON(http.StatusOK, res)
}
