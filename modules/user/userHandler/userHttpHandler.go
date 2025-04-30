package userHandler

import (
	"githib.coom/jetsadawwts/go-microservices/config"
	"githib.coom/jetsadawwts/go-microservices/modules/user/userUsecase"
)


type (
	UserHttpHandlerService interface{}

	userHttpHandler struct {
		cfg *config.Config
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserHttpHandler(cfg *config.Config, userUsecase userUsecase.UserUsecaseService) UserHttpHandlerService {
	return &userHttpHandler{
		cfg: cfg,
		userUsecase: userUsecase,
	}
}

