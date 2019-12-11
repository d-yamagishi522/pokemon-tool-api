package pokemon

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute"
)

// Pokemon struct
type Pokemon struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Type1 string    `json:"type1"`
	Type2 *string   `json:"type2"`
}

// Data struct
type Data struct {
	ID         uuid.UUID             `json:"id"`
	Name       string                `json:"name"`
	Attributes []attribute.Attribute `json:"types"`
	Weaks      []attribute.Attribute `json:"weaks"`
}

// Request struct
type Request struct {
	Name  string                `json:"name"`
	Type1 string                `json:"type1"`
	Type2 *string               `json:"type2"`
	Weaks []attribute.Attribute `json:"weaks"`
}
