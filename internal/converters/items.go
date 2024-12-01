package converters

import (
	"static/internal/models/dto"
	"static/internal/models/entities"
)

type ItemsConverter struct{}

func NewItemsConverter() *ItemsConverter {
	return &ItemsConverter{}
}

func (c *ItemsConverter) ToItemDTO(item entities.Item) dto.Item {
	return dto.Item{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
}

func (c *ItemsConverter) ToItemDTOs(items []entities.Item) []dto.Item {
	res := make([]dto.Item, len(items))
	for i, item := range items {
		res[i] = c.ToItemDTO(item)
	}

	return res
}
