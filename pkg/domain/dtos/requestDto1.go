package dtos

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type RequestDto1 struct {
	Uuid          uuid.UUID
	Client        ClientDto
	TimeStamp     time.Time
	Status        string
	Total         decimal.Decimal
}
