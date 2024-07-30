package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"usrmanagement/controllers"
	"usrmanagement/middlewares"
)

// SetupRouter sets up all routes
func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Include Authorization header
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)

	authorized := router.Group("/api")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/pages", controllers.GetUserPages)
		authorized.GET("/register", controllers.GetUsers)
		authorized.GET("/roles", controllers.GetRoles)
		authorized.POST("/roles", controllers.CreateRole)
		authorized.POST("/pages", controllers.CreatePage)
		authorized.POST("/roles/assign_pages", controllers.AssignPagesToRole)
		/*authorized.GET("/customers", controllers.GetCustomers)
		authorized.GET("/customers/:id", controllers.GetCustomer)
		authorized.POST("/customers", controllers.CreateCustomer)
		authorized.PUT("/customers/:id", controllers.UpdateCustomer)
		authorized.DELETE("/customers/:id", controllers.DeleteCustomer)
		authorized.POST("/loans", controllers.CreateLoan)
		authorized.POST("/collaterals", controllers.CreateCollateral)
		authorized.POST("/guarantors", controllers.CreateGuarantor)
		authorized.POST("/payments", controllers.CreatePayment)*/
	}

	return router
}
