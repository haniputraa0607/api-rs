package userhandler

import (
	"api-rs/schemas"
	userservice "api-rs/services/user"
	"api-rs/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	ListUser(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type userHandler struct {
	userService userservice.UserService
}

func NewUserHandler(userService userservice.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) ListUser(ctx *gin.Context) {
	users, err := h.userService.ListUser()
	if err != nil {
		utility.ApiResponse(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	utility.ApiResponse(ctx, http.StatusOK, "success", users, nil)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var request schemas.LoginRequest
	err := utility.ShouldBind(ctx, &request)
	if err != nil {
		error := utility.FormatValidationError(err)
		utility.ApiResponse(ctx, http.StatusBadRequest, error, nil, nil)
		return
	}

	response, err := h.userService.Login(request.Username, request.Password)
	if err != nil {
		utility.ApiResponse(ctx, http.StatusBadRequest, err.Error(), response, nil)
		return
	}

	utility.ApiResponse(ctx, http.StatusOK, "success", response, nil)
}
