package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func (uc *ShopUseCase) UpdateMenuItem(id int64, req *entity.UpdateMenuItem) (*entity.MenuItem, int, error) {
	arg := db.UpdateMenuItemParams{
		ID: id,
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},
		Description: sql.NullString{
			String: req.Description,
			Valid:  req.Description != "",
		},
		Price: sql.NullInt32{
			Int32: req.Price,
			Valid: req.Price != 0,
		},
	}

	menuItemUpdated, err := uc.store.UpdateMenuItem(context.Background(), arg)
	if err != nil {
		err = fmt.Errorf("failed to update menu item: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := convertMenu(menuItemUpdated)

	return res, http.StatusOK, nil
}
