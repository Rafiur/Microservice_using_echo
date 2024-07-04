package handlers

import (
	"database/sql"

	"example.com/PROJECT_1/repository"
	"example.com/PROJECT_1/service"
	"github.com/labstack/echo/v4"
)

func SetUpRouter(e *echo.Echo, db *sql.DB) {

	//e.GET("/", allEmployees)

	// e.POST("/add", insertEmployee)

	// e.PUT("/update/:id", updateEmployee)

	// e.DELETE("/delete/:id", deleteEmployee)


	employeeRepo := repository.NewEmployeeRepo(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := NewEmployeeHandler(employeeService)

	employeeGroup := e.Group("/employee")

	employeeHandler.MapEmployeeRoutes(employeeGroup)

}
