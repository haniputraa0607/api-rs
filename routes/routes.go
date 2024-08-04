package routes

import (
	examplehandler "api-rs/handlers/example"
	userhandler "api-rs/handlers/user"
	"api-rs/middlewares"
	userrepository "api-rs/repositories/user"
	userservice "api-rs/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	var (
		// Repositories
		userRepository userrepository.UserRepository = userrepository.NewUserRepository(db)

		// Services
		userService userservice.UserService = userservice.NewUserService(userRepository)

		// Handlers
		userHandler    userhandler.UserHandler       = userhandler.NewUserHandler(userService)
		exampleHandler examplehandler.ExampleHandler = examplehandler.NewExampleHandler()

		// Middlewares
		authMiddleware = middlewares.AuthMiddleware()
	)

	apiRoutes := r.Group("api")
	{
		// Example Routes
		exampleRoutes := apiRoutes.Group("example")
		{
			exampleRoutes.GET("/test", exampleHandler.Example)
		}

		// Costumer Routes
		clientRoutes := apiRoutes.Group("client")
		{
			clientRoutes.GET("/test", exampleHandler.Example)
		}

		// Admin Routes
		adminRoutes := apiRoutes.Group("admin")
		{
			adminRoutes.POST("/login", userHandler.Login)

			adminAuthRoutes := adminRoutes.Group("", authMiddleware)
			{
				adminAuthRoutes.GET("/users", userHandler.ListUser)
			}
		}
	}
}
