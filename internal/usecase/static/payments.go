package static

import (
	"context"
	"fmt"

	"static/internal/converters"
	"static/internal/models/dto"
	"static/internal/ports/repository"
	"static/internal/usecase"
)

type paymentsUseCase struct {
	paymentsRepository repository.PaymentsRepository
	paymentsConverter  *converters.PaymentsConverter
}

func NewPaymentsUseCase(
	paymentsRepository repository.PaymentsRepository,
) usecase.PaymentsUseCase {
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
