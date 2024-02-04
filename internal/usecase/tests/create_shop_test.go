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

func randomShop() *db.Shop {
	return &db.Shop{
		ID:       pkg.RandomInt(1, 1000),
		Name:     pkg.RandomString(6),
		IsClosed: false,
	}
}

func Test_create_shop(t *testing.T) {
	shop := randomShop()

	tests := []struct {
		name          string
		req           *entity.CreateShop
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *entity.Shop, st int, err error)
	}{
		{
			name: "OK",
			req: &entity.CreateShop{
				Name:        shop.Name,
				Description: shop.Description.String,
				OpenTime:    shop.OpenTime.Time,
				CloseTime:   shop.CloseTime.Time,
				IsClosed:    shop.IsClosed,
				UserId:      pkg.RandomInt(1, 1000),
			},
			buildStub: func(store *mockdb.MockStore) {
				req := db.CreateShopParams{
					Name:        shop.Name,
					Description: shop.Description,
					OpenTime:    shop.OpenTime,
					CloseTime:   shop.CloseTime,
					IsClosed:    shop.IsClosed,
					UserID:      pkg.RandomInt(1, 1000),
				}
				shopCreated := db.Shop{
					ID:          shop.ID,
					Name:        shop.Name,
					Description: shop.Description,
					OpenTime:    shop.OpenTime,
					CloseTime:   shop.CloseTime,
					IsClosed:    shop.IsClosed,
				}

				store.EXPECT().CreateShop(gomock.Any(), req).
					Times(1).
					Return(shopCreated, nil)
			},
			checkResponse: func(t *testing.T, res *entity.Shop, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				createdShop := res
				require.Equal(t, shop.Name, createdShop.Name)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "InternalError",
			req: &entity.CreateShop{
				Name:        shop.Name,
				Description: shop.Description.String,
				OpenTime:    shop.OpenTime.Time,
				CloseTime:   shop.CloseTime.Time,
				IsClosed:    shop.IsClosed,
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateShop(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Shop{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *entity.Shop, st int, err error) {
				require.Error(t, err)
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
			res, st, err := u.CreateShop(tc.req)

			tc.checkResponse(t, res, st, err)
		})
	}
}
