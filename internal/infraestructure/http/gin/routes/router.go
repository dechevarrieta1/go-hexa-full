package GinRoutes

import (
	GinMiddlewares "go-hexa-full/internal/infraestructure/http/gin/middlewares"
	GinModels "go-hexa-full/internal/infraestructure/http/gin/models"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRouter(routes []GinModels.Route) *gin.Engine {
	r := gin.Default()

	// Middlewares
	r.Use(GinMiddlewares.LoggerMiddleware())

	// Rutas
	for _, route := range routes {
		switch route.Method {
		case "GET":
			r.GET(route.Path, route.Handler)
		case "POST":
			r.POST(route.Path, route.Handler)
		case "PUT":
			r.PUT(route.Path, route.Handler)
		case "DELETE":
			r.DELETE(route.Path, route.Handler)
		default:
			log.Println("[LOG][INFO] Invalid HTTP method") //todo -> check the possibility of not log this bc a possible DDOS attack can print a lot this line....
		}
	}
	return r
}
