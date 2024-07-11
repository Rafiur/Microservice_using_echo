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
	//api calls
	employeeGroup.GET("/all", h.GetAllEmployee)
	employeeGroup.POST("/insert", h.InsertEmployee) //called for both grpc and api
	employeeGroup.PUT("/update/:id", h.UpdateEmployee)
	employeeGroup.DELETE("/delete/:id", h.DeleteEmployee)
	//intra database insertion
	employeeGroup.POST("/addwdept", h.AddEmployeeWithDept)
	employeeGroup.GET("/allwdept", h.GetAllEmployeeWithDept)
	//grpc calls
	employeeGroup.GET("/allsalary", h.GetAllEmployeeWithSalary)
	employeeGroup.PUT("/update_full/:id", h.UpdateEmployeeWithSalary)
	
}

func (h *EmployeeHandler) GetAllEmployee(c echo.Context) error {
	res, err := h.EmployeeService.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, res)
}

func (h *EmployeeHandler) GetAllEmployeeWithSalary(c echo.Context) error {
	res, err := h.EmployeeService.GetAllWithSalary(c.Request().Context())
	if err != nil {
		log.Println("Handler function:", err)
	}
	fmt.Println(res)
	return c.JSON(http.StatusAccepted, res)
}

func (h *EmployeeHandler) GetAllEmployeeWithDept(c echo.Context) error{
	res, err := h.EmployeeService.GetAllWithDeptService(c.Request().Context())
	if err != nil{
		log.Println("Handler function:",err)
	}
	fmt.Println("Handler res:",res)
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

// same databse different table insertion
func (h *EmployeeHandler) AddEmployeeWithDept(c echo.Context) error {
	var payload entity.CreateEmployeeWithDepartment

	if err := c.Bind(&payload); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	res, err := h.EmployeeService.AddWithDeptService(c.Request().Context(), payload)
	fmt.Println(err)
	fmt.Println("Handler res:", res)
	return c.JSON(http.StatusAccepted, res)
}

func (h *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	var payload entity.CreateEmployee

	id := c.Param("id")

	if err := c.Bind(&payload); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

	res, err := h.EmployeeService.Update(c.Request().Context(), payload, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	return c.JSON(http.StatusAccepted, res)
}

func (h *EmployeeHandler) UpdateEmployeeWithSalary(c echo.Context) error {
	var payload entity.CreateEmployee

	id := c.Param("id")

	if err := c.Bind(&payload); err != nil {
		log.Println("Error binding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request binding"})
	}

	res, err := h.EmployeeService.UpdateFull(c.Request().Context(), payload, id)

	fmt.Println(err)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request function"})
	}
	return c.JSON(http.StatusAccepted, res)
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
