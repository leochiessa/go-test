package entities

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type Request struct {
	Uuid          uuid.UUID
	Client        Client
	TimeStamp     time.Time
	Status        string
	Total         decimal.Decimal
	RequestDetail []RequestDetail
}
