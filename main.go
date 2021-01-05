package main

import (
	"pokemon-tool-api/attribute"
	attributeC "pokemon-tool-api/attribute/controller"
	attributeR "pokemon-tool-api/attribute/repository"
	attributeU "pokemon-tool-api/attribute/usecase"
	"pokemon-tool-api/db"
	mid "pokemon-tool-api/middleware"
	"pokemon-tool-api/pokemon"
	pokemonC "pokemon-tool-api/pokemon/controller"
	pokemonR "pokemon-tool-api/pokemon/repository"
	pokemonU "pokemon-tool-api/pokemon/usecase"
	"pokemon-tool-api/pokemonweak"
	pokemonWeakR "pokemon-tool-api/pokemonweak/repository"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))
	database := db.Init()
	database.AutoMigrate(
		&pokemon.Pokemon{},
		&attribute.Attribute{},
		&pokemonweak.PokemonWeak{},
	)
	pokemonRepo := pokemonR.NewPokemonRepository(database)
	pokemonWeakRepo := pokemonWeakR.NewPokemonWeakRepository(database)
	pokemonUse := pokemonU.NewPokemonUsecase(
		pokemonWeakRepo,
		pokemonRepo,
	)
	pokemonC.NewPokemonController(e, pokemonUse)

	attributeRepo := attributeR.NewAttributeRepository(database)
	attributeUse := attributeU.NewAttributeUsecase(attributeRepo)
	attributeC.NewAttributeController(e, attributeUse)

	e.Logger.Fatal(e.Start(":8080"))
}
