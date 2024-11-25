package static

import (
	"context"
	"fmt"

	"static/internal/infrastructure/database/postgres"
	"static/internal/models/domain"
)

const (
	paymentsDB = "payments"
)

type PaymentsRepository interface {
	GetPayments(ctx context.Context) ([]domain.Payment, error)
}

type paymentsRepository struct {
	db *postgres.Postgres
}

func NewPaymentsRepository(
	db *postgres.Postgres,
) PaymentsRepository {
	return &paymentsRepository{
		db: db,
	}
}

func (p *paymentsRepository) GetPayments(ctx context.Context) ([]domain.Payment, error) {
	qb := p.db.Builder.Select(
		"id",
		"name",
	).From(paymentsDB)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql %w", err)
	}

	rows, err := p.db.SqlDB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query context %w", err)
	}
	defer rows.Close()

	var payments []domain.Payment
	for rows.Next() {
		var payment domain.Payment
		if err := rows.Scan(&payment.ID, &payment.Name); err != nil {
			return nil, fmt.Errorf("row scan %w", err)
		}

		payments = append(payments, payment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows %w", err)
	}

	return payments, nil
}
