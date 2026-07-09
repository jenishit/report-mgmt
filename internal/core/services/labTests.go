package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type LabTestsService struct {
	repo port.LabTestsRepository
}

func NewLabTestsService(lts port.LabTestsRepository) *LabTestsService {
	return &LabTestsService{
		repo: lts,
	}
}

func (lts *LabTestsService) CreateDepartment(ctx context.Context, dept *domain.Department) (*domain.Department, error) {
	return lts.repo.CreateDepartment(ctx, dept)
}

func (lts *LabTestsService) CreatePanel(ctx context.Context, panel *domain.Panels) (*domain.Panels, error) {
	return lts.repo.CreatePanel(ctx, panel)
}
func (lts *LabTestsService) CreateTestCatalog(ctx context.Context, test *domain.TestCatalog) (*domain.TestCatalog, error) {
	return lts.repo.CreateTestCatalog(ctx, test)
}
func (lts *LabTestsService) CreatePanelComponent(ctx context.Context, component *domain.PanelComponent) error {
	return lts.repo.CreatePanelComponent(ctx, component)
}
func (lts *LabTestsService) CreateTestParameter(ctx context.Context, parameter *domain.TestParameter) (*domain.TestParameter, error) {
	return lts.repo.CreateTestParameter(ctx, parameter)
}
func (lts *LabTestsService) CreateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error {
	return lts.repo.CreateReferenceRange(ctx, rangeData)
}
func (lts *LabTestsService) GetDepartments(ctx context.Context) ([]*domain.Department, error) {
	return lts.repo.GetDepartments(ctx)
}
func (lts *LabTestsService) GetPanelsByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.Panels, error) {
	return lts.repo.GetPanelsByDepartmentID(ctx, departmentID)
}
func (lts *LabTestsService) GetTestCatalogByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.TestCatalog, error) {
	return lts.repo.GetTestCatalogByDepartmentID(ctx, departmentID)
}
func (lts *LabTestsService) GetPanelComponentsByPanelID(ctx context.Context, panelID uuid.UUID) ([]*domain.PanelComponent, error) {
	return lts.repo.GetPanelComponentsByPanelID(ctx, panelID)
}
func (lts *LabTestsService) GetTestParametersByTestCatalogID(ctx context.Context, testCatalogID uuid.UUID) ([]*domain.TestParameter, error) {
	return lts.repo.GetTestParametersByTestCatalogID(ctx, testCatalogID)
}
func (lts *LabTestsService) GetReferenceRangesByTestParameterID(ctx context.Context, testParameterID uuid.UUID) ([]*domain.ReferenceRange, error) {
	return lts.repo.GetReferenceRangesByTestParameterID(ctx, testParameterID)
}
func (lts *LabTestsService) UpdateDepartment(ctx context.Context, dept *domain.Department) error {
	return lts.repo.UpdateDepartment(ctx, dept)
}
func (lts *LabTestsService) UpdatePanel(ctx context.Context, panel *domain.Panels) error {
	return lts.repo.UpdatePanel(ctx, panel)
}
func (lts *LabTestsService) UpdateTestCatalog(ctx context.Context, test *domain.TestCatalog) error {
	return lts.repo.UpdateTestCatalog(ctx, test)
}
func (lts *LabTestsService) UpdateTestParameter(ctx context.Context, parameter *domain.TestParameter) error {
	return lts.repo.UpdateTestParameter(ctx, parameter)
}
func (lts *LabTestsService) UpdateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error {
	return lts.repo.UpdateReferenceRange(ctx, rangeData)
}
