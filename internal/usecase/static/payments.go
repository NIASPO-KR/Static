package static

import (
	"context"
	"fmt"

	"static/internal/converters"
	"static/internal/models/dto"
	"static/internal/repository/static"
)

type PaymentsUseCase interface {
	GetPayments(ctx context.Context) ([]dto.Payment, error)
}

type paymentsUseCase struct {
	paymentsRepository static.PaymentsRepository
	paymentsConverter  *converters.PaymentsConverter
}

func NewPaymentsUseCase(
	paymentsRepository static.PaymentsRepository,
) PaymentsUseCase {
	return &paymentsUseCase{
		paymentsRepository: paymentsRepository,
		paymentsConverter:  converters.NewPaymentsConverter(),
	}
}

func (puc *paymentsUseCase) GetPayments(ctx context.Context) ([]dto.Payment, error) {
	payments, err := puc.paymentsRepository.GetPayments(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get payments: %w", err)
	}

	return puc.paymentsConverter.ToPaymentDTOs(payments), nil
}
