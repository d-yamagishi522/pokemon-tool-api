package pokemonweak

import "github.com/gofrs/uuid"

// PokemonWeak struct
type PokemonWeak struct {
	PokemonID uuid.UUID
	Attribute string
}
