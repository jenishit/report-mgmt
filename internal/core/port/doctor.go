package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type DoctorRepository interface {
	CreateDoctor(ctx context.Context, doc *domain.Doctor) (*domain.Doctor, error)
	GetDoctorByID(ctx context.Context, id uuid.UUID) (*domain.Doctor, error)
	GetDoctors(ctx context.Context) ([]*domain.Doctor, error)
	UpdateDoctor(ctx context.Context, doc *domain.Doctor) error
}

type DoctorService interface {
	CreateDoctor(ctx context.Context, doc *domain.Doctor) (*domain.Doctor, error)
	GetDoctorByID(ctx context.Context, id uuid.UUID) (*domain.Doctor, error)
	GetDoctors(ctx context.Context) ([]*domain.Doctor, error)
	UpdateDoctor(ctx context.Context, doc *domain.Doctor) error
}
