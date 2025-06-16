package userHandler

import (
	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/user/userUsecase"
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
