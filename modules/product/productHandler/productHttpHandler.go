package productHandler

import (
	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/product/productUsecase"
)

type (
	ProductHttpHandlerService interface{}

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
