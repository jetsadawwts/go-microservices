package productHandler

import (
	"githib.coom/jetsadawwts/go-microservices/config"
	"githib.coom/jetsadawwts/go-microservices/modules/product/productUsecase"
)

type (
	ProductHttpHandlerService interface {}

	productHttpHandler struct {
		cfg *config.Config
		productUsecase productUsecase.ProductUsecaseService
	}
)

func NewProductHttpHandler(cfg *config.Config, productUsecase productUsecase.ProductUsecaseService) ProductHttpHandlerService {
	return &productHttpHandler{
		cfg: cfg,
		productUsecase: productUsecase,
	}
}