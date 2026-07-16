package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type DoctorService struct {
	repo port.DoctorRepository
}

func NewDoctorService(dr port.DoctorRepository) *DoctorService {
	return &DoctorService{
		repo: dr,
	}
}

func (d *DoctorService) CreateDoctor(ctx context.Context, doc *domain.Doctor) (*domain.Doctor, error) {
	return d.repo.CreateDoctor(ctx, doc)
}

func (d *DoctorService) GetDoctorByID(ctx context.Context, id uuid.UUID) (*domain.Doctor, error) {
	return d.repo.GetDoctorByID(ctx, id)
}
func (d *DoctorService) GetDoctors(ctx context.Context) ([]*domain.Doctor, error) {
	return d.repo.GetDoctors(ctx)
}
func (d *DoctorService) UpdateDoctor(ctx context.Context, doc *domain.Doctor) error {
	return d.repo.UpdateDoctor(ctx, doc)
}
