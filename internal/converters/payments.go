package converters

import (
	"static/internal/models/domain"
	"static/internal/models/dto"
)

type PaymentsConverter struct{}

func NewPaymentsConverter() *PaymentsConverter {
	return &PaymentsConverter{}
}

func (c *PaymentsConverter) ToPaymentDTO(item domain.Payment) dto.Payment {
	return dto.Payment{
		ID:   item.ID,
		Name: item.Name,
	}
}

func (c *PaymentsConverter) ToPaymentDTOs(items []domain.Payment) []dto.Payment {
	res := make([]dto.Payment, len(items))
	for i, item := range items {
		res[i] = c.ToPaymentDTO(item)
	}

	return res
}
