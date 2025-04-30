package inventoryHandler

import "githib.coom/jetsadawwts/go-microservices/modules/inventory/inventoryUsecase"

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) inventoryGrpcHandler {
	return inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
