package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func (uc *ShopUseCase) CreateMenu(req []*entity.CreateMenuItem) ([]*entity.MenuItem, int, error) {
	menu := make([]*entity.MenuItem, len(req))

	for i := 0; i < len(req); i++ {
		arg := db.CreateMenuItemParams{
			Name: req[i].Name,
			Description: sql.NullString{
				String: req[i].Description,
				Valid:  true,
			},
			Photo: sql.NullString{
				String: req[i].Photo,
				Valid:  true,
			},
			Price:  req[i].Price,
			ShopID: req[i].ShopID,
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
