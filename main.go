package main

import (
	"fmt"

	"example.com/PROJECT_1/db"
	v1 "example.com/PROJECT_1/handlers"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Employees []Employee

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "abcd1234"
// 	dbname   = "Employee"
// )

// func DBconnection() *sql.DB {
// 	psqlInfo := fmt.Sprintf("host = %s port = %d user = %s "+"password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//defer db.Close()
// 	//fmt.Println(reflect.TypeOf(db))
// 	return db
// }

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

// func InsertIntoDatabase(db *sql.DB, payload Employee) {
// 	insertEmp := `insert into public.information ("name", "email") values('Smith', 'agent.smith@gmail.com')`
// 	_, e := db.Exec(insertEmp)
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func insertEmployee(c echo.Context) error {
// 	db := DBconnection()
// 	defer db.Close() // Close database connection when function exits

// 	fmt.Println("Endpoint Hit: insert Employees")

// 	// Bind request body to Employee struct
// 	emp := new(Employee)
// 	if err := c.Bind(emp); err != nil {
// 		log.Println("Error binding request body:", err)
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
// 	}

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

func main() {

	db, err := db.DBconnection()
	fmt.Println(err)
	e := echo.New()

	v1.SetUpRouter(e, db)
	//fmt.Println("insert successfull")

	e.Logger.Fatal(e.Start(":4000"))

}
