package ginroutes

import (
	ginUserHandlers "go-hexa-full/internal/infraestructure/http/gin/handlers/user"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *ginUserHandlers.UserHandler
}

func NewUserRoutes(handler *ginUserHandlers.UserHandler) *UserRoutes {
	return &UserRoutes{handler: handler}
}

func (ur *UserRoutes) RegisterRoutes(r *gin.Engine) {
	r.GET("/user/:id", ur.handler.GetUser)
	r.POST("/user", ur.handler.CreateUser)
}
