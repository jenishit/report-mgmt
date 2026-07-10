package http

import (
	"fmt"
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

func (ph *PatientHandler) CreatePatient(ctx *gin.Context) {
	var req dto.CreatePatient

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}
	fmt.Println("heree")
	dob, err := time.Parse("2006-01-02", req.DOB)

	fmt.Println(dob)

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

func (ph *PatientHandler) GetPatients(ctx *gin.Context) {
	res, err := ph.svc.GetPatients(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.PtResponses(res)

	handleSuccess(ctx, rsp)
}

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
