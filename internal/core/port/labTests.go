package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type LabTestsRepository interface {
	CreateDepartment(ctx context.Context, dept *domain.Department) (*domain.Department, error)
	CreatePanel(ctx context.Context, panel *domain.Panels) (*domain.Panels, error)
	CreateTestCatalog(ctx context.Context, test *domain.TestCatalog) (*domain.TestCatalog, error)
	CreatePanelComponent(ctx context.Context, component *domain.PanelComponent) error
	CreateTestParameter(ctx context.Context, parameter *domain.TestParameter) (*domain.TestParameter, error)
	CreateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error
	GetDepartments(ctx context.Context) ([]*domain.Department, error)
	GetPanelsByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.Panels, error)
	GetTestCatalogByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.TestCatalog, error)
	GetPanelComponentsByPanelID(ctx context.Context, panelID uuid.UUID) ([]*domain.PanelComponent, error)
	GetTestParametersByTestCatalogID(ctx context.Context, testCatalogID uuid.UUID) ([]*domain.TestParameter, error)
	GetReferenceRangesByTestParameterID(ctx context.Context, testParameterID uuid.UUID) ([]*domain.ReferenceRange, error)
	UpdateDepartment(ctx context.Context, dept *domain.Department) error
	UpdatePanel(ctx context.Context, panel *domain.Panels) error
	UpdateTestCatalog(ctx context.Context, test *domain.TestCatalog) error
	UpdateTestParameter(ctx context.Context, parameter *domain.TestParameter) error
	UpdateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error
}

type LabTestsService interface {
	CreateDepartment(ctx context.Context, dept *domain.Department) (*domain.Department, error)
	CreatePanel(ctx context.Context, panel *domain.Panels) (*domain.Panels, error)
	CreateTestCatalog(ctx context.Context, test *domain.TestCatalog) (*domain.TestCatalog, error)
	CreatePanelComponent(ctx context.Context, component *domain.PanelComponent) error
	CreateTestParameter(ctx context.Context, parameter *domain.TestParameter) (*domain.TestParameter, error)
	CreateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error
	GetDepartments(ctx context.Context) ([]*domain.Department, error)
	GetPanelsByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.Panels, error)
	GetTestCatalogByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.TestCatalog, error)
	GetPanelComponentsByPanelID(ctx context.Context, panelID uuid.UUID) ([]*domain.PanelComponent, error)
	GetTestParametersByTestCatalogID(ctx context.Context, testCatalogID uuid.UUID) ([]*domain.TestParameter, error)
	GetReferenceRangesByTestParameterID(ctx context.Context, testParameterID uuid.UUID) ([]*domain.ReferenceRange, error)
	UpdateDepartment(ctx context.Context, dept *domain.Department) error
	UpdatePanel(ctx context.Context, panel *domain.Panels) error
	UpdateTestCatalog(ctx context.Context, test *domain.TestCatalog) error
	UpdateTestParameter(ctx context.Context, parameter *domain.TestParameter) error
	UpdateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error
}
