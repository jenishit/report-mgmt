package dto

import "github.com/google/uuid"

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
