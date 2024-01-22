package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zura-t/go_delivery_system-shops/internal/entity"
)

func (uc *ShopUseCase) GetShopInfo(id int64) (*entity.Shop, int, error) {
	shopCreated, err := uc.store.GetShop(context.Background(), id)
	if err != nil {
		err = fmt.Errorf("failed to get shop: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := convertShop(shopCreated)

	return res, http.StatusOK, nil
}