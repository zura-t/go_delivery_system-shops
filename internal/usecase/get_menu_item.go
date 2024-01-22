package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zura-t/go_delivery_system-shops/internal/entity"
)

func (uc *ShopUseCase) GetMenuItem(id int64) (*entity.MenuItem, int, error) {
	i, err := uc.store.GetMenuItem(context.Background(), id)
	if err != nil {
		err := fmt.Errorf("failed to get menu item: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := convertMenu(i)
	return res, http.StatusOK, nil
}
