package inventoryHandler

import (
	"context"

	"net/http"
	"strings"

	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/inventory"
	"github.com/jetsadawwts/go-microservices/modules/inventory/inventoryUsecase"
	"github.com/jetsadawwts/go-microservices/pkg/request"
	"github.com/jetsadawwts/go-microservices/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	InventoryHttpHandlerService interface{
		FindUserProducts(c echo.Context) error 
	}
	inventoryHttpHandler        struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}

func (h *inventoryHttpHandler) FindUserProducts(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(inventory.InventorySearchReq)
	userId := strings.TrimPrefix(c.Param("user_id"), "user:");

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.inventoryUsecase.FindUserProducts(ctx, h.cfg, userId, req)

	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
