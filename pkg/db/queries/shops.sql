-- name: CreateShop :one
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
RETURNING *;

-- name: GetShop :one
SELECT * FROM shops
WHERE id = $1 LIMIT 1;

-- name: ListShops :many
SELECT * FROM shops
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: GetShopsAdmin :many
SELECT * FROM shops
WHERE user_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateShop :one
UPDATE shops
SET
name = COALESCE(sqlc.narg(name), name),
description = COALESCE(sqlc.narg(description), description),
open_time = COALESCE(sqlc.narg(open_time), open_time),
close_time = COALESCE(sqlc.narg(close_time), close_time),
is_closed = COALESCE(sqlc.narg(is_closed), is_closed)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: CreateMenuItem :one
INSERT INTO "menuItems" (
  name,
  description,
  photo,
  price,
  shop_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteShop :exec
DELETE FROM shops
WHERE id = $1;

-- name: GetMenuItem :one
SELECT * FROM "menuItems"
WHERE id = $1 LIMIT 1;

-- name: ListMenuItems :many
SELECT * FROM "menuItems"
WHERE shop_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateMenuItem :one
UPDATE "menuItems"
SET
name = COALESCE(sqlc.narg(name), name),
description = COALESCE(sqlc.narg(description), description),
price = COALESCE(sqlc.narg(price), price)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteMenuItem :exec
DELETE FROM "menuItems"
WHERE id = $1;