package userHandler

import "githib.coom/jetsadawwts/go-microservices/modules/user/userUsecase"

type (
	userGrpcHandler struct {
		userUsecase userUsecase.UserUsecaseService
	}
)

func NewUserGrpcHandler(userUsecase userUsecase.UserUsecaseService) userGrpcHandler {
	return userGrpcHandler{
		userUsecase: userUsecase,
	}
}
