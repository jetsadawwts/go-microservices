package inventoryHandler

import "github.com/jetsadawwts/go-microservices/modules/inventory/inventoryUsecase"

type (
	inventoryQueueHandler struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) inventoryQueueHandler {
	return inventoryQueueHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
