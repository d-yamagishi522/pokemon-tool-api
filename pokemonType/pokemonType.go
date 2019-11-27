package pokemontype

import uuid "github.com/satori/go.uuid"

// PokemonType struct
type PokemonType struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
