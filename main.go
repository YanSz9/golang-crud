package main

import (
	"context"
	"log"

	"github.com/YanSz9/golang-crud/src/configuration/database/mongodb"
	"github.com/YanSz9/golang-crud/src/configuration/logger"
	"github.com/YanSz9/golang-crud/src/controller"
	"github.com/YanSz9/golang-crud/src/controller/routes"
	"github.com/YanSz9/golang-crud/src/model/repository"
	"github.com/YanSz9/golang-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDbConncetion(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
