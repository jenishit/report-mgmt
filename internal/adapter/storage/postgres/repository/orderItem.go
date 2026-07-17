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

type OrderRepository struct {
	DB *postgres.DB
}

func NewOrderRepository(db *postgres.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (or *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	query, args, err := sq.
		Insert("order_item").
		Columns(
			"visit_id",
			"test_id",
			"panel_id",
			"status",
			"price",
			"collected_by",
		).
		Values(
			order.VisitID,
			order.TestID,
			nullUUID(order.PanelID),
			order.Status,
			order.Price,
			order.CollectedBy,
		).
		Suffix(`
			RETURNING
			id,
			collected_at
		`).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("OrderRepo.CreateOrder build: %w", err)
	}

	err = or.DB.QueryRow(ctx, query, args...).Scan(
		&order.ID,
		&order.CollectedAt,
	)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (or *OrderRepository) GetOrderByID(ctx context.Context, id uuid.UUID) (*domain.Order, error) {
	var panelID uuid.NullUUID
	var price sql.NullFloat64

	query, args, err := sq.
		Select("id", "visit_id", "test_id", "panel_id", "status", "price", "collected_by", "collected_at").
		From("order_item").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("OrderRepo.GetOrderByID build: %w", err)
	}

	var order domain.Order
	err = or.DB.QueryRow(ctx, query, args...).Scan(
		&order.ID,
		&order.VisitID,
		&order.TestID,
		&panelID,
		&order.Status,
		&price,
		&order.CollectedBy,
		&order.CollectedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order not found")
		}
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	if panelID.Valid {
		order.PanelID = panelID.UUID
	}
	if price.Valid {
		order.Price = price.Float64
	}

	return &order, nil
}

func (or *OrderRepository) GetOrdersByVisitID(ctx context.Context, visitID uuid.UUID) ([]*domain.Order, error) {
	var panelID uuid.NullUUID
	var price sql.NullFloat64

	query, args, err := sq.
		Select("id", "visit_id", "test_id", "panel_id", "status", "price", "collected_by", "collected_at").
		From("order_item").
		Where(sq.Eq{"visit_id": visitID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("OrderRepo.GetOrdersByVisitID build: %w", err)
	}

	rows, err := or.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(
			&order.ID,
			&order.VisitID,
			&order.TestID,
			&panelID,
			&order.Status,
			&price,
			&order.CollectedBy,
			&order.CollectedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		if panelID.Valid {
			order.PanelID = panelID.UUID
		}
		if price.Valid {
			order.Price = price.Float64
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

func (or *OrderRepository) UpdateOrder(ctx context.Context, order *domain.Order) error {
	panelID := nullUUID(order.PanelID)
	status := nullString(string(order.Status))
	price := sql.NullFloat64{}
	if order.Price != 0 {
		price = sql.NullFloat64{Float64: order.Price, Valid: true}
	}
	collectedBy := nullUUID(order.CollectedBy)

	query, args, err := sq.
		Update("order_item").
		Set("panel_id", sq.Expr("COALESCE(?, panel_id)", panelID)).
		Set("status", sq.Expr("COALESCE(?, status)", status)).
		Set("price", sq.Expr("COALESCE(?, price)", price)).
		Set("collected_by", sq.Expr("COALESCE(?, collected_by)", collectedBy)).
		Set("updated_at", sq.Expr("NOW()")).
		Where(sq.Eq{"id": order.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = or.DB.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}

	return nil
}
