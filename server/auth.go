package server

import (
	"log"

	"github.com/jetsadawwts/go-microservices/modules/auth/authHandler"
	authPb "github.com/jetsadawwts/go-microservices/modules/auth/authPb"
	"github.com/jetsadawwts/go-microservices/modules/auth/authRepository"
	"github.com/jetsadawwts/go-microservices/modules/auth/authUsecase"
	"github.com/jetsadawwts/go-microservices/pkg/grpcconn"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	//gRPC
	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)
		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	auth.GET("", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(s.healthCheckService, []int{1,0})))
	//auth.GET("/:user_id", s.middleware.JwtAuthorization(s.middleware.UserIdParamValidation(s.middleware.RbacAuthorization(s.healthCheckService, []int{1,0}))))
	auth.POST("/auth/login", httpHandler.Login)
	auth.POST("/auth/refresh-token", httpHandler.RefreshToken)
	auth.POST("/auth/logout", httpHandler.Logout)
}
