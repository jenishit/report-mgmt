package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Registered Status = "registered"
	InProgress Status = "in_progress"
	Completed  Status = "completed"
	Cancelled  Status = "cancelled"
)

type Visit struct {
	ID           uuid.UUID
	VisitNo      string
	PatientID    uuid.UUID
	DoctorID     *uuid.UUID
	RegisteredBy uuid.UUID
	VisitDate    time.Time
	Status       Status
	CreatedAt    time.Time
	IsDeleted    bool
}

type ListVisit struct {
	ID                uuid.UUID
	VisitNo           string
	PatientFirstName  string
	PatientLastName   string
	DoctorFirstName   string
	DoctorLastName    string
	RegisterFirstName string
	RegisterLastName  string
	VisitDate         time.Time
	Status            Status
}
