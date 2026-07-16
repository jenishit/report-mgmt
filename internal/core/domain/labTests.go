package domain

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID          uuid.UUID
	Name        string
	Description *string
	UpdatedAt   time.Time
}

type Panels struct {
	ID           uuid.UUID
	DepartmentID uuid.UUID
	Name         string
	Code         *string
	PanelPrice   *float64
	IsActive     bool
	UpdatedAt    time.Time
}

type TestCatalog struct {
	ID             uuid.UUID
	DepartmentID   uuid.UUID
	Name           string
	Code           *string
	SampleType     *string
	TestPrice      *float64
	TurnAroundTime *int
	IsActive       bool
	UpdatedAt      time.Time
}

type PanelComponent struct {
	PanelID       uuid.UUID
	TestCatalogID uuid.UUID
	SequenceNo    *int
	UpdatedAt     time.Time
}

type TestParameter struct {
	ID            uuid.UUID
	TestCatalogID uuid.UUID
	Name          string
	Unit          string
	ResultType    string
	SequenceNo    int
	IsActive      bool
	UpdatedAt     time.Time
}

type ReferenceRange struct {
	ID              uuid.UUID
	TestParameterID uuid.UUID
	Gender          string
	MinAge          *int
	MaxAge          *int
	MinValue        *float64
	MaxValue        *float64
	TextValue       *string
	UpdatedAt       time.Time
}
