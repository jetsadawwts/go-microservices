package productHandler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/product"
	"github.com/jetsadawwts/go-microservices/modules/product/productUsecase"
	"github.com/jetsadawwts/go-microservices/pkg/request"
	"github.com/jetsadawwts/go-microservices/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	ProductHttpHandlerService interface{
		CreateProduct(c echo.Context) error
		FindOneProduct(c echo.Context) error
		FindManyProducts(c echo.Context) error
		EditProduct(c echo.Context) error 
		EnableOrDisableProduct(c echo.Context) error 
	}

	productHttpHandler struct {
		cfg            *config.Config
		productUsecase productUsecase.ProductUsecaseService
	}
)

func NewProductHttpHandler(cfg *config.Config, productUsecase productUsecase.ProductUsecaseService) ProductHttpHandlerService {
	return &productHttpHandler{
		cfg:            cfg,
		productUsecase: productUsecase,
	}
}

func (h *productHttpHandler) CreateProduct(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(product.CreateProductReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.productUsecase.CreateProduct(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *productHttpHandler) FindOneProduct(c echo.Context) error {
	ctx := context.Background()

	productId := strings.TrimPrefix(c.Param("product_id"), "product:")

	res, err := h.productUsecase.FindOneProduct(ctx, productId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *productHttpHandler) FindManyProducts(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(product.ProductSearchReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.productUsecase.FindManyProducts(ctx, h.cfg.Paginate.ProductNextPageBasedUrl,req)

	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *productHttpHandler) EditProduct(c echo.Context) error {
	ctx := context.Background()

	productId := strings.TrimPrefix(c.Param("product_id"), "product:")

	wrapper := request.ContextWrapper(c)

	req := new(product.ProductUpdateReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.productUsecase.EditProduct(ctx, productId, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *productHttpHandler) EnableOrDisableProduct(c echo.Context) error {
	ctx := context.Background()

	productId := strings.TrimPrefix(c.Param("product_id"), "product:")

	res, err := h.productUsecase.EnableOrDisableProduct(ctx, productId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"message": fmt.Sprintf("product: %s is successfully is activated to: %s", productId, res),
	})
}

