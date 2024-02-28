// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: shops.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMenuItem = `-- name: CreateMenuItem :one
INSERT INTO "menuItems" (
  name,
  description,
  photo,
  price,
  shop_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, name, description, photo, price, shop_id, created_at
`

type CreateMenuItemParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Photo       sql.NullString `json:"photo"`
	Price       int32          `json:"price"`
	ShopID      int64          `json:"shop_id"`
}

func (q *Queries) CreateMenuItem(ctx context.Context, arg CreateMenuItemParams) (MenuItem, error) {
	row := q.db.QueryRowContext(ctx, createMenuItem,
		arg.Name,
		arg.Description,
		arg.Photo,
		arg.Price,
		arg.ShopID,
	)
	var i MenuItem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Photo,
		&i.Price,
		&i.ShopID,
		&i.CreatedAt,
	)
	return i, err
}

const createShop = `-- name: CreateShop :one
INSERT INTO shops (
  name,
  description,
  open_time,
  close_time,
  is_closed,
  user_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, name, description, open_time, close_time, is_closed, user_id, created_at
`

type CreateShopParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	OpenTime    sql.NullTime   `json:"open_time"`
	CloseTime   sql.NullTime   `json:"close_time"`
	IsClosed    bool           `json:"is_closed"`
	UserID      int64          `json:"user_id"`
}

func (q *Queries) CreateShop(ctx context.Context, arg CreateShopParams) (Shop, error) {
	row := q.db.QueryRowContext(ctx, createShop,
		arg.Name,
		arg.Description,
		arg.OpenTime,
		arg.CloseTime,
		arg.IsClosed,
		arg.UserID,
	)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.OpenTime,
		&i.CloseTime,
		&i.IsClosed,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMenuItem = `-- name: DeleteMenuItem :exec
DELETE FROM "menuItems"
WHERE id = $1
`

func (q *Queries) DeleteMenuItem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMenuItem, id)
	return err
}

const deleteShop = `-- name: DeleteShop :exec
DELETE FROM shops
WHERE id = $1 and user_id = $2
`

type DeleteShopParams struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

func (q *Queries) DeleteShop(ctx context.Context, arg DeleteShopParams) error {
	_, err := q.db.ExecContext(ctx, deleteShop, arg.ID, arg.UserID)
	return err
}

const getMenuItem = `-- name: GetMenuItem :one
SELECT id, name, description, photo, price, shop_id, created_at FROM "menuItems"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMenuItem(ctx context.Context, id int64) (MenuItem, error) {
	row := q.db.QueryRowContext(ctx, getMenuItem, id)
	var i MenuItem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Photo,
		&i.Price,
		&i.ShopID,
		&i.CreatedAt,
	)
	return i, err
}

const getShop = `-- name: GetShop :one
SELECT id, name, description, open_time, close_time, is_closed, user_id, created_at FROM shops
WHERE id = $1
`

func (q *Queries) GetShop(ctx context.Context, id int64) (Shop, error) {
	row := q.db.QueryRowContext(ctx, getShop, id)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.OpenTime,
		&i.CloseTime,
		&i.IsClosed,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getShopWithMenuItemId = `-- name: GetShopWithMenuItemId :one
SELECT "menuItems".id, "menuItems".name, "menuItems".description, photo, price, shop_id, "menuItems".created_at, shops.id, shops.name, shops.description, open_time, close_time, is_closed, user_id, shops.created_at FROM "menuItems"
JOIN shops ON "menuItems".shop_id = shops.id
WHERE "menuItems".id = $1
`

type GetShopWithMenuItemIdRow struct {
	ID            int64          `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	Photo         sql.NullString `json:"photo"`
	Price         int32          `json:"price"`
	ShopID        int64          `json:"shop_id"`
	CreatedAt     time.Time      `json:"created_at"`
	ID_2          int64          `json:"id_2"`
	Name_2        string         `json:"name_2"`
	Description_2 sql.NullString `json:"description_2"`
	OpenTime      sql.NullTime   `json:"open_time"`
	CloseTime     sql.NullTime   `json:"close_time"`
	IsClosed      bool           `json:"is_closed"`
	UserID        int64          `json:"user_id"`
	CreatedAt_2   time.Time      `json:"created_at_2"`
}

func (q *Queries) GetShopWithMenuItemId(ctx context.Context, id int64) (GetShopWithMenuItemIdRow, error) {
	row := q.db.QueryRowContext(ctx, getShopWithMenuItemId, id)
	var i GetShopWithMenuItemIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Photo,
		&i.Price,
		&i.ShopID,
		&i.CreatedAt,
		&i.ID_2,
		&i.Name_2,
		&i.Description_2,
		&i.OpenTime,
		&i.CloseTime,
		&i.IsClosed,
		&i.UserID,
		&i.CreatedAt_2,
	)
	return i, err
}

const getShopsAdmin = `-- name: GetShopsAdmin :many
SELECT id, name, description, open_time, close_time, is_closed, user_id, created_at FROM shops
WHERE user_id = $1
`

func (q *Queries) GetShopsAdmin(ctx context.Context, userID int64) ([]Shop, error) {
	rows, err := q.db.QueryContext(ctx, getShopsAdmin, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Shop{}
	for rows.Next() {
		var i Shop
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.OpenTime,
			&i.CloseTime,
			&i.IsClosed,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMenuItems = `-- name: ListMenuItems :many
SELECT id, name, description, photo, price, shop_id, created_at FROM "menuItems"
WHERE shop_id = $1
ORDER BY name
`

func (q *Queries) ListMenuItems(ctx context.Context, shopID int64) ([]MenuItem, error) {
	rows, err := q.db.QueryContext(ctx, listMenuItems, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MenuItem{}
	for rows.Next() {
		var i MenuItem
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Photo,
			&i.Price,
			&i.ShopID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listShops = `-- name: ListShops :many
SELECT id, name, description, open_time, close_time, is_closed, user_id, created_at FROM shops
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListShopsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListShops(ctx context.Context, arg ListShopsParams) ([]Shop, error) {
	rows, err := q.db.QueryContext(ctx, listShops, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Shop{}
	for rows.Next() {
		var i Shop
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.OpenTime,
			&i.CloseTime,
			&i.IsClosed,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMenuItem = `-- name: UpdateMenuItem :one
UPDATE "menuItems"
SET
name = COALESCE($1, name),
description = COALESCE($2, description),
price = COALESCE($3, price)
WHERE id = $4
RETURNING id, name, description, photo, price, shop_id, created_at
`

type UpdateMenuItemParams struct {
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	Price       sql.NullInt32  `json:"price"`
	ID          int64          `json:"id"`
}

func (q *Queries) UpdateMenuItem(ctx context.Context, arg UpdateMenuItemParams) (MenuItem, error) {
	row := q.db.QueryRowContext(ctx, updateMenuItem,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ID,
	)
	var i MenuItem
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Photo,
		&i.Price,
		&i.ShopID,
		&i.CreatedAt,
	)
	return i, err
}

const updateShop = `-- name: UpdateShop :one
UPDATE shops
SET
name = COALESCE($1, name),
description = COALESCE($2, description),
open_time = COALESCE($3, open_time),
close_time = COALESCE($4, close_time),
is_closed = COALESCE($5, is_closed)
WHERE id = $6 and user_id = $7
RETURNING id, name, description, open_time, close_time, is_closed, user_id, created_at
`

type UpdateShopParams struct {
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
	OpenTime    sql.NullTime   `json:"open_time"`
	CloseTime   sql.NullTime   `json:"close_time"`
	IsClosed    sql.NullBool   `json:"is_closed"`
	ID          int64          `json:"id"`
	UserID      int64          `json:"user_id"`
}

func (q *Queries) UpdateShop(ctx context.Context, arg UpdateShopParams) (Shop, error) {
	row := q.db.QueryRowContext(ctx, updateShop,
		arg.Name,
		arg.Description,
		arg.OpenTime,
		arg.CloseTime,
		arg.IsClosed,
		arg.ID,
		arg.UserID,
	)
	var i Shop
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.OpenTime,
		&i.CloseTime,
		&i.IsClosed,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
