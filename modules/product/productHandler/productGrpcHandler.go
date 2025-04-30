package productHandler

import "githib.coom/jetsadawwts/go-microservices/modules/product/productUsecase"

type (
	productGrpcHandler struct {
		productUsecase productUsecase.ProductUsecaseService
	}
)

func NewproductGrpcHandler(productUsecase productUsecase.ProductUsecaseService) productGrpcHandler {
	return productGrpcHandler{
		productUsecase: productUsecase,
	}
}