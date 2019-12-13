package pokemon

import (
	"github.com/gofrs/uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute"
)

// Pokemon struct
type Pokemon struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Type1 string    `json:"type1"`
	Type2 *string   `json:"type2"`
	// Sort  int       `json:"sort"`
}

// Data struct
type Data struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Type1 string    `json:"type1"`
	Type2 *string   `json:"type2"`
	Weaks []string  `json:"weaks"`
}

// Request struct
type Request struct {
	Name  string                `json:"name"`
	Type1 string                `json:"type1"`
	Type2 *string               `json:"type2"`
	Weaks []attribute.Attribute `json:"weaks"`
}
