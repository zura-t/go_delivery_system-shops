package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
)

func (uc *ShopUseCase) DeleteShop(id int64) (string, int, error) {
	err := uc.store.DeleteShop(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("shop not found: %s", err)
			return "", http.StatusNotFound, err
		}
		err = fmt.Errorf("failed to delete shop: %s", err)
		return "", http.StatusInternalServerError, err
	}

	return "shop has been deleted", http.StatusOK, nil
}