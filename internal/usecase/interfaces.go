package usecase

import "github.com/zura-t/go_delivery_system-shops/internal/entity"

type Shop interface {
	CreateShop(req *entity.CreateShop) (*entity.Shop, int, error)
	GetShops(limit int32, offset int32) ([]*entity.Shop, int, error)
	GetShopsAdmin(user_id int64) ([]*entity.Shop, int, error)
	GetShopInfo(id int64) (*entity.Shop, int, error)
	UpdateShop(id int64, req *entity.UpdateShopInfo) (*entity.Shop, int, error)
	CreateMenu(req *entity.CreateMenuItem) ([]*entity.GetMenuItem, int, error)
	GetMenu(shopId int64) ([]*entity.GetMenuItem, int, error)
	UpdateMenuItem(id int64, req *entity.UpdateMenuItem) (*entity.GetMenuItem, int, error)
	GetMenuItem(id int64) (*entity.GetMenuItem, int, error)
	DeleteShop(id int64, user_id int64) (string, int, error)
	DeleteMenuItem(id int64, user_id int64) (string, int, error)
}
