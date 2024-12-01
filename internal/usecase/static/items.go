package static

import (
	"context"
	"fmt"

	"static/internal/converters"
	"static/internal/models/dto"
	"static/internal/ports/repository"
	"static/internal/usecase"
)

type itemsUseCase struct {
	itemsRepository repository.ItemsRepository
	itemsConverter  *converters.ItemsConverter
}

func NewItemsUseCase(
	itemsRepository repository.ItemsRepository,
) usecase.ItemsUseCase {
	return &itemsUseCase{
		itemsRepository: itemsRepository,
		itemsConverter:  converters.NewItemsConverter(),
	}
}

func (iuc *itemsUseCase) GetItems(ctx context.Context) ([]dto.Item, error) {
	items, err := iuc.itemsRepository.GetItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo get items: %w", err)
	}

	return iuc.itemsConverter.ToItemDTOs(items), nil
}
