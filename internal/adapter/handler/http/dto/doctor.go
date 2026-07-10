package dto

import (
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type CreateDoctor struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	Email          string `json:"email" binding:"required,email"` //the acceptable string type is of email type
	Phone          string `json:"phone"`
	Qualification  string `json:"qualification"`
	RegistrationNo string `json:"registration_no" binding:"required"`
}

type DoctorResponse struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"` //the acceptable string type is of email type
	Phone          string    `json:"phone"`
	Qualification  string    `json:"qualification"`
	RegistrationNo string    `json:"registration_no"`
}

func DocResponse(d *domain.Doctor) *DoctorResponse {
	return &DoctorResponse{
		ID:             d.ID,
		FirstName:      d.FirstName,
		LastName:       d.LastName,
		Email:          d.Email,
		Phone:          *d.Phone,
		Qualification:  d.Qualification,
		RegistrationNo: d.RegistrationNo,
	}
}

func DocResponses(d []*domain.Doctor) []*DoctorResponse {
	docs := make([]*DoctorResponse, 0, len(d))

	for _, doc := range d {
		docs = append(docs, &DoctorResponse{
			ID:             doc.ID,
			FirstName:      doc.FirstName,
			LastName:       doc.LastName,
			Email:          doc.Email,
			Phone:          *doc.Phone,
			Qualification:  doc.Qualification,
			RegistrationNo: doc.RegistrationNo,
		})
	}
	return docs
}
