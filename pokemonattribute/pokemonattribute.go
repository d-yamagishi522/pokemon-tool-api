package pokemonattribute

import uuid "github.com/satori/go.uuid"

// PokemonAttribute struct
type PokemonAttribute struct {
	PokemonID   uuid.UUID
	AttributeID uuid.UUID
}
