package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type CreateVisit struct {
	VisitNo   string     `json:"visit_no" binding:"required"`
	PatientID uuid.UUID  `json:"patient_id" binding:"required"`
	DoctorID  *uuid.UUID `json:"doctor_id"`
	Status    string     `json:"status"`
}

type ListVisits struct {
	ID           uuid.UUID `json:"visit_id"`
	VisitNo      string    `json:"visit_no"`
	PatientName  string    `json:"patient_name"`
	DoctorName   string    `json:"doctor_name"`
	RegisterName string    `json:"registered_by"`
	VisitDate    time.Time `json:"visit_date"`
	Status       string    `json:"status"`
}

func VisitResponse(v *domain.ListVisit) *ListVisits {
	return &ListVisits{
		ID:           v.ID,
		VisitNo:      v.VisitNo,
		PatientName:  v.PatientFirstName + " " + v.PatientLastName,
		DoctorName:   v.DoctorFirstName + " " + v.DoctorLastName,
		RegisterName: v.RegisterFirstName + " " + v.RegisterLastName,
		Status:       string(v.Status),
	}
}

func VisitsResponse(vs []*domain.ListVisit) []*ListVisits {
	visits := make([]*ListVisits, 0, len(vs))

	for _, v := range vs {
		visits = append(visits, &ListVisits{
			ID:           v.ID,
			VisitNo:      v.VisitNo,
			PatientName:  v.PatientFirstName + " " + v.PatientLastName,
			DoctorName:   v.DoctorFirstName + " " + v.DoctorLastName,
			RegisterName: v.RegisterFirstName + " " + v.RegisterLastName,
			Status:       string(v.Status),
		})
	}
	return visits
}
