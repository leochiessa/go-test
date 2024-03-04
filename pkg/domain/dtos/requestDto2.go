package dtos

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type RequestDto2 struct {
	Uuid          uuid.UUID
	Client        ClientDto
	TimeStamp     time.Time
	Status        string
	Total         decimal.Decimal
	RequestDetail []RequestDetailDto
}
