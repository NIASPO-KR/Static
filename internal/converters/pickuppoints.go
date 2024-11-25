package converters

import (
	"static/internal/models/domain"
	"static/internal/models/dto"
)

type PickupPointsConverter struct{}

func NewPickupPointsConverter() *PickupPointsConverter {
	return &PickupPointsConverter{}
}

func (c *PickupPointsConverter) ToPickupPointDTO(item domain.PickupPoint) dto.PickupPoint {
	return dto.PickupPoint{
		ID:      item.ID,
		Address: item.Address,
	}
}

func (c *PickupPointsConverter) ToPickupPointDTOs(items []domain.PickupPoint) []dto.PickupPoint {
	res := make([]dto.PickupPoint, len(items))
	for i, item := range items {
		res[i] = c.ToPickupPointDTO(item)
	}

	return res
}
