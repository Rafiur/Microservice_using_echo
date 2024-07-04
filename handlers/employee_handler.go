package handlers

import (
	"fmt"
	"log"
	"net/http"

	"example.com/PROJECT_1/entity"
	"example.com/PROJECT_1/service"
	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	EmployeeService *service.EmployeeService
}

func NewEmployeeHandler(employeeService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		EmployeeService: employeeService,
	}
}

func (h *EmployeeHandler) MapEmployeeRoutes(employeeGroup *echo.Group) {
	employeeGroup.GET("/all", h.GetAllEmployee)
	employeeGroup.POST("/insert", h.InsertEmployee)
	employeeGroup.PUT("/update/:id", h.UpdateEmployee)
	employeeGroup.DELETE("/delete/:id", h.DeleteEmployee)
}

func (h *EmployeeHandler) GetAllEmployee(c echo.Context) error {
	res, err := h.EmployeeService.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, res)
}

func (h *EmployeeHandler) InsertEmployee(c echo.Context) error {
	var payload entity.CreateEmployee

	if err := c.Bind(&payload); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	res, err := h.EmployeeService.Insert(c.Request().Context(), payload)
	fmt.Println(err)
	return c.JSON(http.StatusAccepted, res)
}

func (h *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	var payload entity.CreateEmployee

	id := c.Param("id")

	if err := c.Bind(payload); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	res, err := h.EmployeeService.Update(c.Request().Context(), payload, id)

	fmt.Println(res)
	return err
}

func (h *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	var payload entity.Employee

	id := c.Param("id")

	err := h.EmployeeService.Delete(c.Request().Context(), id, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return err
}
