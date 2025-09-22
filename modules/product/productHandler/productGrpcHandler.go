package productHandler

import (
	"context"

	productPb "github.com/jetsadawwts/go-microservices/modules/product/productPb"
	"github.com/jetsadawwts/go-microservices/modules/product/productUsecase"
)

type (
	productGrpcHandler struct {
		productPb.UnimplementedProductGrpcServiceServer
		productUsecase productUsecase.ProductUsecaseService
	}
)

func NewproductGrpcHandler(productUsecase productUsecase.ProductUsecaseService) *productGrpcHandler {
	return &productGrpcHandler{
		productUsecase: productUsecase,
	}
}

func (g *productGrpcHandler) FindProductsInIds(ctx context.Context, req *productPb.FindProductsInIdsReq) (*productPb.FindProductsInIdsRes, error) {
	return g.productUsecase.FindProductsInIds(ctx, req)
}
