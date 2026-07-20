package di

import (
	"koda-b8-backend1/internal/handler"
	"koda-b8-backend1/internal/model"
	"koda-b8-backend1/internal/repo"
	"koda-b8-backend1/internal/svc"
)

type Container struct {
	userData []model.User

	userRepo    *repo.UserRepo
	userService *svc.UserService
	userHandler *handler.UserHandler
}

func (c *Container) initDeps() {
	c.userRepo = repo.NewUserRepo(&c.userData)
	c.userService = svc.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func NewContainer() *Container {
	container := &Container{
		userData: []model.User{},
	}

	container.initDeps()

	return container
}