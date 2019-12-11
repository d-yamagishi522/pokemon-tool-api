package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute"
	attributeC "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute/controller"
	attributeR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute/repository"
	attributeU "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute/usecase"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/db"
	mid "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/middleware"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
	pokemonC "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/controller"
	pokemonR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/repository"
	pokemonU "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon/usecase"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonattribute"
	pokemonAttributeR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonattribute/repository"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonweak"
	pokemonWeakR "gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonweak/repository"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))
	database := db.Init()
	database.AutoMigrate(
		&pokemon.Pokemon{},
		&attribute.Attribute{},
		&pokemonattribute.PokemonAttribute{},
		&pokemonweak.PokemonWeak{},
	)
	pokemonRepo := pokemonR.NewPokemonRepository(database)
	pokemonAttributeRepo := pokemonAttributeR.NewPokemonAttributeRepository(database)
	pokemonWeakRepo := pokemonWeakR.NewPokemonWeakRepository(database)
	pokemonUse := pokemonU.NewPokemonUsecase(
		pokemonAttributeRepo,
		pokemonWeakRepo,
		pokemonRepo,
	)
	pokemonC.NewPokemonController(e, pokemonUse)

	attributeRepo := attributeR.NewAttributeRepository(database)
	attributeUse := attributeU.NewAttributeUsecase(attributeRepo)
	attributeC.NewAttributeController(e, attributeUse)

	e.Logger.Fatal(e.Start(":8080"))
}
