package dto

import (
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type LabRequest struct {
	LabName        string `json:"lab_name" binding:"required"`
	Tagline        string `json:"tagline"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	RegistrationNo string `json:"registration_no" binding:"required"`
	ReportFooter   string `json:"report_footer"`
}

type LabResponse struct {
	ID             uuid.UUID `json:"id"`
	LabName        string    `json:"lab_name"`
	Tagline        string    `json:"tagline"`
	Address        string    `json:"address"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	RegistrationNo string    `json:"registration_no"`
	ReportFooter   string    `json:"report_footer"`
	UpdatedBy      uuid.UUID `json:"updated_by"`
}

func LabsResponse(l *domain.LabSettings) *LabResponse {
	return &LabResponse{
		ID:             l.ID,
		LabName:        l.LabName,
		Tagline:        *l.Tagline,
		Address:        *l.Address,
		Phone:          *l.Phone,
		Email:          *l.Email,
		RegistrationNo: *l.RegistrationNo,
		ReportFooter:   *l.ReportFooter,
		UpdatedBy:      l.UpdatedBy,
	}
}

func LabsResponses(l []*domain.LabSettings) []*LabResponse {
	labs := make([]*LabResponse, 0, len(l))

	for _, lab := range l {
		labs = append(labs, &LabResponse{
			ID:             lab.ID,
			LabName:        lab.LabName,
			Tagline:        *lab.Tagline,
			Address:        *lab.Address,
			Phone:          *lab.Phone,
			Email:          *lab.Email,
			RegistrationNo: *lab.RegistrationNo,
			ReportFooter:   *lab.ReportFooter,
			UpdatedBy:      lab.UpdatedBy,
		})
	}
	return labs
}
