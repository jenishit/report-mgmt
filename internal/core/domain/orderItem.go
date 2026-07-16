package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	Complete      OrderStatus = "completed"
	Collected     OrderStatus = "collected"
	ResultEntered OrderStatus = "result_entered"
)

type Order struct {
	ID          uuid.UUID
	VisitID     uuid.UUID
	TestID      uuid.UUID
	PanelID     uuid.UUID
	Status      OrderStatus
	Price       float64
	CollectedBy uuid.UUID
	CollectedAt time.Time
}
