package ginroutes

import (
	ginRoutesModels "go-hexa-full/internal/infraestructure/http/gin/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter(routeModules ...ginRoutesModels.RouteInterface) *gin.Engine {
	r := gin.Default()

	for _, module := range routeModules {
		module.RegisterRoutes(r)
	}

	return r
}
