package productUsecase

import "githib.coom/jetsadawwts/go-microservices/modules/product/productRepository"


type (
	ProductUsecaseService interface {}

	productUsecase struct {
		productRepository productRepository.ProductRepositoryService
	}
)

func NewProductUsecase(productRepository productRepository.ProductRepositoryService) ProductUsecaseService {
	return &productUsecase{
		productRepository: productRepository,
	}
}