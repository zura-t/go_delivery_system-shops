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

type UpdateMenuItemRequest struct {
	Id   int64
	Data *entity.UpdateMenuItem
}

func Test_update_menu_item(t *testing.T) {
	menuItem := randomMenuItem()
	newName := pkg.RandomString(6)
	newPrice := int32(pkg.RandomInt(5, 20))

	tests := []struct {
		name          string
		req           *UpdateMenuItemRequest
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, res *entity.MenuItem, st int, err error)
	}{
		{
			name: "OK",
			req: &UpdateMenuItemRequest{
				Id: menuItem.ID,
				Data: &entity.UpdateMenuItem{
					Name:        newName,
					Description: "newDescription",
					Price:       newPrice,
				},
			},
			buildStub: func(store *mockdb.MockStore) {
				req := db.UpdateMenuItemParams{
					ID: menuItem.ID,
					Name: sql.NullString{
						Valid:  true,
						String: newName,
					},
					Description: sql.NullString{
						Valid:  true,
						String: "newDescription",
					},
					Price: sql.NullInt32{
						Valid: true,
						Int32: newPrice,
					},
				}
				menuItemUpdated := db.MenuItem{
					ID:   menuItem.ID,
					Name: newName,
					Description: sql.NullString{
						Valid:  true,
						String: "newDescription",
					},
					Price: newPrice,
				}

				store.EXPECT().UpdateMenuItem(gomock.Any(), req).
					Times(1).
					Return(menuItemUpdated, nil)
			},
			checkResponse: func(t *testing.T, res *entity.MenuItem, st int, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				updatedMenuItem := res
				require.Equal(t, newName, updatedMenuItem.Name)
				require.Equal(t, newPrice, updatedMenuItem.Price)
				require.Equal(t, http.StatusOK, st)
			},
		},
		{
			name: "NotFound",
			req: &UpdateMenuItemRequest{
				Id: menuItem.ID,
				Data: &entity.UpdateMenuItem{
					Name:        newName,
					Description: "newDescription",
					Price:       newPrice,
				},
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateMenuItem(gomock.Any(), gomock.Any()).Times(1).Return(db.MenuItem{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, res *entity.MenuItem, st int, err error) {
				require.Error(t, err)
				require.Empty(t, res)
				require.Equal(t, http.StatusNotFound, st)
			},
		},
		{
			name: "InternalError",
			req: &UpdateMenuItemRequest{
				Id:   menuItem.ID,
				Data: &entity.UpdateMenuItem{},
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateMenuItem(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.MenuItem{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, res *entity.MenuItem, st int, err error) {
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
			res, st, err := u.UpdateMenuItem(tc.req.Id, tc.req.Data)

			tc.checkResponse(t, res, st, err)
		})
	}
}
