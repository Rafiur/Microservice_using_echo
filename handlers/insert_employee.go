package handlers

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // type Employee struct {
// // 	Id    int64  `json:"id"`
// // 	Name  string `json:"name"`
// // 	Email string `json:"email"`
// // }

// func insertEmployee(c echo.Context) error {
// 	db := DBconnection()
// 	defer db.Close() // Close database connection when function exits

// 	fmt.Println("Endpoint Hit: insert Employees")

// 	// Bind request body to Employee struct

// 	qry := `INSERT INTO public.information (name, email) VALUES($1, $2) returning id`

// 	var id int64

// 	row := db.QueryRowContext(c.Request().Context(), qry, emp.Name, emp.Email).Scan(&id)
// 	// if err != nil {
// 	// 	log.Println("Error inserting employee:", err)
// 	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 	// }

// 	// emp.Id = id
// 	fmt.Println(row)

// 	fmt.Println("Employee inserted successfully")
// 	return c.JSON(http.StatusOK, emp)
// }
