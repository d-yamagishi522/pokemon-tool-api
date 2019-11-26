package main

import (
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/db"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/pokemon"
)

func main() {
	database := db.Init()
	database.AutoMigrate(
		&pokemon.Pokemon{},
	)
	p := []*pokemon.Pokemon{}
	database.Model(&p).Find(&p)
}
