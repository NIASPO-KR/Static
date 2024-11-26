package static

import (
	"context"
	"fmt"

	"static/internal/converters"
	"static/internal/models/dto"
	"static/internal/repository/static"
)

type PickupPointsUseCase interface {
	GetPickupPoints(ctx context.Context) ([]dto.PickupPoint, error)
}

type pickupPointUseCase struct {
	pickupPointsRepository static.PickupPointsRepository
	pickupPointsConverter  *converters.PickupPointsConverter
}

func NewPickupPointUseCase(
	pickupPointsRepository static.PickupPointsRepository,
) PickupPointsUseCase {
	return &pickupPointUseCase{
		pickupPointsRepository: pickupPointsRepository,
		pickupPointsConverter:  converters.NewPickupPointsConverter(),
	}
}

func (ppuc *pickupPointUseCase) GetPickupPoints(ctx context.Context) ([]dto.PickupPoint, error) {
	pickupPoints, err := ppuc.pickupPointsRepository.GetPickupPoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get pickup points: %w", err)
	}

	return ppuc.pickupPointsConverter.ToPickupPointDTOs(pickupPoints), nil
}
