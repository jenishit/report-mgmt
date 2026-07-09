package dto

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type DepartmentRequest struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"dept_name" binding:"required"`
	Description string    `json:"dept_description"`
}

type PanelRequest struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"panel_name" binding:"required"`
	DepartmentID uuid.UUID `json:"dept_id" binding:"required"`
	Code         string    `json:"panel_code"`
	PanelPrice   float64   `json:"panel_price"`
}

type TestCatalogRequest struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"test_name" binding:"required"`
	DepartmentID   uuid.UUID `json:"dept_id" binding:"required"`
	Code           string    `json:"test_code"`
	TestPrice      float64   `json:"test_price"`
	SampleType     string    `json:"sample_type"`
	TurnaroundTime int       `json:"turnaround_time"`
}

type PanelComponentRequest struct {
	PanelID    uuid.UUID `json:"panel_id" binding:"required"`
	TestID     uuid.UUID `json:"test_id" binding:"required"`
	SequenceNo int       `json:"sequence_no"`
}

type TestParameterRequest struct {
	ID         uuid.UUID `json:"id"`
	TestID     uuid.UUID `json:"test_id" binding:"required"`
	Name       string    `json:"parameter_name" binding:"required"`
	Unit       string    `json:"unit"`
	ResultType string    `json:"result_type"`
	SequenceNo int       `json:"sequence_no"`
}

type ReferenceRangeRequest struct {
	ID          uuid.UUID `json:"id"`
	ParameterID uuid.UUID `json:"parameter_id" binding:"required"`
	Gender      string    `json:"gender"`
	MinAge      int       `json:"min_age"`
	MaxAge      int       `json:"max_age"`
	MinValue    float64   `json:"min_value"`
	MaxValue    float64   `json:"max_value"`
	TextRange   string    `json:"text_range"`
	Note        string    `json:"note"`
}

type DepartmentResponse struct {
	ID          uuid.UUID `json:"dept_id"`
	Name        string    `json:"dept_name"`
	Description string    `json:"dept_description"`
}
type PanelResponse struct {
	ID           uuid.UUID `json:"panel_id"`
	Name         string    `json:"panel_name" binding:"required"`
	DepartmentID uuid.UUID `json:"dept_id" binding:"required"`
	Code         string    `json:"panel_code"`
	PanelPrice   float64   `json:"panel_price"`
}
type TestCatalogResponse struct {
	ID             uuid.UUID `json:"test_catalog_id"`
	Name           string    `json:"test_name"`
	DepartmentID   uuid.UUID `json:"dept_id"`
	Code           string    `json:"test_code"`
	TestPrice      float64   `json:"test_price"`
	SampleType     string    `json:"sample_type"`
	TurnaroundTime int       `json:"turnaround_time"`
}
type PanelComponentResponse struct {
	PanelID    uuid.UUID `json:"panel_id" binding:"required"`
	TestID     uuid.UUID `json:"test_id" binding:"required"`
	SequenceNo int       `json:"sequence_no"`
}
type TestParameterResponse struct {
	ID         uuid.UUID `json:"parameter_id"`
	TestID     uuid.UUID `json:"test_id" binding:"required"`
	Name       string    `json:"parameter_name" binding:"required"`
	Unit       string    `json:"unit"`
	ResultType string    `json:"result_type"`
	SequenceNo int       `json:"sequence_no"`
}
type ReferenceRangeResponse struct {
	ID          uuid.UUID `json:"ref_id"`
	ParameterID uuid.UUID `json:"parameter_id" binding:"required"`
	Gender      string    `json:"gender"`
	Age         string    `json:"age"`
	Value       string    `json:"value"`
	TextRange   string    `json:"text_range"`
	Note        string    `json:"note"`
}

func DeptResponse(d *domain.Department) *DepartmentResponse {
	return &DepartmentResponse{
		ID:          d.ID,
		Name:        d.Name,
		Description: *d.Description,
	}
}

func DeptsResponse(d []*domain.Department) []*DepartmentResponse {
	depts := make([]*DepartmentResponse, 0, len(d))

	for _, dept := range d {
		depts = append(depts, &DepartmentResponse{
			ID:          dept.ID,
			Name:        dept.Name,
			Description: *dept.Description,
		})
	}
	return depts
}

func GetPanelResponse(p *domain.Panels) *PanelResponse {
	return &PanelResponse{
		ID:           p.ID,
		Name:         p.Name,
		DepartmentID: p.DepartmentID,
		Code:         *p.Code,
		PanelPrice:   *p.PanelPrice,
	}
}

func GetPanelsResponse(p []*domain.Panels) []*PanelResponse {
	panels := make([]*PanelResponse, 0, len(p))

	for _, panel := range p {
		panels = append(panels, &PanelResponse{
			ID:           panel.ID,
			Name:         panel.Name,
			DepartmentID: panel.DepartmentID,
			Code:         *panel.Code,
			PanelPrice:   *panel.PanelPrice,
		})
	}

	return panels
}

func GetCatalogResponse(c *domain.TestCatalog) *TestCatalogResponse {
	return &TestCatalogResponse{
		ID:             c.ID,
		DepartmentID:   c.DepartmentID,
		Name:           c.Name,
		Code:           *c.Code,
		TestPrice:      *c.TestPrice,
		SampleType:     *c.SampleType,
		TurnaroundTime: *c.TurnAroundTime,
	}
}

func GetCatalogsResponse(c []*domain.TestCatalog) []*TestCatalogResponse {
	catalogs := make([]*TestCatalogResponse, 0, len(c))

	for _, catalog := range c {
		catalogs = append(catalogs, &TestCatalogResponse{
			ID:             catalog.ID,
			DepartmentID:   catalog.DepartmentID,
			Name:           catalog.Name,
			Code:           *catalog.Code,
			TestPrice:      *catalog.TestPrice,
			SampleType:     *catalog.SampleType,
			TurnaroundTime: *catalog.TurnAroundTime,
		})
	}

	return catalogs
}

func GetPanelComponentResponse(pc *domain.PanelComponent) *PanelComponentResponse {
	return &PanelComponentResponse{
		PanelID:    pc.PanelID,
		TestID:     pc.TestCatalogID,
		SequenceNo: *pc.SequenceNo,
	}
}

func GetPanelComponentsResponse(pc []*domain.PanelComponent) []*PanelComponentResponse {
	components := make([]*PanelComponentResponse, 0, len(pc))

	for _, component := range pc {
		components = append(components, &PanelComponentResponse{
			PanelID:    component.PanelID,
			TestID:     component.TestCatalogID,
			SequenceNo: *component.SequenceNo,
		})
	}

	return components
}

func GetTestParameterResponse(tp *domain.TestParameter) *TestParameterResponse {
	return &TestParameterResponse{
		ID:         tp.ID,
		TestID:     tp.TestCatalogID,
		Name:       tp.Name,
		Unit:       tp.Unit,
		ResultType: tp.ResultType,
		SequenceNo: tp.SequenceNo,
	}
}

func GetTestParametersResponse(tp []*domain.TestParameter) []*TestParameterResponse {
	parameters := make([]*TestParameterResponse, 0, len(tp))

	for _, parameter := range tp {
		parameters = append(parameters, &TestParameterResponse{
			ID:         parameter.ID,
			TestID:     parameter.TestCatalogID,
			Name:       parameter.Name,
			Unit:       parameter.Unit,
			ResultType: parameter.ResultType,
			SequenceNo: parameter.SequenceNo,
		})
	}

	return parameters
}

func GetReferenceRangeResponse(rr *domain.ReferenceRange) *ReferenceRangeResponse {
	var age string
	if rr.MinAge != nil && rr.MaxAge != nil {
		age = fmt.Sprintf("%d-%d", *rr.MinAge, *rr.MaxAge)
	}

	var value string
	if rr.MinValue != nil && rr.MaxValue != nil {
		value = fmt.Sprintf("%g-%g", *rr.MinValue, *rr.MaxValue)
	}

	var textRange string
	if rr.TextValue != nil {
		textRange = *rr.TextValue
	}

	return &ReferenceRangeResponse{
		ID:          rr.ID,
		ParameterID: rr.TestParameterID,
		Gender:      rr.Gender,
		Age:         age,
		Value:       value,
		TextRange:   textRange,
	}
}

func GetReferenceRangesResponse(rr []*domain.ReferenceRange) []*ReferenceRangeResponse {
	ranges := make([]*ReferenceRangeResponse, 0, len(rr))

	for _, r := range rr {
		var age string
		var value string
		var textRange string

		if r.MinAge != nil && r.MaxAge != nil {
			age = fmt.Sprintf("%d-%d", *r.MinAge, *r.MaxAge)
		}

		if r.MinValue != nil && r.MaxValue != nil {
			value = fmt.Sprintf("%g-%g", *r.MinValue, *r.MaxValue)
		}

		if r.TextValue != nil {
			textRange = *r.TextValue
		}

		ranges = append(ranges, &ReferenceRangeResponse{
			ID:          r.ID,
			ParameterID: r.TestParameterID,
			Gender:      r.Gender,
			Age:         age,
			Value:       value,
			TextRange:   textRange,
		})
	}

	return ranges
}
