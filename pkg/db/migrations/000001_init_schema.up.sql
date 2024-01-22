CREATE TABLE "shops" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "open_time" timestamptz,
  "close_time" timestamptz,
  "is_closed" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "menuItems" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "photo" varchar,
  "price" int NOT NULL,
  "shop_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "menuItems" ADD FOREIGN KEY ("shop_id") REFERENCES "shops" ("id");