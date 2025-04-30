package userHandler

import (
	"githib.coom/jetsadawwts/go-microservices/config"
	"githib.coom/jetsadawwts/go-microservices/modules/user/userUsecase"
)

type (
	userQueueHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserQueueHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) userQueueHandler {
	return userQueueHandler{
		cfg:         cfg,
		userUsecase: userUsecase,
	}
}
