package handlers

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // type Employee struct {
// // 	Id    int64  `json:"id"`
// // 	Name  string `json:"name"`
// // 	Email string `json:"email"`
// // }

// type Employees []Employee

// func allEmployees(c echo.Context) error {
// 	db := DBconnection()
// 	defer db.Close() // Close database connection when function exits

// 	qry := `SELECT id, name, email FROM public.information`

// 	rows, err := db.QueryContext(c.Request().Context(), qry)
// 	if err != nil {
// 		log.Println("Error querying database:", err)
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 	}
// 	defer rows.Close()

// 	employees := Employees{} // Slice to hold all employees

// 	for rows.Next() {
// 		var id int64
// 		var name, email string
// 		err := rows.Scan(&id, &name, &email)
// 		if err != nil {
// 			log.Println("Error scanning row:", err)
// 			continue
// 		}
// 		employee := Employee{Id: id, Name: name, Email: email}
// 		employees = append(employees, employee)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Println("Error iterating rows:", err)
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
// 	}

// 	fmt.Println("Endpoint Hit: All Employees")
// 	return c.JSON(http.StatusOK, employees)
// }
