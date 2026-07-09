package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type LabTestsRepository struct {
	DB *postgres.DB
}

func NewLabTestsRepository(db *postgres.DB) *LabTestsRepository {
	return &LabTestsRepository{
		DB: db,
	}
}

func (lr *LabTestsRepository) CreateDepartment(ctx context.Context, dept *domain.Department) (*domain.Department, error) {
	query, args, err := sq.
		Insert("departments").
		Columns(
			"name",
			"description",
		).
		Values(
			dept.Name,
			dept.Description,
		).Suffix(`
		RETURNING
			id,
			name
	`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var depts domain.Department

	err = lr.DB.QueryRow(ctx, query, args...).Scan(
		&depts.ID,
		&depts.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("inserting department: %w", err)
	}

	return &depts, nil
}

func (lr *LabTestsRepository) CreatePanel(ctx context.Context, panel *domain.Panels) (*domain.Panels, error) {
	query, args, err := sq.
		Insert("panels").
		Columns(
			"department_id",
			"name",
			"code",
			"panel_price",
		).
		Values(
			panel.DepartmentID,
			panel.Name,
			panel.Code,
			panel.PanelPrice,
		).Suffix(`
		RETURNING
			id,
			name
	`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var panels domain.Panels

	err = lr.DB.QueryRow(ctx, query, args...).Scan(
		&panels.ID,
		&panels.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("inserting lab Panel: %w", err)
	}

	return &panels, nil
}

func (lr *LabTestsRepository) CreateTestCatalog(ctx context.Context, test *domain.TestCatalog) (*domain.TestCatalog, error) {
	query, args, err := sq.
		Insert("test_catalog").
		Columns(
			"department_id",
			"name",
			"code",
			"sample_type",
			"price",
			"turnaround_hours",
		).
		Values(
			test.DepartmentID,
			test.Name,
			test.Code,
			test.SampleType,
			test.TestPrice,
			test.TurnAroundTime,
		).Suffix(`
		RETURNING
			id,
			name
	`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var tests domain.TestCatalog

	err = lr.DB.QueryRow(ctx, query, args...).Scan(
		&tests.ID,
		&tests.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("inserting TestCatalog: %w", err)
	}

	return &tests, nil
}

func (lr *LabTestsRepository) CreatePanelComponent(ctx context.Context, component *domain.PanelComponent) error {
	query, args, err := sq.
		Insert("panel_components").
		Columns(
			"panel_id",
			"test_id",
			"sequence_no",
		).
		Values(
			component.PanelID,
			component.TestCatalogID,
			component.SequenceNo,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = lr.DB.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("inserting PanelComponent: %w", err)
	}

	return nil
}

func (lr *LabTestsRepository) CreateTestParameter(ctx context.Context, parameter *domain.TestParameter) (*domain.TestParameter, error) {
	query, args, err := sq.
		Insert("test_parameters").
		Columns(
			"test_id",
			"name",
			"unit",
			"result_type",
			"sequence_no",
		).
		Values(
			parameter.TestCatalogID,
			parameter.Name,
			parameter.Unit,
			parameter.ResultType,
			parameter.SequenceNo,
		).Suffix(`
		RETURNING
			id,
			name
	`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var params domain.TestParameter

	err = lr.DB.QueryRow(ctx, query, args...).Scan(
		&params.ID,
		&params.Name,
	)

	if err != nil {
		return nil, fmt.Errorf("inserting TestParameter: %w", err)
	}

	return &params, nil
}

func (lr *LabTestsRepository) CreateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error {
	query, args, err := sq.
		Insert("reference_ranges").
		Columns(
			"parameter_id",
			"gender",
			"age_min_years",
			"age_max_years",
			"low_value",
			"high_value",
			"text_range",
		).
		Values(
			rangeData.TestParameterID,
			rangeData.Gender,
			rangeData.MinAge,
			rangeData.MaxAge,
			rangeData.MinValue,
			rangeData.MaxValue,
			rangeData.TextValue,
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	_, err = lr.DB.Exec(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("inserting ReferenceRange: %w", err)
	}

	return nil
}

func (lr *LabTestsRepository) GetDepartments(ctx context.Context) ([]*domain.Department, error) {
	query, args, err := sq.
		Select(
			"id",
			"name",
			"description",
		).From("departments").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := lr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying departments: %w", err)
	}
	defer rows.Close()

	var departments []*domain.Department

	for rows.Next() {
		var dept domain.Department
		err := rows.Scan(
			&dept.ID,
			&dept.Name,
			&dept.Description,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning department row: %w", err)
		}
		departments = append(departments, &dept)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating department rows: %w", err)
	}

	return departments, nil
}

func (lr *LabTestsRepository) GetPanelsByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.Panels, error) {
	query, args, err := sq.
		Select(
			"id",
			"department_id",
			"name",
			"code",
			"panel_price",
			"is_active",
		).From("panels").
		Where(sq.Eq{"department_id": departmentID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := lr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying panels: %w", err)
	}
	defer rows.Close()

	var panels []*domain.Panels

	for rows.Next() {
		var panel domain.Panels
		err := rows.Scan(
			&panel.ID,
			&panel.DepartmentID,
			&panel.Name,
			&panel.Code,
			&panel.PanelPrice,
			&panel.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning panel row: %w", err)
		}
		panels = append(panels, &panel)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating panel rows: %w", err)
	}

	return panels, nil
}

func (lr *LabTestsRepository) GetTestCatalogByDepartmentID(ctx context.Context, departmentID uuid.UUID) ([]*domain.TestCatalog, error) {
	query, args, err := sq.
		Select(
			"id",
			"department_id",
			"name",
			"code",
			"sample_type",
			"price",
			"turnaround_hours",
			"is_active",
		).From("test_catalog").
		Where(sq.Eq{"department_id": departmentID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := lr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying test catalog: %w", err)
	}
	defer rows.Close()

	var testCatalogs []*domain.TestCatalog

	for rows.Next() {
		var test domain.TestCatalog
		err := rows.Scan(
			&test.ID,
			&test.DepartmentID,
			&test.Name,
			&test.Code,
			&test.SampleType,
			&test.TestPrice,
			&test.TurnAroundTime,
			&test.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning test catalog row: %w", err)
		}
		testCatalogs = append(testCatalogs, &test)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating test catalog rows: %w", err)
	}

	return testCatalogs, nil
}

func (lr *LabTestsRepository) GetPanelComponentsByPanelID(ctx context.Context, panelID uuid.UUID) ([]*domain.PanelComponent, error) {
	query, args, err := sq.
		Select(
			"panel_id",
			"test_id",
			"sequence_no",
		).From("panel_components").
		Where(sq.Eq{"panel_id": panelID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := lr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying panel components: %w", err)
	}
	defer rows.Close()

	var components []*domain.PanelComponent

	for rows.Next() {
		var component domain.PanelComponent
		err := rows.Scan(
			&component.PanelID,
			&component.TestCatalogID,
			&component.SequenceNo,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning panel component row: %w", err)
		}
		components = append(components, &component)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating panel component rows: %w", err)
	}

	return components, nil
}

func (lr *LabTestsRepository) GetTestParametersByTestCatalogID(ctx context.Context, testCatalogID uuid.UUID) ([]*domain.TestParameter, error) {
	query, args, err := sq.
		Select(
			"id",
			"test_id",
			"name",
			"unit",
			"result_type",
			"sequence_no",
		).From("test_parameters").
		Where(sq.Eq{"test_id": testCatalogID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := lr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying test parameters: %w", err)
	}
	defer rows.Close()

	var parameters []*domain.TestParameter

	for rows.Next() {
		var parameter domain.TestParameter
		err := rows.Scan(
			&parameter.ID,
			&parameter.TestCatalogID,
			&parameter.Name,
			&parameter.Unit,
			&parameter.ResultType,
			&parameter.SequenceNo,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning test parameter row: %w", err)
		}
		parameters = append(parameters, &parameter)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating test parameter rows: %w", err)
	}

	return parameters, nil
}

func (lr *LabTestsRepository) GetReferenceRangesByTestParameterID(ctx context.Context, testParameterID uuid.UUID) ([]*domain.ReferenceRange, error) {
	query, args, err := sq.
		Select(
			"id",
			"parameter_id",
			"gender",
			"age_min_years",
			"age_max_years",
			"low_value",
			"high_value",
			"text_range",
		).From("reference_ranges").
		Where(sq.Eq{"parameter_id": testParameterID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := lr.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying reference ranges: %w", err)
	}
	defer rows.Close()

	var ranges []*domain.ReferenceRange

	for rows.Next() {
		var rangeData domain.ReferenceRange
		err := rows.Scan(
			&rangeData.ID,
			&rangeData.TestParameterID,
			&rangeData.Gender,
			&rangeData.MinAge,
			&rangeData.MaxAge,
			&rangeData.MinValue,
			&rangeData.MaxValue,
			&rangeData.TextValue,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning reference range row: %w", err)
		}
		ranges = append(ranges, &rangeData)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating reference range rows: %w", err)
	}

	return ranges, nil
}

func (lr *LabTestsRepository) UpdateDepartment(ctx context.Context, dept *domain.Department) error {

	builder := sq.
		Update("departments").
		Where(sq.Eq{"id": dept.ID}).
		PlaceholderFormat(sq.Dollar)

	if dept.Name != "" {
		builder = builder.Set("name", dept.Name)
	}
	if dept.Description != nil {
		builder = builder.Set("description", dept.Description)
	}

	query, args, err := builder.ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}
	fmt.Println(query)
	fmt.Println(args...)

	_, err = lr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}

func (lr *LabTestsRepository) UpdatePanel(ctx context.Context, panel *domain.Panels) error {
	now := time.Now()

	builder := sq.
		Update("panels").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": panel.ID})

	if panel.Name != "" {
		builder = builder.Set("name", panel.Name)
	}
	if panel.Code != nil {
		builder = builder.Set("code", panel.Code)
	}
	if panel.PanelPrice != nil {
		builder = builder.Set("panel_price", panel.PanelPrice)
	}

	builder = builder.Set("updated_at", now)

	query, args, err := builder.ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = lr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update panel: %w", err)
	}

	return nil
}

func (lr *LabTestsRepository) UpdateTestCatalog(ctx context.Context, test *domain.TestCatalog) error {
	now := time.Now()

	builder := sq.
		Update("test_catalog").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": test.ID})

	if test.Name != "" {
		builder = builder.Set("name", test.Name)
	}
	if test.Code != nil {
		builder = builder.Set("code", test.Code)
	}
	if test.SampleType != nil {
		builder = builder.Set("sample_type", test.SampleType)
	}
	if test.TestPrice != nil {
		builder = builder.Set("price", test.TestPrice)
	}
	if test.TurnAroundTime != nil {
		builder = builder.Set("turnaround_hours", test.TurnAroundTime)
	}

	builder = builder.Set("updated_at", now)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = lr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update test catalog: %w", err)
	}

	return nil
}

func (lr *LabTestsRepository) UpdateTestParameter(ctx context.Context, parameter *domain.TestParameter) error {
	now := time.Now()

	builder := sq.
		Update("test_parameters").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": parameter.ID})

	if parameter.Name != "" {
		builder = builder.Set("name", parameter.Name)
	}
	if parameter.Unit != "" {
		builder = builder.Set("unit", parameter.Unit)
	}
	if parameter.ResultType != "" {
		builder = builder.Set("result_type", parameter.ResultType)
	}
	if parameter.SequenceNo != 0 {
		builder = builder.Set("sequence_no", parameter.SequenceNo)
	}

	builder = builder.Set("updated_at", now)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = lr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update test parameter: %w", err)
	}

	return nil
}

func (lr *LabTestsRepository) UpdateReferenceRange(ctx context.Context, rangeData *domain.ReferenceRange) error {
	now := time.Now()

	builder := sq.
		Update("reference_ranges").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": rangeData.ID})

	if rangeData.Gender != "" {
		builder = builder.Set("gender", rangeData.Gender)
	}
	if rangeData.MinAge != nil {
		builder = builder.Set("age_min_years", rangeData.MinAge)
	}
	if rangeData.MaxAge != nil {
		builder = builder.Set("age_max_years", rangeData.MaxAge)
	}
	if rangeData.MinValue != nil {
		builder = builder.Set("low_value", rangeData.MinValue)
	}
	if rangeData.MaxValue != nil {
		builder = builder.Set("high_value", rangeData.MaxValue)
	}
	if rangeData.TextValue != nil {
		builder = builder.Set("text_range", rangeData.TextValue)
	}

	builder = builder.Set("updated_at", now)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = lr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update reference range: %w", err)
	}

	return nil
}
