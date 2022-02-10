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

	// User Endpoints - https://domain/users/:id
	userRepository := repository.NewUserRepository(database.DB)
	userController := controllers.NewUserController(userRepository)
	userRouter := r.Group("/users")
	{
		userRouter.GET("/", userController.GetAllUsers)      // Index
		userRouter.POST("/", userController.CreateUser)      // Create
		userRouter.GET("/:id", userController.GetUserById)   // Read
		userRouter.PUT("/:id", userController.UpdateUser)    // Update
		userRouter.DELETE("/:id", userController.DeleteUser) // Delete
	}

	// Project Repository & Controller - https://domain/projects/:id
	projectRepository := repository.NewProjectRepository(database.DB)
	projectController := controllers.NewProjectController(projectRepository)
	projectRouter := r.Group("/projects")
	{
		projectRouter.GET("/", projectController.GetAllProjects)      // Index
		projectRouter.GET("/:id", projectController.GetProjectById)   // Read
		projectRouter.PUT("/:id", projectController.UpdateProject)    // Update
		projectRouter.DELETE("/:id", projectController.DeleteProject) // Delete

		// Project PortfolioRecords Endpoints - https://domain/projects/:id/records
		projectPortfolioRouter := projectRouter.Group("/:id/records")
		{
			projectPortfolioRouter.POST("/")
		}
	}

	// Client Endpoints - https://domain/clients/:id
	clientRepository := repository.NewClientRepository(database.DB)
	clientController := controllers.NewClientController(clientRepository)
	clientRouter := r.Group("/clients")
	{
		clientRouter.GET("/", clientController.GetAllClients)      // Index
		clientRouter.POST("/", clientController.CreateClient)      // Create
		clientRouter.GET("/:id", clientController.GetClientById)   // Read
		clientRouter.PUT("/:id", clientController.UpdateClient)    // Update
		clientRouter.DELETE("/:id", clientController.DeleteClient) // Delete

		// Client Project Endpoints - https://domain/clients/:id/projects
		clientProjectRouter := clientRouter.Group("/:id/projects")
		{
			clientProjectRouter.GET("/", projectController.GetAllClientProjects) // Index
			clientProjectRouter.POST("/", projectController.CreateProject)       // Create
		}
	}

}
