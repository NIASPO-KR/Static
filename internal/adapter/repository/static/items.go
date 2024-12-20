package static

import (
	"context"
	"fmt"

	"static/internal/infrastructure/database/postgres"
	"static/internal/models/entities"
	"static/internal/ports/repository"
)

const (
	itemsDB = "items"
)

type itemsRepository struct {
	db *postgres.Postgres
}

func NewItemsRepository(
	db *postgres.Postgres,
) repository.ItemsRepository {
	return &itemsRepository{
		db: db,
	}
}

func (i *itemsRepository) GetItems(ctx context.Context) ([]entities.Item, error) {
	qb := i.db.Builder.Select(
		"id",
		"name",
		"price",
	).From(itemsDB)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql %w", err)
	}

	rows, err := i.db.SqlDB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query context %w", err)
	}
	defer rows.Close()

	var items []entities.Item
	for rows.Next() {
		var item entities.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price); err != nil {
			return nil, fmt.Errorf("row scan %w", err)
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows %w", err)
	}

	return items, nil
}
