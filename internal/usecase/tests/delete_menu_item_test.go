package usecase_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/zura-t/go_delivery_system-shops/pkg/db/mock"
)

type DeleteMenu struct {
	Id     int64
	UserId int64
}

func Test_delete_menu_item(t *testing.T) {
	userId := randomShop().UserID
	menuItem := randomMenuItem()
	r := DeleteMenu{
		Id:     menuItem.ID,
		UserId: userId,
	}

	tests := []struct {
		name          string
		req           DeleteMenu
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res string, st int, err error)
	}{
		{
			name: "OK",
			req:  r,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteMenuItem(gomock.Any(), menuItem.ID).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, res string, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "NotFound",
			req:  r,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteMenuItem(gomock.Any(), menuItem.ID).Times(1).Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, res string, st int, err error) {
				require.Error(t, err)
				require.Empty(t, res)
				require.Equal(t, http.StatusNotFound, st)
			},
		},
		{
			name: "InternalError",
			req:  r,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteMenuItem(gomock.Any(), menuItem.ID).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res string, st int, err error) {
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
			res, st, err := u.DeleteMenuItem(tc.req.Id, tc.req.UserId)

			tc.checkResponse(t, res, st, err)
		})
	}
}
