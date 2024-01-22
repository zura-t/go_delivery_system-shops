package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	"github.com/zura-t/go_delivery_system-shops/internal/usecase"
	"github.com/zura-t/go_delivery_system-shops/pkg/logger"
)

type shopRoutes struct {
	shopUsecase usecase.Shop
	logger      logger.Interface
}

func newShopRoutes(handler *gin.RouterGroup, shopUsecase usecase.Shop, logger logger.Interface) {
	routes := &shopRoutes{shopUsecase, logger}

	handler.POST("/shops", routes.createShop)
	// handler.POST("/shops", routes.loginUser)
	// handler.POST("/shops", routes.logout)
	// handler.POST("/renew_token", server.renewAccessToken)
}

type CreateShopRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	OpenTime    time.Time `json:"open_time" binding:"required"`
	CloseTime   time.Time `json:"close_time" binding:"required"`
	IsClosed    bool      `json:"is_closed" binding:"required"`
}

func (r *shopRoutes) createShop(ctx *gin.Context) {
	var req CreateShopRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	shop, st, err := r.shopUsecase.CreateShop(&entity.CreateShop{
		Name:        req.Name,
		Description: req.Description,
		OpenTime:    req.OpenTime,
		CloseTime:   req.CloseTime,
		IsClosed:    req.IsClosed,
	})
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - CreateShop")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, shop)
}

type IdParam struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (r *shopRoutes) getShops(ctx *gin.Context) {
	// var req
	shops, st, err := r.shopUsecase.GetShops()
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - GetShops")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, shops)
}

func (r *shopRoutes) getShopInfo(ctx *gin.Context) {
	var req IdParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	shop, st, err := r.shopUsecase.GetShopInfo(req.Id)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - GetShopInfo")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, shop)
}

type CreateMenuRequest struct {
	menuItems []struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Photo       string `json:"photo"`
		Price       int32  `json:"price" binding:"required,min=1"`
		ShopID      int64  `json:"shop_id" binding:"required,min=1"`
	}
}

func (r *shopRoutes) createMenu(ctx *gin.Context) {
	var req CreateMenuRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	request := make([]*entity.CreateMenuItem, len(req.menuItems))
	for i := 0; i < len(request); i++ {
		request[i] = &entity.CreateMenuItem{
			Name:        req.menuItems[i].Name,
			Description: req.menuItems[i].Description,
			Price:       req.menuItems[i].Price,
		}
	}
	menuCreated, st, err := r.shopUsecase.CreateMenu(request)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - CreateMenu")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, menuCreated)
}

type GetMenuReq struct {
	ShopId int64 `uri:"id"  binding:"required,min=1"`
}

func (r *shopRoutes) getMenu(ctx *gin.Context) {
	var req GetMenuReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	menu, st, err := r.shopUsecase.GetMenu(req.ShopId)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - GetMenu")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, menu)
}

func (r *shopRoutes) getMenuItem(ctx *gin.Context) {
	var req IdParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	menuItem, st, err := r.shopUsecase.GetMenuItem(req.Id)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - GetMenuItem")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, menuItem)
}

type UpdateShopRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OpenTime    time.Time `json:"open_time"`
	CloseTime   time.Time `json:"close_time"`
	IsClosed    bool      `json:"is_closed"`
}

func (r *shopRoutes) updateShop(ctx *gin.Context) {
	var params IdParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var req UpdateShopRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	request := &entity.UpdateShopInfo{
		Name:        req.Name,
		Description: req.Description,
		OpenTime:    req.OpenTime,
		CloseTime:   req.CloseTime,
		IsClosed:    req.IsClosed,
	}
	shopUpdated, st, err := r.shopUsecase.UpdateShop(params.Id, request)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - UpdateShop")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, shopUpdated)
}

type UpdateMenuItemRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Price       int32  `json:"price"`
}

func (r *shopRoutes) updateMenuItem(ctx *gin.Context) {
	var params IdParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var req UpdateMenuItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	request := &entity.UpdateMenuItem{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
	shopUpdated, st, err := r.shopUsecase.UpdateMenuItem(params.Id, request)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - UpdateMenuItem")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, shopUpdated)
}

func (r *shopRoutes) deleteShop(ctx *gin.Context) {
	var req IdParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res, st, err := r.shopUsecase.DeleteShop(req.Id)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - DeleteShop")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, res)
}

func (r *shopRoutes) deleteMenuItem(ctx *gin.Context) {
	var req IdParam
	if err := ctx.ShouldBindUri(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res, st, err := r.shopUsecase.DeleteMenuItem(req.Id)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - DeleteMenuItem()")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, res)
}
