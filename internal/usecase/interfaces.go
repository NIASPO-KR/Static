package usecase

import (
	"context"

	"static/internal/models/dto"
)

type ItemsUseCase interface {
	GetItems(ctx context.Context) ([]dto.Item, error)
}

type PaymentsUseCase interface {
	GetPayments(ctx context.Context) ([]dto.Payment, error)
}

type PickupPointsUseCase interface {
	GetPickupPoints(ctx context.Context) ([]dto.PickupPoint, error)
}
