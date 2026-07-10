package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type CreatePatient struct {
	FirstName string        `json:"first_name" binding:"required"`
	LastName  string        `json:"last_name" binding:"required"`
	Email     string        `json:"email" binding:"required,email"` //the acceptable string type is of email type
	Phone     string        `json:"phone"`
	MRN       string        `json:"mrn"`
	DOB       string        `json:"dob" binding:"required"`
	Gender    domain.Gender `json:"gender"`
	Address   string        `json:"address"`
}

type PatientResponse struct {
	ID        uuid.UUID     `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"` //the acceptable string type is of email type
	Phone     string        `json:"phone"`
	MRN       string        `json:"mrn"`
	DOB       time.Time     `json:"dob"`
	Gender    domain.Gender `json:"gender"`
	Address   string        `json:"address"`
}

func PtResponse(p *domain.Patient) *PatientResponse {
	return &PatientResponse{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
		Phone:     *p.Phone,
		MRN:       p.MRN,
		DOB:       p.DOB,
		Gender:    p.Gender,
		Address:   p.Address,
	}
}

func PtResponses(p []*domain.Patient) []*PatientResponse {
	pts := make([]*PatientResponse, 0, len(p))

	for _, pt := range p {
		pts = append(pts, &PatientResponse{
			ID:        pt.ID,
			FirstName: pt.FirstName,
			LastName:  pt.LastName,
			Email:     pt.Email,
			Phone:     *pt.Phone,
			MRN:       pt.MRN,
			DOB:       pt.DOB,
			Gender:    pt.Gender,
			Address:   pt.Address,
		})
	}
	return pts
}
