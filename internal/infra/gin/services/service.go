package GinService

import UserService "go-hexa-full/internal/core/user/service"

type GinService struct {
	UserService *UserService.UserService
}
