package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/zura-t/go_delivery_system-shops/internal/entity"
)

func (uc *ShopUseCase) GetShopInfo(id int64) (*entity.Shop, int, error) {
	shopCreated, err := uc.store.GetShop(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("shop not found: %s", err)
			return nil, http.StatusNotFound, err
		}
		err = fmt.Errorf("failed to get shop: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := convertShop(shopCreated)

	return res, http.StatusOK, nil
}