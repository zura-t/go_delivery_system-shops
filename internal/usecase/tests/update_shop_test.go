package usecase_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	"github.com/zura-t/go_delivery_system-shops/pkg"
	mockdb "github.com/zura-t/go_delivery_system-shops/pkg/db/mock"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

type UpdateShopRequest struct {
	Id   int64
	Data *entity.UpdateShopInfo
}

func Test_update_shop(t *testing.T) {
	shop := randomShop()
	newName := pkg.RandomString(6)

	tests := []struct {
		name          string
		req           *UpdateShopRequest
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *entity.Shop, st int, err error)
	}{
		{
			name: "OK",
			req: &UpdateShopRequest{
				Id: shop.ID,
				Data: &entity.UpdateShopInfo{
					Name:        newName,
					Description: shop.Description.String,
					OpenTime:    shop.OpenTime.Time,
					CloseTime:   shop.CloseTime.Time,
					IsClosed:    shop.IsClosed,
				},
			},
			buildStub: func(store *mockdb.MockStore) {
				req := db.UpdateShopParams{
					ID: shop.ID,
					Name: sql.NullString{
						Valid:  true,
						String: newName,
					},
					IsClosed: sql.NullBool{
						Valid: true,
						Bool:  shop.IsClosed,
					},
				}
				shopUpdated := db.Shop{
					ID:       shop.ID,
					Name:     newName,
					IsClosed: shop.IsClosed,
				}

				store.EXPECT().UpdateShop(gomock.Any(), req).
					Times(1).
					Return(shopUpdated, nil)
			},
			checkResponse: func(t *testing.T, res *entity.Shop, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				updatedShop := res
				require.Equal(t, shop.ID, updatedShop.ID)
				require.Equal(t, newName, updatedShop.Name)
				require.Equal(t, shop.IsClosed, updatedShop.IsClosed)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "NotFound",
			req: &UpdateShopRequest{
				Id: shop.ID,
				Data: &entity.UpdateShopInfo{
					Name: newName,
				},
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateShop(gomock.Any(), gomock.Any()).Times(1).Return(db.Shop{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, res *entity.Shop, st int, err error) {
				require.Error(t, err)
				require.Empty(t, res)
				require.Equal(t, http.StatusNotFound, st)
			},
		},
		{
			name: "InternalError",
			req: &UpdateShopRequest{
				Id: shop.ID,
				Data: &entity.UpdateShopInfo{
					Name: newName,
				},
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateShop(gomock.Any(), gomock.Any()).Times(1).Return(db.Shop{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *entity.Shop, st int, err error) {
				require.Error(t, err)
				require.Empty(t, res)
				require.Equal(t, http.StatusInternalServerError, st)
			},
		},
	}

	for i := range tests {
		tc := tests[i]

		t.Run(tc.name, func(t *testing.T) {
			storeCtrl := gomock.NewController(t)
			defer storeCtrl.Finish()
			store := mockdb.NewMockStore(storeCtrl)

			u := newTestServer(t, store)

			tc.buildStub(store)
			res, st, err := u.UpdateShop(tc.req.Id, tc.req.Data)

			tc.checkResponse(t, res, st, err)
		})
	}
}
