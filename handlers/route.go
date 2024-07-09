package handlers

import (
	"database/sql"

	"example.com/PROJECT_1/entity/proto"
	"example.com/PROJECT_1/repository"
	"example.com/PROJECT_1/service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func SetUpRouter(e *echo.Echo, db *sql.DB, conn *grpc.ClientConn) {

	//e.GET("/", allEmployees)

	// e.POST("/add", insertEmployee)

	// e.PUT("/update/:id", updateEmployee)

	// e.DELETE("/delete/:id", deleteEmployee)

	grpcClient := proto.NewEmployeeToSalaryClient(conn)
	grpcRepo := repository.NewGrcpClientRepository(grpcClient)

	employeeRepo := repository.NewEmployeeRepo(db)
	employeeService := service.NewEmployeeService(employeeRepo, grpcRepo)
	employeeHandler := NewEmployeeHandler(employeeService)

	employeeGroup := e.Group("/employee")

	employeeHandler.MapEmployeeRoutes(employeeGroup)

}
