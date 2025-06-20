package userHandler

import (
	"context"
	"net/http"
	"strings"

	"github.com/jetsadawwts/go-microservices/config"
	"github.com/jetsadawwts/go-microservices/modules/user"
	"github.com/jetsadawwts/go-microservices/modules/user/userUsecase"
	"github.com/jetsadawwts/go-microservices/pkg/request"
	"github.com/jetsadawwts/go-microservices/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	UserHttpHandlerService interface {
		CreateUser(c echo.Context) error
		FindOneUserProfile(c echo.Context) error
		AddUserMoney(c echo.Context) error
		GetUserSavingAccount(c echo.Context) error
	}

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

func (h *userHttpHandler) CreateUser(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.CreateUser(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) FindOneUserProfile(c echo.Context) error {
	ctx := context.Background()

	userId := strings.TrimPrefix(c.Param("user_id"), "user:")

	res, err := h.userUsecase.FindOneUserProfile(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *userHttpHandler) AddUserMoney(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(user.CreateUserTransactionReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.userUsecase.AddUserMoney(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *userHttpHandler) GetUserSavingAccount(c echo.Context) error {
	ctx := context.Background()

	userId := c.Param("user_id")

	res, err := h.userUsecase.GetUserSavingAccount(ctx, userId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
