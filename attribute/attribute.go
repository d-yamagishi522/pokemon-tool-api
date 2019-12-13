package attribute

import "github.com/gofrs/uuid"

// Attribute struct
type Attribute struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"type"`
}
