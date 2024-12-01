package converters

import (
	"static/internal/models/dto"
	"static/internal/models/entities"
)

type PaymentsConverter struct{}

func NewPaymentsConverter() *PaymentsConverter {
	return &PaymentsConverter{}
}

func (c *PaymentsConverter) ToPaymentDTO(item entities.Payment) dto.Payment {
	return dto.Payment{
		ID:   item.ID,
		Name: item.Name,
	}
}

func (c *PaymentsConverter) ToPaymentDTOs(items []entities.Payment) []dto.Payment {
	res := make([]dto.Payment, len(items))
	for i, item := range items {
		res[i] = c.ToPaymentDTO(item)
	}

	return res
}
