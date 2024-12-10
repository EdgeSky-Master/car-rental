package routes

import (
	"car-rental/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r:= gin.Default()
	carRoutes := r.Group("/cars")
	{
		carRoutes.GET("/", controllers.GetCars)
		carRoutes.POST("/", controllers.CreateCar)
		carRoutes.PUT("/:id", controllers.UpdateCar)
		carRoutes.DELETE("/:id", controllers.DeleteCar)
	}
	
	bookingRoutes := r.Group("/bookings")
	{
		bookingRoutes.GET("/", controllers.GetBookings)
		bookingRoutes.POST("/", controllers.CreateBooking)
		bookingRoutes.PUT("/:id", controllers.UpdateBooking)
		bookingRoutes.DELETE("/:id", controllers.DeleteBooking)
	}
	
	customerRoutes := r.Group("/customers")
	{
		customerRoutes.GET("/", controllers.GetCustomers)
		customerRoutes.POST("/", controllers.CreateCustomer)
		customerRoutes.PUT("/:id", controllers.UpdateCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomer)
	}

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)

	}
return r
}