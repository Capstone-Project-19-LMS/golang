package routes

import (
	middlewares "golang/app/middlewares"
	middlewareCostumer "golang/app/middlewares/costumer"
	middlewareInstructor "golang/app/middlewares/instructor"
	"golang/controllers/costumerController"
	instructorController "golang/controllers/instructorController"
	"golang/helper"
	"golang/repository/customerRepository"
	instructorrepository "golang/repository/instructorRepository"
	"golang/service/costumerService"
	instructorservice "golang/service/instructorService"
	"golang/util"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	/*
		Repositories
	*/ 
	// customer
	customerRepository := customerRepository.NewCustomerRepository(db)
	// instructor
	instructorRepository := instructorrepository.Newinstructorrepository(db)
	
	/*
		Services
	*/ 
	// customer
	costumerService := costumerService.NewcostumerService(customerRepository)
	// instructor
	instructorService := instructorservice.NewinstructorService(instructorRepository)
	
	/*
	Controllers
	*/ 
	// customer
	costumerController := costumerController.CostumerController{
		CostumerService: costumerService,
	}
	// instructor
	instructorController := instructorController.InstructorController{
		InstructorService: instructorService,
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
	configCostumer := middleware.JWTConfig{
		Claims:     &middlewareCostumer.JwtCostumerClaims{},
		SigningKey: []byte(util.GetConfig("TOKEN_SECRET")),
	}
	configInstructor := middleware.JWTConfig{
		Claims:     &middlewareInstructor.JwtInstructorClaims{},
		SigningKey: []byte(util.GetConfig("TOKEN_SECRET")),
	}

	app.Use(configLogger.Init())

	// costumer
	costumer := app.Group("/custumer")
	costumer.POST("/register", costumerController.Register)
	costumer.POST("/login", costumerController.Login)

	privateCostumer := app.Group("/custumer", middleware.JWTWithConfig(configCostumer))

	// private costumer access
	privateCostumer.POST("/logout", costumerController.Logout)

	// -->

	// instructor
	instructor := app.Group("/instructor")
	instructor.POST("/register", instructorController.Register)
	instructor.POST("/login", instructorController.Login)

	privateInstructor := app.Group("/instructor", middleware.JWTWithConfig(configInstructor))

	// private instructor access
	privateInstructor.POST("/logout", instructorController.Logout)

	// -->

	return app
}
