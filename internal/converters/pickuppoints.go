package converters

import (
	"static/internal/models/dto"
	"static/internal/models/entities"
)

type PickupPointsConverter struct{}

func NewPickupPointsConverter() *PickupPointsConverter {
	return &PickupPointsConverter{}
}

func (c *PickupPointsConverter) ToPickupPointDTO(item entities.PickupPoint) dto.PickupPoint {
	return dto.PickupPoint{
		ID:      item.ID,
		Address: item.Address,
	}
}

func (c *PickupPointsConverter) ToPickupPointDTOs(items []entities.PickupPoint) []dto.PickupPoint {
	res := make([]dto.PickupPoint, len(items))
	for i, item := range items {
		res[i] = c.ToPickupPointDTO(item)
	}

	return res
}
