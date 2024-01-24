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

func randomMenuItem() *db.MenuItem {
	return &db.MenuItem{
		ID:   pkg.RandomInt(1, 1000),
		Name: pkg.RandomString(6),
		Description: sql.NullString{
			Valid:  true,
			String: pkg.RandomString(10),
		},
		Price: int32(pkg.RandomInt(10, 50)),
	}
}

func randomMenuItemReq() *entity.CreateMenuItem {
	return &entity.CreateMenuItem{
		Name:        pkg.RandomString(6),
		Description: pkg.RandomString(10),
		Price:       int32(pkg.RandomInt(10, 50)),
	}
}

func randomMenuItemDbReq() *db.CreateMenuItemParams {
	return &db.CreateMenuItemParams{
		Name: pkg.RandomString(6),
		Description: sql.NullString{
			Valid:  true,
			String: pkg.RandomString(10),
		},
		Price: int32(pkg.RandomInt(10, 50)),
	}
}

func Test_create_menu(t *testing.T) {
	// shop := randomShop()
	n := 5
	req := make([]*entity.CreateMenuItem, n)
	for i := 0; i < n; i++ {
		req[i] = randomMenuItemReq()
	}

	reqToDb := make([]db.CreateMenuItemParams, n)
	for i := 0; i < n; i++ {
		reqToDb[i] = *randomMenuItemDbReq()
	}

	dbRes := make([]db.MenuItem, n)
	for i := 0; i < n; i++ {
		dbRes[i] = *randomMenuItem()
	}

	tests := []struct {
		name          string
		req           []*entity.CreateMenuItem
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res []*entity.MenuItem, st int, err error)
	}{
		{
			name: "OK",
			req:  req,
			buildStub: func(store *mockdb.MockStore) {
					store.EXPECT().CreateMenuItem(gomock.Any(), reqToDb[0]).
					Times(n)
				
				// Return(menuItemCreated, nil)
			},
			checkResponse: func(t *testing.T, res []*entity.MenuItem, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "InternalError",
			req:  req,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().CreateMenuItem(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.MenuItem{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res []*entity.MenuItem, st int, err error) {
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
			res, st, err := u.CreateMenu(tc.req)

			tc.checkResponse(t, res, st, err)
		})
	}
}
