package entities

import "github.com/gofrs/uuid"

type Client struct {
	Uuid    uuid.UUID
	Name    string
	Address string
}
