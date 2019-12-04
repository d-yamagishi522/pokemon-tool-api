package pokemon

import (
	uuid "github.com/satori/go.uuid"
	"gitlab.com/pokemon-party-meta-chart/pokemon-tool-api/attribute"
)

// Pokemon struct
type Pokemon struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Response struct
type Response struct {
	ID         uuid.UUID             `json:"id"`
	Name       string                `json:"name"`
	Attributes []attribute.Attribute `json:"types"`
	Weak       []attribute.Attribute `json:"weaks"`
}
