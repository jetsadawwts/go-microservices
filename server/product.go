package server

import (
	"githib.coom/jetsadawwts/go-microservices/modules/product/productHandler"
	"githib.coom/jetsadawwts/go-microservices/modules/product/productRepository"
	"githib.coom/jetsadawwts/go-microservices/modules/product/productUsecase"
)

func (s * server) productService() {
	repo := productRepository.NewProductRepository(s.db)
	usecase := productUsecase.NewProductUsecase(repo)
	httpHandler := productHandler.NewProductHttpHandler(s.cfg, usecase)
	grpcHandler := productHandler.NewproductGrpcHandler(usecase)


	_ = httpHandler
	_ = grpcHandler


	product := s.app.Group("/product_v1")

	product.GET("", s.healthCheckService)
}