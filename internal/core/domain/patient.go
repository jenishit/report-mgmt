package domain

import (
	"time"

	"github.com/google/uuid"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

type Patient struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     *string
	MRN       string
	DOB       time.Time
	Gender    Gender
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}
