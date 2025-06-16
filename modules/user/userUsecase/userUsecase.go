package userUsecase

import "github.com/jetsadawwts/go-microservices/modules/user/userRepository"

type (
	UserUsecaseService interface{}

	userUsecase struct {
		userRepository userRepository.UserRepositoryService
	}
)

func NewUserUsecase(userRepository userRepository.UserRepositoryService) UserUsecaseService {
	return &userUsecase{
		userRepository: userRepository,
	}
}
