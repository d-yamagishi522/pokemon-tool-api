package main

import (
	"sync"

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
	)

	odb.AutoMigrate(
		&pokemon.Pokemon{},
		&attribute.Attribute{},
	)

	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		odb.DB().Exec(sql.InsertMockAttribute)
		wait.Done()
	}()
	wait.Wait()
}
