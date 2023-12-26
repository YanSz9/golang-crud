package main

import (
	"github.com/YanSz9/golang-crud/src/controller"
	"github.com/YanSz9/golang-crud/src/model/repository"
	"github.com/YanSz9/golang-crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
