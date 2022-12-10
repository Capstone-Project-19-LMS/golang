package routes

import (
	middlewares "golang/app/middlewares"
	middlewareCostumer "golang/app/middlewares/costumer"
	middlewareInstructor "golang/app/middlewares/instructor"
	assignmentcontroller "golang/controllers/assignmentController"
	"golang/controllers/categoryController"
	"golang/controllers/costumerController"
	"golang/controllers/courseController"
	customerassignmentcontroller "golang/controllers/customerAssignmentController"
	instructorController "golang/controllers/instructorController"
	mediamodulecontroller "golang/controllers/mediaModuleController"
	"golang/controllers/moduleController"
	"golang/helper"
	assignmentrepository "golang/repository/assignmentRepository"
	"golang/repository/categoryRepository"
	"golang/repository/courseRepository"
	customerassignmentrepository "golang/repository/customerAssignmentRepository"
	"golang/repository/customerRepository"
	instructorrepository "golang/repository/instructorRepository"
	mediamodulerepository "golang/repository/mediaModuleRepository"
	modulerepository "golang/repository/moduleRepository"
	assignmentservice "golang/service/assignmentService"
	"golang/service/categoryService"
	"golang/service/costumerService"
	"golang/service/courseService"
	"golang/service/customerAssignmentService"
	instructorservice "golang/service/instructorService"
	mediamoduleservice "golang/service/mediaModuleService"
	moduleservice "golang/service/moduleService"
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
	categoryRepository := categoryRepository.NewCategoryRepository(db)
	courseRepository := courseRepository.NewCourseRepository(db)
	moduleRepository := modulerepository.NewModuleRepository(db)
	mediamodulerepository := mediamodulerepository.NewMediaModuleRepository(db)
	assignmentRepository := assignmentrepository.NewAssignmentRepository(db)
	customerAssignmentRepository := customerassignmentrepository.NewcustomerAssignmentRepository(db)
	/*
		Services
	*/
	// customer
	costumerService := costumerService.NewcostumerService(customerRepository)

	// instructor
	instructorService := instructorservice.NewinstructorService(instructorRepository)
	categoryService := categoryService.NewCategoryService(categoryRepository)
	courseService := courseService.NewCourseService(courseRepository, categoryRepository)
	moduleService := moduleservice.NewModuleService(moduleRepository)
	mediamoduleservice := mediamoduleservice.NewMediaModuleService(mediamodulerepository)
	assignmentService := assignmentservice.NewAssignmentService(assignmentRepository)
	customerAssignmentService := customerAssignmentService.NewcustomerAssignmentService(customerAssignmentRepository)

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
	categoryController := categoryController.CategoryController{
		CategoryService: categoryService,
	}
	courseController := courseController.CourseController{
		CourseService: courseService,
	}

	moduleController := moduleController.ModuleController{
		ModuleService: moduleService,
	}

	mediaModuleController := mediamodulecontroller.MediaModuleController{
		MediaModuleService: mediamoduleservice,
	}

	assignmentController := assignmentcontroller.AssignmentController{
		AssignmentService: assignmentService,
	}

	customerAssignmentController := customerassignmentcontroller.CustomerAssignmentController{
		CustomerAssignmentService: customerAssignmentService,
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
	costumer := app.Group("/customer")
	costumer.POST("/register", costumerController.Register)
	costumer.POST("/login", costumerController.Login)

	privateCostumer := app.Group("/customer", middleware.JWTWithConfig(configCostumer))
	privateCostumer.Use(middlewares.CheckTokenMiddlewareCustomer)
	// private costumer access
	privateCostumer.POST("/logout", costumerController.Logout)

	// -->

	// instructor
	instructor := app.Group("/instructor")
	instructor.POST("/register", instructorController.Register)
	instructor.POST("/login", instructorController.Login)

	privateInstructor := app.Group("/instructor", middleware.JWTWithConfig(configInstructor))
	privateInstructor.Use(middlewares.CheckTokenMiddlewareInstructor)
	/*
		private instructor access
	*/
	privateInstructor.POST("/logout", instructorController.Logout)

	// category

	//instructor access
	privateInstructor.POST("/category/create", categoryController.CreateCategory)
	privateInstructor.DELETE("/category/delete/:id", categoryController.DeleteCategory)
	privateInstructor.GET("/category/get_all", categoryController.GetAllCategory)
	privateInstructor.GET("/category/get_by_id/:id", categoryController.GetCategoryByID)
	privateInstructor.PUT("/category/update/:id", categoryController.UpdateCategory)
	//costumer access
	privateCostumer.GET("/category/get_all", categoryController.GetAllCategory)
	privateCostumer.GET("/category/get_by_id/:id", categoryController.GetCategoryByID)

	// course

	//instructor access
	privateInstructor.POST("/course/create", courseController.CreateCourse)
	privateInstructor.DELETE("/course/delete/:id", courseController.DeleteCourse)
	privateInstructor.GET("/course/get_by_id/:id", courseController.GetCourseByID)
	privateInstructor.GET("/course/get_all", courseController.GetAllCourse)
	privateInstructor.PUT("/course/update/:id", courseController.UpdateCourse)
	//costumer access
	privateCostumer.GET("/course/get_by_id/:id", courseController.GetCourseByID)
	privateCostumer.GET("/course/get_all", courseController.GetAllCourse)

	//module
	//instructor access
	privateInstructor.POST("/module/create", moduleController.CreateModule)
	privateInstructor.DELETE("/module/delete/:id", moduleController.DeleteModule)
	privateInstructor.GET("/module/get_all", moduleController.GetAllModule)
	privateInstructor.GET("/module/get_by_id/:id", moduleController.GetModuleByID)
	privateInstructor.GET("/module/get_by_course_id/:course_id", moduleController.GetModuleByCourseID)
	privateInstructor.PUT("/module/update/:id", moduleController.UpdateModule)
	//costumer access
	privateCostumer.GET("/module/get_all", moduleController.GetAllModule)
	privateCostumer.GET("/module/get_by_id/:id", moduleController.GetModuleByID)
	privateCostumer.GET("/module/get_by_course_id/:course_id", moduleController.GetModuleByCourseID)

	//media module
	//instructor access
	privateInstructor.POST("/media_module/create", mediaModuleController.CreateMediaModule)
	privateInstructor.DELETE("/media_module/delete/:id", mediaModuleController.DeleteMediaModule)
	privateInstructor.GET("/media_module/get_all", mediaModuleController.GetAllMediaModule)
	privateInstructor.GET("/media_module/get_by_id/:id", mediaModuleController.GetMediaModuleByID)
	privateInstructor.PUT("/media_module/update/:id", mediaModuleController.UpdateMediaModule)
	//costumer access
	privateCostumer.GET("/media_module/get_all", mediaModuleController.GetAllMediaModule)
	privateCostumer.GET("/media_module/get_by_id/:id", mediaModuleController.GetMediaModuleByID)

	//assignment
	//instructor access
	privateInstructor.POST("/assignment/create", assignmentController.CreateAssignment)
	privateInstructor.DELETE("/assignment/delete/:id", assignmentController.DeleteAssignment)
	privateInstructor.GET("/assignment/get_all", assignmentController.GetAllAssignment)
	privateInstructor.GET("/assignment/get_by_id/:id", assignmentController.GetAssignmentByID)
	privateInstructor.PUT("/assignment/update/:id", assignmentController.UpdateAssignment)
	//costumer access
	privateCostumer.GET("/assignment/get_all", assignmentController.GetAllAssignment)
	privateCostumer.GET("/assignment/get_by_id/:id", assignmentController.GetAssignmentByID)

	//customer assignment
	//instructor access
	privateInstructor.POST("/customer_assignment/create", customerAssignmentController.CreateCustomerAssignment)
	privateInstructor.DELETE("/customer_assignment/delete/:id", customerAssignmentController.DeleteCustomerAssignment)
	privateInstructor.GET("/customer_assignment/get_all", customerAssignmentController.GetAllCustomerAssignment)
	privateInstructor.GET("/customer_assignment/get_by_id/:id", customerAssignmentController.GetCustomerAssignmentByID)
	privateInstructor.PUT("/customer_assignment/update/:id", customerAssignmentController.UpdateCustomerAssignment)
	//costumer access
	privateCostumer.POST("/customer_assignment/create", customerAssignmentController.CreateCustomerAssignment)
	privateCostumer.DELETE("/customer_assignment/delete/:id", customerAssignmentController.DeleteCustomerAssignment)
	privateCostumer.GET("/customer_assignment/get_all", customerAssignmentController.GetAllCustomerAssignment)
	privateCostumer.GET("/customer_assignment/get_by_id/:id", customerAssignmentController.GetCustomerAssignmentByID)
	privateCostumer.PUT("/customer_assignment/update/:id", customerAssignmentController.UpdateCustomerAssignment)

	return app
}
