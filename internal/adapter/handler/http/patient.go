package http

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type PatientHandler struct {
	svc port.PatientService
}

func NewPatientHandler(svc port.PatientService) *PatientHandler {
	return &PatientHandler{
		svc: svc,
	}
}

// CreatePatient creates a new patient
// @Summary Create patient
// @Description Create a new patient record
// @Tags Patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreatePatient true "Patient details"
// @Success 200 {object} response{data=domain.Patient}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Router /patient [post]
func (ph *PatientHandler) CreatePatient(ctx *gin.Context) {
	var req dto.CreatePatient

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	dob, err := time.Parse("2006-01-02", req.DOB)

	if err != nil {
		handleError(ctx, err)
		return
	}

	pt := &domain.Patient{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     &req.Phone,
		Email:     req.Email,
		MRN:       req.MRN,
		DOB:       dob,
		Gender:    req.Gender,
		Address:   req.Address,
	}

	d, err := ph.svc.CreatePatient(ctx, pt)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, d)
}

// GetPatientByID returns a patient by ID
// @Summary Get patient by ID
// @Description Get a patient's details by ID
// @Tags Patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Patient ID"
// @Success 200 {object} response{data=dto.PatientResponse}
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /patient/{id} [get]
func (ph *PatientHandler) GetPatientByID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	pt, err := ph.svc.GetPatientByID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, pt)
}

// GetPatients returns all patients
// @Summary List patients
// @Description Get all patients
// @Tags Patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response{data=[]dto.PatientResponse}
// @Failure 401 {object} errorResponse
// @Router /patient [get]
func (ph *PatientHandler) GetPatients(ctx *gin.Context) {
	res, err := ph.svc.GetPatients(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.PtResponses(res)

	handleSuccess(ctx, rsp)
}

// UpdatePatient updates a patient
// @Summary Update patient
// @Description Update an existing patient record
// @Tags Patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Patient ID"
// @Param request body dto.CreatePatient true "Patient details"
// @Success 200 {object} response
// @Failure 400 {object} errorResponse
// @Failure 401 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /patient/{id} [patch]
func (ph *PatientHandler) UpdatePatient(ctx *gin.Context) {
	ID := ctx.Param("id")
	ptID, err := uuid.Parse(ID)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	var req dto.CreatePatient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	payload, exists := ctx.Get(authorizationPayloadKey)
	if !exists {
		validationError(ctx, domain.ErrEmptyAuthorizationHeader)
		return
	}

	userPayload, ok := payload.(*domain.TokenPayload)
	if !ok {
		validationError(ctx, domain.ErrInvalidAuthorizationHeader)
		return
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		log.Fatal(err)
	}

	pt := &domain.Patient{
		ID:        ptID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     &req.Phone,
		Email:     req.Email,
		MRN:       req.MRN,
		DOB:       dob,
		Gender:    req.Gender,
		Address:   req.Address,
		UpdatedBy: userPayload.UserId,
	}

	err = ph.svc.UpdatePatient(ctx, pt)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Patient updated successfully"})
}
