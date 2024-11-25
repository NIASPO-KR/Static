package static

import (
	"context"
	"fmt"

	"static/internal/infrastructure/database/postgres"
	"static/internal/models/domain"
)

const (
	pickupPointsDB = "pickup_points"
)

type PickupPointsRepository interface {
	GetPickupPoints(ctx context.Context) ([]domain.PickupPoint, error)
}

type pickupPointsRepository struct {
	db *postgres.Postgres
}

func NewPickupPointsRepository(
	db *postgres.Postgres,
) PickupPointsRepository {
	return &pickupPointsRepository{
		db: db,
	}
}

func (pp *pickupPointsRepository) GetPickupPoints(ctx context.Context) ([]domain.PickupPoint, error) {
	qb := pp.db.Builder.Select(
		"id",
		"address",
	).From(pickupPointsDB)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql %w", err)
	}

	rows, err := pp.db.SqlDB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query context %w", err)
	}
	defer rows.Close()

	var pickupPoints []domain.PickupPoint
	for rows.Next() {
		var pickupPoint domain.PickupPoint
		if err := rows.Scan(&pickupPoint.ID, &pickupPoint.Address); err != nil {
			return nil, fmt.Errorf("row scan %w", err)
		}

		pickupPoints = append(pickupPoints, pickupPoint)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows %w", err)
	}

	return pickupPoints, nil
}
