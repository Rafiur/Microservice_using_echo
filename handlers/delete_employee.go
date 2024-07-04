package handlers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func deleteEmployee(c echo.Context) error {
// 	db := DBconnection()
// 	defer db.Close() // Close database connection when function exits

// 	fmt.Println("Endpoint Hit: delete Employees")

// 	id := c.Param("id")

// 	qry := `DELETE from public.information WHERE id=$1`

// 	_, err := db.ExecContext(c.Request().Context(), qry, id)
// 	if err != nil {
// 		log.Println("Error deleting employee:", err)
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 	}

// 	fmt.Println("Employee deleted successfully")
// 	return c.JSON(http.StatusOK, "Employee deleted successfully")

// }
