package productUsecase

import "github.com/jetsadawwts/go-microservices/modules/product/productRepository"

type (
	ProductUsecaseService interface{}

	productUsecase struct {
		productRepository productRepository.ProductRepositoryService
	}
)

func NewProductUsecase(productRepository productRepository.ProductRepositoryService) ProductUsecaseService {
	return &productUsecase{
		productRepository: productRepository,
	}
}
