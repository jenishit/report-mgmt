package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type LabTestsHandler struct {
	svc port.LabTestsService
}

func NewLabTestsHandler(svc port.LabTestsService) *LabTestsHandler {
	return &LabTestsHandler{
		svc: svc,
	}
}

func (lth *LabTestsHandler) CreateDepartment(ctx *gin.Context) {
	var req dto.DepartmentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	dept := &domain.Department{
		Name:        req.Name,
		Description: &req.Description,
	}

	d, err := lth.svc.CreateDepartment(ctx, dept)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, d)
}

func (lth *LabTestsHandler) CreatePanel(ctx *gin.Context) {
	var req dto.PanelRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	panel := &domain.Panels{
		Name:         req.Name,
		DepartmentID: req.DepartmentID,
		Code:         &req.Code,
		PanelPrice:   &req.PanelPrice,
	}

	panel, err := lth.svc.CreatePanel(ctx, panel)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, panel)
}

func (lth *LabTestsHandler) CreateTestCatalog(ctx *gin.Context) {
	var req dto.TestCatalogRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	catalog := &domain.TestCatalog{
		DepartmentID:   req.DepartmentID,
		Name:           req.Name,
		Code:           &req.Code,
		SampleType:     &req.SampleType,
		TestPrice:      &req.TestPrice,
		TurnAroundTime: &req.TurnaroundTime,
	}

	cat, err := lth.svc.CreateTestCatalog(ctx, catalog)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, cat)
}

func (lth *LabTestsHandler) CreatePanelComponent(ctx *gin.Context) {
	var req dto.PanelComponentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	component := &domain.PanelComponent{
		PanelID:       req.PanelID,
		TestCatalogID: req.TestID,
		SequenceNo:    &req.SequenceNo,
	}

	err := lth.svc.CreatePanelComponent(ctx, component)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Panel components created successfully"})
}

func (lth *LabTestsHandler) CreateTestParameter(ctx *gin.Context) {
	var req dto.TestParameterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	parameter := &domain.TestParameter{
		TestCatalogID: req.TestID,
		Name:          req.Name,
		Unit:          req.Unit,
		ResultType:    req.ResultType,
		SequenceNo:    req.SequenceNo,
	}

	param, err := lth.svc.CreateTestParameter(ctx, parameter)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Test parameter created successfully", "data": param})
}

func (lth *LabTestsHandler) CreateReferenceRange(ctx *gin.Context) {
	var req dto.ReferenceRangeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	ref := &domain.ReferenceRange{
		TestParameterID: req.ParameterID,
		Gender:          req.Gender,
		MinAge:          &req.MinAge,
		MaxAge:          &req.MaxAge,
		MinValue:        &req.MinValue,
		MaxValue:        &req.MaxValue,
		TextValue:       &req.TextRange,
	}

	err := lth.svc.CreateReferenceRange(ctx, ref)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Successfully created reference range"})
}

func (lth *LabTestsHandler) GetDepartments(ctx *gin.Context) {
	res, err := lth.svc.GetDepartments(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.DeptsResponse(res)

	handleSuccess(ctx, rsp)
}

func (lth *LabTestsHandler) GetPanelsByDepartmentID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	res, err := lth.svc.GetPanelsByDepartmentID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.GetPanelsResponse(res)

	handleSuccess(ctx, rsp)
}

func (lth *LabTestsHandler) GetTestCatalogByDepartmentID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	res, err := lth.svc.GetTestCatalogByDepartmentID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.GetCatalogsResponse(res)

	handleSuccess(ctx, rsp)
}

func (lth *LabTestsHandler) GetPanelComponentsByPanelID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	res, err := lth.svc.GetPanelComponentsByPanelID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.GetPanelComponentsResponse(res)

	handleSuccess(ctx, rsp)
}

func (lth *LabTestsHandler) GetTestParametersByTestCatalogID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	res, err := lth.svc.GetTestParametersByTestCatalogID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.GetTestParametersResponse(res)

	handleSuccess(ctx, rsp)
}

func (lth *LabTestsHandler) GetReferenceRangesByTestParameterID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	res, err := lth.svc.GetReferenceRangesByTestParameterID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.GetReferenceRangesResponse(res)

	handleSuccess(ctx, rsp)
}

func (lth *LabTestsHandler) UpdateDepartment(ctx *gin.Context) {
	var req dto.DepartmentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	dept := &domain.Department{
		ID:          req.ID,
		Name:        req.Name,
		Description: &req.Description,
	}

	err := lth.svc.UpdateDepartment(ctx, dept)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Department Updated successfully"})
}

func (lth *LabTestsHandler) UpdatePanel(ctx *gin.Context) {
	var req dto.PanelRequest
	now := time.Now()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	panel := &domain.Panels{
		ID:           req.ID,
		Name:         req.Name,
		DepartmentID: req.DepartmentID,
		Code:         &req.Code,
		PanelPrice:   &req.PanelPrice,
		UpdatedAt:    now,
	}

	err := lth.svc.UpdatePanel(ctx, panel)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Panel updated"})
}

func (lth *LabTestsHandler) UpdateTestCatalog(ctx *gin.Context) {
	var req dto.TestCatalogRequest
	now := time.Now()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	catalog := &domain.TestCatalog{
		DepartmentID:   req.DepartmentID,
		Name:           req.Name,
		Code:           &req.Code,
		SampleType:     &req.SampleType,
		TestPrice:      &req.TestPrice,
		TurnAroundTime: &req.TurnaroundTime,
		UpdatedAt:      now,
	}

	err := lth.svc.UpdateTestCatalog(ctx, catalog)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Test catalog updated successfully"})
}

func (lth *LabTestsHandler) UpdateTestParameter(ctx *gin.Context) {
	var req dto.TestParameterRequest
	now := time.Now()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	parameter := &domain.TestParameter{
		TestCatalogID: req.TestID,
		Name:          req.Name,
		Unit:          req.Unit,
		ResultType:    req.ResultType,
		SequenceNo:    req.SequenceNo,
		UpdatedAt:     now,
	}

	err := lth.svc.UpdateTestParameter(ctx, parameter)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Test Parameter updated successfully"})
}

func (lth *LabTestsHandler) UpdateReferenceRange(ctx *gin.Context) {
	var req dto.ReferenceRangeRequest
	now := time.Now()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	ref := &domain.ReferenceRange{
		TestParameterID: req.ParameterID,
		Gender:          req.Gender,
		MinAge:          &req.MinAge,
		MaxAge:          &req.MaxAge,
		MinValue:        &req.MinValue,
		MaxValue:        &req.MaxValue,
		TextValue:       &req.TextRange,
		UpdatedAt:       now,
	}

	err := lth.svc.UpdateReferenceRange(ctx, ref)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Reference range is updated"})
}
