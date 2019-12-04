package attribute

import uuid "github.com/satori/go.uuid"

// Attribute struct
type Attribute struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"type"`
}
