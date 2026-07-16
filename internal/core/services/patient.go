package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type PatientService struct {
	repo port.PatientRepository
}

func NewPatientService(pt port.PatientRepository) *PatientService {
	return &PatientService{
		repo: pt,
	}
}

func (d *PatientService) CreatePatient(ctx context.Context, pt *domain.Patient) (*domain.Patient, error) {
	return d.repo.CreatePatient(ctx, pt)
}

func (d *PatientService) GetPatientByID(ctx context.Context, id uuid.UUID) (*domain.Patient, error) {
	return d.repo.GetPatientByID(ctx, id)
}
func (d *PatientService) GetPatients(ctx context.Context) ([]*domain.Patient, error) {
	return d.repo.GetPatients(ctx)
}
func (d *PatientService) UpdatePatient(ctx context.Context, pt *domain.Patient) error {
	return d.repo.UpdatePatient(ctx, pt)
}
