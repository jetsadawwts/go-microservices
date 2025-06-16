package userHandler

import (
	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/user/userUsecase"
)

type (
	UserHttpHandlerService interface{}

	userHttpHandler struct {
		cfg         *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{
		cfg:         cfg,
		userUsecase: userUsecase,
	}
}
