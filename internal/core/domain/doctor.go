package domain

import (
	"time"

	"github.com/google/uuid"
)

type Doctor struct {
	ID             uuid.UUID
	FirstName      string
	LastName       string
	Email          string
	Phone          *string
	Qualification  string
	RegistrationNo string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	UpdatedBy      uuid.UUID
}
