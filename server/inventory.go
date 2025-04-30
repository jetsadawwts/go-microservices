package server

import (
	"githib.coom/jetsadawwts/go-microservices/modules/inventory/inventoryHandler"
	"githib.coom/jetsadawwts/go-microservices/modules/inventory/inventoryRepository"
	"githib.coom/jetsadawwts/go-microservices/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	inventory.GET("", s.healthCheckService)
}
