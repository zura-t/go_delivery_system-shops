package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

func (uc *ShopUseCase) DeleteMenuItem(id int64) (string, int, error) {
	err := uc.store.DeleteMenuItem(context.Background(), id)
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