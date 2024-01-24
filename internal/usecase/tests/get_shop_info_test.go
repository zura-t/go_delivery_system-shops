package usecase_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	mockdb "github.com/zura-t/go_delivery_system-shops/pkg/db/mock"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func Test_get_shop(t *testing.T) {
	shop := randomShop()

	tests := []struct {
		name          string
		req           int64
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *entity.Shop, st int, err error)
	}{
		{
			name: "OK",
			req:  shop.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetShop(gomock.Any(), shop.ID).
					Times(1).
					Return(*shop, nil)
			},
			checkResponse: func(t *testing.T, res *entity.Shop, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, shop.Name, res.Name)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "NotFound",
			req:  shop.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetShop(gomock.Any(), shop.ID).Times(1).Return(db.Shop{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, res *entity.Shop, st int, err error) {
				require.Error(t, err)
				require.Empty(t, res)
				require.Equal(t, http.StatusNotFound, st)
			},
		},
		{
			name: "InternalError",
			req:  shop.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetShop(gomock.Any(), shop.ID).
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
			res, st, err := u.GetShopInfo(tc.req)

			tc.checkResponse(t, res, st, err)
		})
	}
}
