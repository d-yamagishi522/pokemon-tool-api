package main

import (
	"sync"

	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemonweak"

	"github.com/jinzhu/gorm"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/db"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/db/seeds/sql"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
)

var odb *gorm.DB

func main() {
	odb := db.Init()
	defer odb.Close()

	odb.DropTableIfExists(
		&pokemon.Pokemon{},
		&attribute.Attribute{},
		&pokemonweak.PokemonWeak{},
	)

	odb.AutoMigrate(
		&pokemon.Pokemon{},
		&attribute.Attribute{},
		&pokemonweak.PokemonWeak{},
	)

	var wait sync.WaitGroup
	wait.Add(3)

	go func() {
		odb.DB().Exec(sql.InsertMockAttribute)
		wait.Done()
	}()

	go func() {
		odb.DB().Exec(sql.InsertMockPokemon)
		wait.Done()
	}()

	go func() {
		odb.DB().Exec(sql.InsertMockPokemonWeak)
		wait.Done()
	}()

	wait.Wait()
}
