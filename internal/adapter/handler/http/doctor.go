package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/handler/http/dto"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
	"github.com/jenish-brainztechs/go-backend/internal/core/port"
)

type DoctorHandler struct {
	svc port.DoctorService
}

func NewDoctorHandler(svc port.DoctorService) *DoctorHandler {
	return &DoctorHandler{
		svc: svc,
	}
}

func (dh *DoctorHandler) CreateDoctor(ctx *gin.Context) {
	var req dto.CreateDoctor

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	doc := &domain.Doctor{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Phone:          &req.Phone,
		Email:          req.Email,
		RegistrationNo: req.RegistrationNo,
		Qualification:  req.Qualification,
	}

	d, err := dh.svc.CreateDoctor(ctx, doc)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, d)
}

func (dh *DoctorHandler) GetDoctorByID(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := uuid.Parse(id)

	if err != nil {
		parseError(err)
		return
	}

	doc, err := dh.svc.GetDoctorByID(ctx, uid)

	if err != nil {
		handleError(ctx, err)
		return
	}
	handleSuccess(ctx, doc)
}

func (dh *DoctorHandler) GetDoctors(ctx *gin.Context) {
	res, err := dh.svc.GetDoctors(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := dto.DocResponses(res)

	handleSuccess(ctx, rsp)
}

func (dh *DoctorHandler) UpdateDoctor(ctx *gin.Context) {
	ID := ctx.Param("id")
	docID, err := uuid.Parse(ID)
	if err != nil {
		handleError(ctx, domain.ErrInvalidUUID)
		return
	}

	var req dto.CreateDoctor
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

	doc := &domain.Doctor{
		ID:             docID,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Phone:          &req.Phone,
		Email:          req.Email,
		RegistrationNo: req.RegistrationNo,
		Qualification:  req.Qualification,
		UpdatedBy:      userPayload.UserId,
	}

	err = dh.svc.UpdateDoctor(ctx, doc)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, gin.H{"message": "Doctor updated successfully"})
}
