package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jenish-brainztechs/go-backend/internal/adapter/storage/postgres"
	"github.com/jenish-brainztechs/go-backend/internal/core/domain"
)

type VisitRepository struct {
	DB *postgres.DB
}

func NewVisitRepository(db *postgres.DB) *VisitRepository {
	return &VisitRepository{
		DB: db,
	}
}

func (vr *VisitRepository) CreateVisit(ctx context.Context, visit *domain.Visit) (*domain.Visit, error) {
	query, args, err := sq.
		Insert("visits").
		Columns(
			"visit_no",
			"patient_id",
			"doctor_id",
			"registered_by",
			"visit_date",
			"v_status",
		).
		Values(
			visit.VisitNo,
			visit.PatientID,
			visit.DoctorID,
			visit.RegisteredBy,
			visit.VisitDate,
			visit.Status,
		).
		Suffix(`
		RETURNING
		id
	`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("DoctorRepo.CreateDoctor build: %w", err)
	}

	err = vr.DB.QueryRow(ctx, query, args...).Scan(
		&visit.ID,
	)

	if err != nil {
		return nil, err
	}

	return visit, nil
}

func (vr *VisitRepository) GetVisitByID(ctx context.Context, id uuid.UUID) (*domain.ListVisit, error) {
	var (
		doctorFirstName, doctorLastName, registerFirstName, registerLastName sql.NullString
		status                                                               string
	)

	query := `
		SELECT
			V.ID,
			V.VISIT_NO,
			P.FIRST_NAME,
			P.LAST_NAME,
			D.FIRST_NAME,
			D.LAST_NAME,
			PR.FIRST_NAME,
			PR.LAST_NAME,
			V.VISIT_DATE,
			V.V_STATUS
		FROM VISITS V
		LEFT JOIN PATIENTS P ON P.ID = V.PATIENT_ID
		LEFT JOIN DOCTORS D ON D.ID = V.DOCTOR_ID
		LEFT JOIN USERS U ON U.ID = V.REGISTERED_BY
		LEFT JOIN PROFILE PR ON PR.USER_ID = U.ID
		WHERE V.ID = $1;
	`

	var v domain.ListVisit

	err := vr.DB.QueryRow(ctx, query, id).Scan(
		&v.ID,
		&v.VisitNo,
		&v.PatientFirstName,
		&v.PatientLastName,
		&doctorFirstName,
		&doctorLastName,
		&registerFirstName,
		&registerLastName,
		&v.VisitDate,
		&status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("visit not found")
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if doctorFirstName.Valid {
		v.DoctorFirstName = doctorFirstName.String
	}
	if doctorLastName.Valid {
		v.DoctorLastName = doctorLastName.String
	}
	if registerFirstName.Valid {
		v.RegisterFirstName = registerFirstName.String
	}
	if registerLastName.Valid {
		v.RegisterLastName = registerLastName.String
	}

	v.Status = domain.Status(status)

	return &v, nil
}

func (vr *VisitRepository) GetVisitByPatientID(ctx context.Context, id uuid.UUID) ([]*domain.ListVisit, error) {
	var doctorFirstName, doctorLastName, registerFirstName, registerLastName sql.NullString
	var status string
	query := `
		SELECT
			V.ID,
			V.VISIT_NO,
			P.FIRST_NAME,
			P.LAST_NAME,
			D.FIRST_NAME,
			D.LAST_NAME,
			PR.FIRST_NAME,
			PR.LAST_NAME,
			V.VISIT_DATE,
			V.V_STATUS
		FROM
			VISITS V
			LEFT JOIN PATIENTS P ON P.ID = V.PATIENT_ID
			LEFT JOIN DOCTORS D ON D.ID = V.DOCTOR_ID
			LEFT JOIN USERS U ON U.ID = V.REGISTERED_BY
			LEFT JOIN PROFILE PR ON PR.USER_ID = U.ID
		WHERE
			V.PATIENT_ID = $1;
	`

	rows, err := vr.DB.Query(ctx, query, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visits []*domain.ListVisit

	for rows.Next() {
		var v domain.ListVisit

		err := rows.Scan(
			&v.ID,
			&v.VisitNo,
			&v.PatientFirstName,
			&v.PatientLastName,
			&doctorFirstName,
			&doctorLastName,
			&registerFirstName,
			&registerLastName,
			&v.VisitDate,
			&status,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if doctorFirstName.Valid {
			v.DoctorFirstName = doctorFirstName.String
		}

		if doctorLastName.Valid {
			v.DoctorLastName = doctorLastName.String
		}

		if registerFirstName.Valid {
			v.RegisterFirstName = registerFirstName.String
		}

		if registerLastName.Valid {
			v.RegisterLastName = registerLastName.String
		}
		fmt.Println("Status from DB:", status)
		v.Status = domain.Status(status)

		visits = append(visits, &v)
	}

	return visits, nil
}

func (vr *VisitRepository) UpdateVisitByID(ctx context.Context, v *domain.Visit) error {
	visit_no := nullString(v.VisitNo)
	patient_id := nullUUID(v.PatientID)
	doctor_id := nullUUID(v.DoctorID)
	v_status := nullString(string(v.Status))

	query, args, err := sq.
		Update("visits").
		Set("visit_no", sq.Expr("COALESCE(?, visit_no)", visit_no)).
		Set("patient_id", sq.Expr("COALESCE(?, patient_id)", patient_id)).
		Set("doctor_id", sq.Expr("COALESCE(?, doctor_id)", doctor_id)).
		Set("v_status", sq.Expr("COALESCE(?, v_status)", v_status)).
		Set("updated_at", sq.Expr("NOW()")).
		Where(sq.Eq{"id": v.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = vr.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update visit: %w", err)
	}

	return nil
}
