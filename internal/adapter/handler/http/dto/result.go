package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type CreateResultRequest struct {
	OrderID     uuid.UUID      `json:"order_id" binding:"required"`
	ParameterID uuid.UUID      `json:"parameter_id" binding:"required"`
	ResultValue string         `json:"result_value" binding:"required"`
	Flag        domain.Flag    `json:"flag"`
	VerifiedBy  uuid.UUID      `json:"verified_by"`
	Remarks     string         `json:"remarks"`
}

type BatchResultItem struct {
	ParameterID uuid.UUID      `json:"parameter_id" binding:"required"`
	ResultValue string         `json:"result_value" binding:"required"`
	Flag        domain.Flag    `json:"flag"`
	Remarks     string         `json:"remarks"`
}

type BatchCreateResultRequest struct {
	OrderID    uuid.UUID         `json:"order_id" binding:"required"`
	VerifiedBy uuid.UUID         `json:"verified_by"`
	Results    []BatchResultItem `json:"results" binding:"required,min=1,dive"`
}

type UpdateResultRequest struct {
	ResultValue string      `json:"result_value"`
	Flag        domain.Flag `json:"flag"`
	VerifiedBy  uuid.UUID   `json:"verified_by"`
	Remarks     string      `json:"remarks"`
}

type ResultResponse struct {
	ID          uuid.UUID      `json:"id"`
	OrderID     uuid.UUID      `json:"order_id"`
	ParameterID uuid.UUID      `json:"parameter_id"`
	ResultValue string         `json:"result_value"`
	Flag        domain.Flag    `json:"flag"`
	PerformedBy uuid.UUID      `json:"performed_by"`
	PerformedAt time.Time      `json:"performed_at"`
	VerifiedBy  uuid.UUID      `json:"verified_by"`
	Remarks     *string        `json:"remarks"`
}

func ResultResponseFromDomain(r *domain.Result) *ResultResponse {
	resp := &ResultResponse{
		ID:          r.ID,
		OrderID:     r.OrderID,
		ParameterID: r.ParameterID,
		ResultValue: r.ResultValue,
		Flag:        r.Flag,
		PerformedBy: r.PerformedBy,
		PerformedAt: r.PerformedAt,
		VerifiedBy:  r.VerifiedBy,
	}
	if r.Remarks != "" {
		resp.Remarks = &r.Remarks
	}
	return resp
}

func ResultsResponseFromDomain(results []*domain.Result) []*ResultResponse {
	res := make([]*ResultResponse, 0, len(results))
	for _, r := range results {
		res = append(res, ResultResponseFromDomain(r))
	}
	return res
}
