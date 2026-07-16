package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type CreateReportRequest struct {
	ReportNo string    `json:"report_no" binding:"required"`
	VisitID  uuid.UUID `json:"visit_id"`
}

type UpdateReportRequest struct {
	Status  domain.ReportStatus `json:"status"`
	PDFPath string              `json:"pdf_path"`
}

type ReportResponse struct {
	ID          uuid.UUID          `json:"id"`
	ReportNo    string             `json:"report_no"`
	VisitID     *uuid.UUID         `json:"visit_id"`
	GeneratedBy uuid.UUID          `json:"generated_by"`
	GeneratedAt time.Time          `json:"generated_at"`
	Status      domain.ReportStatus `json:"status"`
	PDFPath     *string            `json:"pdf_path"`
}

type ReportPrintRequest struct {
	CopyNumber int `json:"copy_number" binding:"required"`
}

type ReportPrintResponse struct {
	ID         uuid.UUID `json:"id"`
	ReportID   uuid.UUID `json:"report_id"`
	PrintedBy  uuid.UUID `json:"printed_by"`
	PrintedAt  time.Time `json:"printed_at"`
	CopyNumber int       `json:"copy_number"`
}

func ReportResponseFromDomain(r *domain.Report) *ReportResponse {
	resp := &ReportResponse{
		ID:          r.ID,
		ReportNo:    r.ReportNo,
		GeneratedBy: r.GeneratedBy,
		GeneratedAt: r.GeneratedAt,
		Status:      r.Status,
	}
	if r.VisitID != uuid.Nil {
		resp.VisitID = &r.VisitID
	}
	if r.PDFPath != "" {
		resp.PDFPath = &r.PDFPath
	}
	return resp
}

func ReportsResponseFromDomain(reports []*domain.Report) []*ReportResponse {
	res := make([]*ReportResponse, 0, len(reports))
	for _, r := range reports {
		res = append(res, ReportResponseFromDomain(r))
	}
	return res
}

func ReportPrintResponseFromDomain(p *domain.ReportPrint) *ReportPrintResponse {
	return &ReportPrintResponse{
		ID:         p.ID,
		ReportID:   p.ReportID,
		PrintedBy:  p.PrintedBy,
		PrintedAt:  p.PrintedAt,
		CopyNumber: p.CopyNumber,
	}
}

func ReportPrintsResponseFromDomain(prints []*domain.ReportPrint) []*ReportPrintResponse {
	res := make([]*ReportPrintResponse, 0, len(prints))
	for _, p := range prints {
		res = append(res, ReportPrintResponseFromDomain(p))
	}
	return res
}
