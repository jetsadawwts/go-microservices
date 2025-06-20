package server

import (
	"log"

	"github.com/jetsadawwts/go-microservices/modules/user/userHandler"
	userPb "github.com/jetsadawwts/go-microservices/modules/user/userPb"
	"github.com/jetsadawwts/go-microservices/modules/user/userRepository"
	"github.com/jetsadawwts/go-microservices/modules/user/userUsecase"
	"github.com/jetsadawwts/go-microservices/pkg/grpcconn"
)

func (s *server) userService() {
	repo := userRepository.NewUserRepository(s.db)
	usecase := userUsecase.NewUserUsecase(repo)
	httpHandler := userHandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userHandler.NewUserGrpcHandler(usecase)
	queueHandler := userHandler.NewUserQueueHandler(s.cfg, usecase)

	//gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.UserUrl)
		userPb.RegisterUserGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("User gRPC server listening on %s", s.cfg.Grpc.UserUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler
	_ = queueHandler

	user := s.app.Group("/user_v1")

	// Health Check
	user.GET("", s.healthCheckService)

	user.POST("/user/register", httpHandler.CreateUser)
	user.GET("/user/:user_id", httpHandler.FindOneUserProfile)
	user.POST("/user/add-money", httpHandler.AddUserMoney)
	user.GET("/user/account/:user_id", httpHandler.GetUserSavingAccount)

}
