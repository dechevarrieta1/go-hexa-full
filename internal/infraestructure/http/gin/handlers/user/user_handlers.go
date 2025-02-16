package ginuserhandlers

import (
	userModels "go-hexa-full/internal/core/user/models"
	userService "go-hexa-full/internal/core/user/service"
	ginutils "go-hexa-full/internal/infraestructure/http/gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user userModels.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ginutils.GinResponseHandler(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(user); err != nil {
		ginutils.GinResponseHandler(c, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	ginutils.GinResponseHandler(c, user, nil, http.StatusCreated, "User created successfully")
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		ginutils.GinResponseHandler(c, nil, err.Error(), http.StatusNotFound)
		return
	}

	ginutils.GinResponseHandler(c, user, nil, http.StatusOK, "User retrieved successfully")
}
