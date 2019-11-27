package main

import (
	"github.com/labstack/echo"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/db"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
	pokemonC "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/controller"
	pokemonR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/repository"
	pokemonU "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/usecase"
)

func main() {
	e := echo.New()

	database := db.Init()
	database.AutoMigrate(
		&pokemon.Pokemon{},
	)
	pokemonRepo := pokemonR.NewPokemonRepository(database)
	pokemonUse := pokemonU.NewPokemonUsecase(
		pokemonRepo,
	)
	pokemonC.NewPokemonController(e, pokemonUse)

	e.Logger.Fatal(e.Start(":8080"))
}
