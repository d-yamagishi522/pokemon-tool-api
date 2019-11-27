package pokemon

import uuid "github.com/satori/go.uuid"

// Pokemon struct
type Pokemon struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
