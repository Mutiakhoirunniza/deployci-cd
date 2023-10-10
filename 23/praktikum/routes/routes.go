package routes

import (
	"20/controller"
	"20/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	// proses pembuatan object dan inject
	UserRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(*&UserRepository)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] status=${status} method=${method} uri=${uri} latency=${latency_human} \n",
	}))

	//endpoint
	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
}
