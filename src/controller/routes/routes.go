package routes

import (
	"github.com/YanSz9/golang-crud/src/controller"
	"github.com/YanSz9/golang-crud/src/model"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/login", userController.LoginUser)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
