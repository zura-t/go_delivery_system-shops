package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func (uc *ShopUseCase) DeleteShop(id int64, user_id int64) (string, int, error) {
	err := uc.store.DeleteShop(context.Background(), db.DeleteShopParams{
		ID: id,
		UserID: user_id,
	})
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