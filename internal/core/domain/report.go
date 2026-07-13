package domain

import (
	"time"

	"github.com/google/uuid"
)

type ReportStatus string

const (
	ReportDraft   ReportStatus = "draft"
	ReportFinal   ReportStatus = "final"
	ReportAmended ReportStatus = "amended"
)

type Report struct {
	ID          uuid.UUID
	ReportNo    string
	VisitID     uuid.UUID
	GeneratedBy uuid.UUID
	GeneratedAt time.Time
	Status      ReportStatus
	PDFPath     string
}

type ReportPrint struct {
	ID         uuid.UUID
	ReportID   uuid.UUID
	PrintedBy  uuid.UUID
	PrintedAt  time.Time
	CopyNumber int
}
