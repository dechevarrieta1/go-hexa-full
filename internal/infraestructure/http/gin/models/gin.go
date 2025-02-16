package ginroutesmodels

import "github.com/gin-gonic/gin"

type RouteInterface interface {
	RegisterRoutes(*gin.Engine)
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
