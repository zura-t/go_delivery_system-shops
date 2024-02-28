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

func newShopRoutes(handler *gin.Engine, shopUsecase usecase.Shop, logger logger.Interface) {
	routes := &shopRoutes{shopUsecase, logger}

	shopsRoutes := handler.Group("/shops")

	shopsRoutes.POST("/", routes.createShop)
	shopsRoutes.GET("/:id", routes.getShopInfo)
	shopsRoutes.GET("/", routes.getShops)
	shopsRoutes.GET("/admin", routes.getShopsAdmin)
	shopsRoutes.PATCH("/:id", routes.updateShop)
	shopsRoutes.DELETE("/:id", routes.deleteShop)

	menuItemRoutes := handler.Group("/menu_items")

	menuItemRoutes.POST("/", routes.createMenu)
	menuItemRoutes.GET("/list/:id", routes.getMenu)
	menuItemRoutes.PATCH("/:id", routes.updateMenuItem)
	menuItemRoutes.GET("/:id", routes.getMenuItem)
	menuItemRoutes.DELETE("/:id", routes.deleteMenuItem)
}

type CreateShopRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	OpenTime    time.Time `json:"open_time" binding:"required"`
	CloseTime   time.Time `json:"close_time" binding:"required"`
	UserId      int64     `json:"user_id" binding:"required,min=1"`
	IsClosed    bool      `json:"is_closed"`
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
		UserId:      req.UserId,
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

type GetShopsRequest struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

func (r *shopRoutes) getShops(ctx *gin.Context) {
	var req GetShopsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	shops, st, err := r.shopUsecase.GetShops(req.Limit, req.Offset)

	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - GetShops")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, shops)
}

type UserIdQuery struct {
	UserId int64 `form:"user_id" binding:"required,min=1"`
}

func (r *shopRoutes) getShopsAdmin(ctx *gin.Context) {
	var req UserIdQuery
	if err := ctx.ShouldBind(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	shops, st, err := r.shopUsecase.GetShopsAdmin(req.UserId)
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

type MenuItem struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Price       int32  `json:"price" binding:"required,min=1"`
}

type CreateMenuRequest struct {
	MenuItems []MenuItem `json:"menu_items" binding:"required,min=1"`
	ShopId    int64      `json:"shop_id" binding:"required,min=1"`
	UserId    int64      `json:"user_id" binding:"required,min=1"`
}

func (r *shopRoutes) createMenu(ctx *gin.Context) {
	var reqs CreateMenuRequest
	if err := ctx.ShouldBind(&reqs); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var menuItems = make([]entity.MenuItem, len(reqs.MenuItems))
	for i := 0; i < len(reqs.MenuItems); i++ {
		menuItems[i] = entity.MenuItem{
			Name:        reqs.MenuItems[i].Name,
			Description: reqs.MenuItems[i].Description,
			Price:       reqs.MenuItems[i].Price,
		}
	}

	menuCreated, st, err := r.shopUsecase.CreateMenu(&entity.CreateMenuItem{
		MenuItems: menuItems,
		ShopId:    reqs.ShopId,
		UserId:    reqs.UserId,
	})
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - CreateMenu")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, menuCreated)
}

type GetMenuReq struct {
	ShopId int64 `uri:"id" binding:"required,min=1"`
}

func (r *shopRoutes) getMenu(ctx *gin.Context) {
	var req GetMenuReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	menuitems, st, err := r.shopUsecase.GetMenu(req.ShopId)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - GetMenu")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, menuitems)
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
	UserId      int64     `json:"user_id" binding:"required,min=1"`
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
		UserId:      req.UserId,
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
	UserId      int64  `json:"user_id" binding:"required,min=1"`
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
		UserId:      req.UserId,
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

	var params UserIdQuery
	if err := ctx.ShouldBindQuery(&params); err != nil {
		r.logger.Error(err, "http - v1 - shop routes - deleteShop")
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, st, err := r.shopUsecase.DeleteShop(req.Id, params.UserId)
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

	var reqQ UserIdQuery
	if err := ctx.ShouldBindQuery(&reqQ); err != nil {
		r.logger.Error(err, "http - v1 - shop routes - deleteShop")
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, st, err := r.shopUsecase.DeleteMenuItem(req.Id, reqQ.UserId)
	if err != nil {
		r.logger.Error(err, "package: v1", "http - v1 - DeleteMenuItem()")
		errorResponse(ctx, st, err.Error())
		return
	}

	ctx.JSON(st, res)
}
