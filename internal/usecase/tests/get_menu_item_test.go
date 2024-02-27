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

func Test_get_menu_item(t *testing.T) {
	menuItem := randomMenuItem()

	tests := []struct {
		name          string
		req           int64
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *entity.GetMenuItem, st int, err error)
	}{
		{
			name: "OK",
			req:  menuItem.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetMenuItem(gomock.Any(), menuItem.ID).
					Times(1).
					Return(*menuItem, nil)
			},
			checkResponse: func(t *testing.T, res *entity.GetMenuItem, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, menuItem.Name, res.Name)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "NotFound",
			req:  menuItem.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetMenuItem(gomock.Any(), gomock.Any()).Times(1).Return(db.MenuItem{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, res *entity.GetMenuItem, st int, err error) {
				require.Error(t, err)
				require.Empty(t, res)
				require.Equal(t, http.StatusNotFound, st)
			},
		},
		{
			name: "InternalError",
			req:  menuItem.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().GetMenuItem(gomock.Any(), menuItem.ID).
					Times(1).
					Return(db.MenuItem{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *entity.GetMenuItem, st int, err error) {
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
			res, st, err := u.GetMenuItem(tc.req)

			tc.checkResponse(t, res, st, err)
		})
	}
}
