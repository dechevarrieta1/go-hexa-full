package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
}

func NewGinServer(router *gin.Engine) *GinServer {
	return &GinServer{router: router}
}

func (s *GinServer) Start() {
	fmt.Println("Server running on http://localhost:8080")
	s.router.Run(":8080")
}
