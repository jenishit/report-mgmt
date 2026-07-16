package domain

import (
	"time"

	"github.com/google/uuid"
)

type LabSettings struct {
	ID             uuid.UUID
	LabName        string
	Tagline        *string
	Address        *string
	Phone          *string
	Email          *string
	RegistrationNo *string
	ReportFooter   *string
	UpdatedAt      time.Time
	UpdatedBy      uuid.UUID
}
