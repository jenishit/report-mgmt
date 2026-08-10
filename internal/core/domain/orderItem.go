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

type ListOrders struct {
	ID                 uuid.UUID
	VisitNo            string
	PtFirstName        string
	PtLastName         string
	TestName           string
	TestCode           string
	TestPrice          float64
	PanelName          string
	PanelCode          string
	PanelPrice         float64
	Status             OrderStatus
	Price              float64
	CollectorFirstName string
	CollectorLastName  string
	CollectedAt        time.Time
}
