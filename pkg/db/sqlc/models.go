// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type MenuItem struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Photo       sql.NullString `json:"photo"`
	Price       int32          `json:"price"`
	ShopID      int64          `json:"shop_id"`
	CreatedAt   time.Time      `json:"created_at"`
}

type Shop struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	OpenTime    sql.NullTime   `json:"open_time"`
	CloseTime   sql.NullTime   `json:"close_time"`
	IsClosed    bool           `json:"is_closed"`
	CreatedAt   time.Time      `json:"created_at"`
}