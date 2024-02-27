package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zura-t/go_delivery_system-shops/internal/entity"
)

func (uc *ShopUseCase) GetShopsAdmin(user_id int64) ([]*entity.Shop, int, error) {
	shops, err := uc.store.GetShopsAdmin(context.Background(), user_id)
	if err != nil {
		err = fmt.Errorf("failed to find shop list: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := make([]*entity.Shop, len(shops))

	for i := 0; i < len(shops); i++ {
		shop := convertShop(shops[i])
		res[i] = shop
	}

	return res, http.StatusOK, nil
}