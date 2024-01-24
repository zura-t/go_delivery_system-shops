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

func Test_get_shops(t *testing.T) {
	n := 5
	shops := make([]db.Shop, n)
	for i := 0; i < n; i++ {
		shops[i] = *randomShop()
	}

	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res []*entity.Shop, st int, err error)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListShopsParams{}

				store.EXPECT().
					ListShops(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(shops, nil)
			},
			checkResponse: func(t *testing.T, res []*entity.Shop, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "InternalError",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListShops(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Shop{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res []*entity.Shop, st int, err error) {
				require.Error(t, err)
				require.Equal(t, http.StatusInternalServerError, st)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			u := newTestServer(t, store)

			tc.buildStubs(store)
			res, st, err := u.GetShops()

			tc.checkResponse(t, res, st, err)
		})
	}
}
