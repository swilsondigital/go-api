package routes

import (
	"goapi/controllers"
	"goapi/database"
	"goapi/pages"
	"goapi/repository"

	"github.com/gin-gonic/gin"
)

/**
* Setup routing
 **/
func Setup(r *gin.Engine) {
	// General Page Endpoints
	r.GET("/", pages.ShowHomePage)

	// User Endpoints
	userRepository := repository.NewUserRepository(database.DB)
	userController := controllers.NewUserController(userRepository)
	userRouter := r.Group("/users")
	{
		userRouter.GET("/", userController.GetAllUsers)
		userRouter.POST("/", userController.CreateUser)
		userRouter.GET("/:id", userController.GetUserById)
		userRouter.PUT("/:id", userController.UpdateUser)
		userRouter.DELETE("/:id", userController.DeleteUser)
	}

	// Client Endpoints
	clientRepository := repository.NewClientRepository(database.DB)
	clientController := controllers.NewClientController(clientRepository)
	clientRouter := r.Group("/clients")
	{
		clientRouter.GET("/", clientController.GetAllClients)
		clientRouter.POST("/", clientController.CreateClient)
		clientRouter.GET("/:id", clientController.GetClientById)
		clientRouter.PUT("/:id", clientController.UpdateClient)
		clientRouter.DELETE("/:id", clientController.DeleteClient)
	}

}
