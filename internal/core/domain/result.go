package domain

import (
	"time"

	"github.com/google/uuid"
)

type Flag string

const (
	Normal   Flag = "normal"
	Low      Flag = "low"
	High     Flag = "high"
	Critical Flag = "critical"
	NA       Flag = "na"
)

type Result struct {
	ID          uuid.UUID
	OrderID     uuid.UUID
	ParameterID uuid.UUID
	ResultValue string
	Flag        Flag
	PerformedBy uuid.UUID
	PerformedAt time.Time
	VerifiedBy  uuid.UUID
	Remarks     string
}
