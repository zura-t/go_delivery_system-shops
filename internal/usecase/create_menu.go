package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func (uc *ShopUseCase) CreateMenu(req *entity.CreateMenuItem) ([]*entity.GetMenuItem, int, error) {
	shop, err := uc.store.GetShop(context.Background(), req.ShopId)
	if shop.UserID != req.UserId {
		err = fmt.Errorf("you are not an admin: %s", err)
		return nil, http.StatusBadRequest, err
	}

	menu := make([]*entity.GetMenuItem, len(req.MenuItems))

	for i := 0; i < len(req.MenuItems); i++ {
		arg := db.CreateMenuItemParams{
			Name: req.MenuItems[i].Name,
			Description: sql.NullString{
				String: req.MenuItems[i].Description,
				Valid:  true,
			},
			Photo: sql.NullString{
				String: req.MenuItems[i].Photo,
				Valid:  true,
			},
			Price:  req.MenuItems[i].Price,
			ShopID: req.ShopId,
		}

		menuItem, err := uc.store.CreateMenuItem(context.Background(), arg)
		if err != nil {
			err = fmt.Errorf("failed to create menu: %s", err)
			return nil, http.StatusInternalServerError, err
		}

		menu[i] = convertMenu(menuItem)
	}

	return menu, http.StatusOK, nil
}
