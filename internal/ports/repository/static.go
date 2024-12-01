package repository

import (
	"context"

	"static/internal/models/entities"
)

type ItemsRepository interface {
	GetItems(ctx context.Context) ([]entities.Item, error)
}

type PaymentsRepository interface {
	GetPayments(ctx context.Context) ([]entities.Payment, error)
}
type PickupPointsRepository interface {
	GetPickupPoints(ctx context.Context) ([]entities.PickupPoint, error)
}
