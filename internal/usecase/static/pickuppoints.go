package static

import (
	"context"
	"fmt"

	"static/internal/converters"
	"static/internal/models/dto"
	"static/internal/ports/repository"
	"static/internal/usecase"
)

type pickupPointUseCase struct {
	pickupPointsRepository repository.PickupPointsRepository
	pickupPointsConverter  *converters.PickupPointsConverter
}

func NewPickupPointUseCase(
	pickupPointsRepository repository.PickupPointsRepository,
) usecase.PickupPointsUseCase {
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
