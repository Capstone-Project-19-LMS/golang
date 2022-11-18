package routes

import (
	"golang/app/middlewares"
	"golang/controllers/userController"
	"golang/helper"
	"golang/repository/customerRepository"
	"golang/service/userService"
	"golang/util"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	// Repositories
	customerRepository := customerRepository.NewCustomerRepository(db)

	// Services
	userService := userService.NewUserService(customerRepository)

	// Controllers
	userController := userController.UserController{
		UserService: userService,
	}

	app := echo.New()

	app.Validator = &helper.CustomValidator{
		Validator: validator.New(),
	}
	
	
	/* 
	API Routes
	*/
	configLogger := middlewares.ConfigLogger{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}
	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte(util.GetConfig("TOKEN_SECRET")),
	}

	app.Use(configLogger.Init())

	app.POST("/register", userController.Register)
	app.POST("/login", userController.Login)

	auth := app.Group("/user", middleware.JWTWithConfig(config))

	auth.POST("/logout", userController.Logout)

	return app
}
