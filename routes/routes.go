package routes

import (
	contacthandler "api-rs/handlers/contact"
	examplehandler "api-rs/handlers/example"
	userhandler "api-rs/handlers/user"
	"api-rs/middlewares"
	contactrepository "api-rs/repositories/contact"
	userrepository "api-rs/repositories/user"
	contactservice "api-rs/services/contact"
	userservice "api-rs/services/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	var (
		// Repositories
		userRepository    userrepository.UserRepository       = userrepository.NewUserRepository(db)
		contactRepository contactrepository.ContactRepository = contactrepository.NewContactRepository(db)

		// Services
		userService    userservice.UserService       = userservice.NewUserService(userRepository)
		contactService contactservice.ContactService = contactservice.NewContactService(contactRepository)

		// Handlers
		exampleHandler examplehandler.ExampleHandler = examplehandler.NewExampleHandler()
		userHandler    userhandler.UserHandler       = userhandler.NewUserHandler(userService)
		contactHandler contacthandler.ContactHandler = contacthandler.NewContactHandler(contactService)

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

			clientRoutes.GET("/contact", contactHandler.GetContactClient)
		}

		// Admin Routes
		adminRoutes := apiRoutes.Group("admin")
		{
			// No Auth
			adminRoutes.POST("/login", userHandler.Login)

			// With Auth
			adminAuthRoutes := adminRoutes.Group("", authMiddleware)
			{
				adminAuthRoutes.GET("/users", userHandler.ListUser)

				// Contact Routes
				contactAdminRoutes := adminAuthRoutes.Group("contact")
				{
					contactAdminRoutes.GET("", contactHandler.GetContacts)
					contactAdminRoutes.POST("", contactHandler.CreateContact)
					contactAdminRoutes.GET("/:id", contactHandler.GetContact)
					contactAdminRoutes.PUT("/:id", contactHandler.UpdateContact)
					contactAdminRoutes.DELETE("/:id", contactHandler.DeleteContact)
				}
			}
		}
	}
}
