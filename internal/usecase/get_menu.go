package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func convertMenu(menuItem db.MenuItem) *entity.GetMenuItem {
	return &entity.GetMenuItem{
		ID:          menuItem.ID,
		Name:        menuItem.Name,
		Description: menuItem.Description.String,
		Photo:       menuItem.Photo.String,
		Price:       menuItem.Price,
		ShopId:      menuItem.ShopID,
		CreatedAt:   menuItem.CreatedAt,
	}
}

func (uc *ShopUseCase) GetMenu(shopId int64) ([]*entity.GetMenuItem, int, error) {
	menu, err := uc.store.ListMenuItems(context.Background(), shopId)
	if err != nil {
		err = fmt.Errorf("failed to get menu: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := make([]*entity.GetMenuItem, len(menu))

	for i := 0; i < len(menu); i++ {
		menu := convertMenu(menu[i])
		res[i] = menu
	}

	return res, http.StatusOK, nil
}
