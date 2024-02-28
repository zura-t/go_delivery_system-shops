package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

func (uc *ShopUseCase) DeleteMenuItem(id int64, user_id int64) (string, int, error) {
	shop, err := uc.store.GetShopWithMenuItemId(context.Background(), id)
	if shop.UserID != user_id {
		err = fmt.Errorf("failed to delete menu item: %s", err)
		return "", http.StatusLocked, err
	}

	err = uc.store.DeleteMenuItem(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("menu item not found: %s", err)
			return "", http.StatusNotFound, err
		}
		err = fmt.Errorf("failed to delete menu item: %s", err)
		return "", http.StatusInternalServerError, err
	}

	return "menu item has been deleted", http.StatusOK, nil
}