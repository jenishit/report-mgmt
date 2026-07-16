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

// CreateDepartment creates a new department
// @Summary Create department
// @Description Create a new lab department (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.DepartmentRequest true "Department details"
// @Success 200 {object} response{data=domain.Department}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/create-department [post]
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

// CreatePanel creates a new panel
// @Summary Create panel
// @Description Create a new test panel (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.PanelRequest true "Panel details"
// @Success 200 {object} response{data=domain.Panels}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/create-panel [post]
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

// CreateTestCatalog creates a new test catalog entry
// @Summary Create test catalog
// @Description Create a new test catalog entry (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TestCatalogRequest true "Test catalog details"
// @Success 200 {object} response{data=domain.TestCatalog}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/create-catalog [post]
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

// CreatePanelComponent creates a panel-component association
// @Summary Create panel component
// @Description Associate a test catalog with a panel (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.PanelComponentRequest true "Panel component details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/create-panel-component [post]
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

// CreateTestParameter creates a new test parameter
// @Summary Create test parameter
// @Description Create a new test parameter for a test catalog (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TestParameterRequest true "Test parameter details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/create-test-parameter [post]
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

// CreateReferenceRange creates a new reference range
// @Summary Create reference range
// @Description Create a new reference range for a test parameter (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ReferenceRangeRequest true "Reference range details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/create-reference [post]
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

// GetDepartments returns all departments
// @Summary List departments
// @Description Get all lab departments
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response{data=[]dto.DepartmentResponse}
// @Failure 401 {object} errorResponse
// @Router /lab-test/list-department [get]
func (lth *LabTestsHandler) GetDepartments(ctx *gin.Context) {
	res, err := lth.svc.GetDepartments(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.DeptsResponse(res)

	handleSuccess(ctx, rsp)
}

// GetPanelsByDepartmentID returns panels for a department
// @Summary List panels by department
// @Description Get all panels for a given department
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Department ID"
// @Success 200 {object} response{data=[]dto.PanelResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /lab-test/list-panel/{id} [get]
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

// GetTestCatalogByDepartmentID returns test catalogs for a department
// @Summary List test catalogs by department
// @Description Get all test catalogs for a given department
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Department ID"
// @Success 200 {object} response{data=[]dto.TestCatalogResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /lab-test/list-catalog/{id} [get]
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

// GetPanelComponentsByPanelID returns panel components for a panel
// @Summary List panel components by panel
// @Description Get all panel components for a given panel
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Panel ID"
// @Success 200 {object} response{data=[]dto.PanelComponentResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /lab-test/list-panel-catalog/{id} [get]
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

// GetTestParametersByTestCatalogID returns test parameters for a test catalog
// @Summary List test parameters by catalog
// @Description Get all test parameters for a given test catalog
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Test Catalog ID"
// @Success 200 {object} response{data=[]dto.TestParameterResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /lab-test/list-test-parameter/{id} [get]
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

// GetReferenceRangesByTestParameterID returns reference ranges for a test parameter
// @Summary List reference ranges by parameter
// @Description Get all reference ranges for a given test parameter
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Test Parameter ID"
// @Success 200 {object} response{data=[]dto.ReferenceRangeResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /lab-test/list-reference/{id} [get]
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

// UpdateDepartment updates a department
// @Summary Update department
// @Description Update an existing department (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.DepartmentRequest true "Department details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/update-department [patch]
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

// UpdatePanel updates a panel
// @Summary Update panel
// @Description Update an existing panel (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.PanelRequest true "Panel details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/update-panel [patch]
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

// UpdateTestCatalog updates a test catalog
// @Summary Update test catalog
// @Description Update an existing test catalog (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TestCatalogRequest true "Test catalog details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/update-catalog [patch]
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

// UpdateTestParameter updates a test parameter
// @Summary Update test parameter
// @Description Update an existing test parameter (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.TestParameterRequest true "Test parameter details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/update-test-parameter [patch]
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

// UpdateReferenceRange updates a reference range
// @Summary Update reference range
// @Description Update an existing reference range (admin only)
// @Tags Lab Tests
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ReferenceRangeRequest true "Reference range details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /admin/lab-test/update-reference [patch]
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
