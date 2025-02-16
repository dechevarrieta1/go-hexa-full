package main

import (
	mongoadapter "go-hexa-full/internal/adapters/db/mongo"
	userservice "go-hexa-full/internal/core/user/service"
	"go-hexa-full/internal/infraestructure/http/gin"
	ginUserHandlers "go-hexa-full/internal/infraestructure/http/gin/handlers/user"
	ginroutes "go-hexa-full/internal/infraestructure/http/gin/routes"
	"log"
	"os"
)

var (
	mongoUri = os.Getenv("MONGO_URI")
)

func main() {
	repo, err := mongoadapter.NewMongoRepository(mongoUri, "sample_mflix", "users")
	if err != nil {
		log.Println("[ERROR][main]Can't establish connection with Mongo", err)
	}
	userService := userservice.NewUserService(repo)
	userHandler := ginUserHandlers.NewUserHandler(userService)

	userRoutes := ginroutes.NewUserRoutes(userHandler)
	router := ginroutes.SetupRouter(userRoutes)

	server := gin.NewGinServer(router)
	server.Start()
}
