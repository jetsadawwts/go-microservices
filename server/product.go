package server

import (
	"log"

	"github.com/jetsadawwts/go-microservices/modules/product/productHandler"
	productPb "github.com/jetsadawwts/go-microservices/modules/product/productPb"
	"github.com/jetsadawwts/go-microservices/modules/product/productRepository"
	"github.com/jetsadawwts/go-microservices/modules/product/productUsecase"
	"github.com/jetsadawwts/go-microservices/pkg/grpcconn"
)

func (s *server) productService() {
	repo := productRepository.NewProductRepository(s.db)
	usecase := productUsecase.NewProductUsecase(repo)
	httpHandler := productHandler.NewProductHttpHandler(s.cfg, usecase)
	grpcHandler := productHandler.NewproductGrpcHandler(usecase)

	//gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ProductUrl)
		productPb.RegisterProductGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Product gRPC server listening on %s", s.cfg.Grpc.ProductUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	product := s.app.Group("/product_v1")

	product.GET("", s.healthCheckService)
}
