package handlers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// // type Employee struct {
// // 	Id    int64  `json:"id"`
// // 	Name  string `json:"name"`
// // 	Email string `json:"email"`
// // }

// func updateEmployee(c echo.Context) error {
// 	db := DBconnection()
// 	defer db.Close() // Close database connection when function exits

// 	fmt.Println("Endpoint Hit: update Employees")

// 	id := c.Param("id")

// 	emp := new(Employee)

// 	if err := c.Bind(emp); err != nil {
// 		log.Println("Error binding request body:", err)
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
// 	}

// 	// Fetch current details of the employee
// 	// qry_1 := `SELECT id, name, email FROM public.information WHERE id=$1`

// 	// row := db.QueryRow(qry_1, id)
// 	// err := row.Scan(&emp.Id, &emp.Name, &emp.Email)
// 	// if err != nil {
// 	// 	log.Println("Error fetching employee details:", err)
// 	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 	// }

// 	// Update the employee details
// 	qryUpdate := `UPDATE public.information SET name=$1, email=$2 WHERE id=$3 RETURNING *`
// 	_, err := db.ExecContext(c.Request().Context(), qryUpdate, emp.Name, emp.Email, id)
// 	if err != nil {
// 		log.Println("Error updating employee:", err)
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 	}

// 	// Fetch updated details of the employee
// 	// qry_1 := `SELECT id, name, email FROM public.information WHERE id=$1`

// 	// row := db.QueryRow(qry_1, id)
// 	// err_1 := row.Scan(&emp.Id, &emp.Name, &emp.Email)
// 	// if err_1 != nil {
// 	// 	log.Println("Error fetching employee details:", err_1)
// 	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 	// }

// 	idInt, _ := strconv.Atoi(id)

// 	emp.Id = int64(idInt)

// 	fmt.Println("Employee updates successfully")
// 	return c.JSON(http.StatusOK, emp)

// }
