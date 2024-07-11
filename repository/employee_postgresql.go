package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"example.com/PROJECT_1/entity"
)

type EmployeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *EmployeeRepo {
	return &EmployeeRepo{
		db: db,
	}
}

func (repo *EmployeeRepo) Insert(ctx context.Context, employee entity.CreateEmployee) (entity.Employee, error) {

	var response entity.Employee

	qry := `INSERT INTO public.information (name, email) VALUES($1, $2) returning *`

	err := repo.db.QueryRowContext(ctx, qry, employee.Name, employee.Email).Scan(&response.Name, &response.Email, &response.Id)

	return response, err
}
//same databse different table insertion
func (repo *EmployeeRepo) AddWithDeptRepo(ctx context.Context, employee_dept entity.CreateEmployeeWithDepartment) (entity.ResponseCreatewithDept, error) {

	var response entity.ResponseCreatewithDept

	qry := `
		WITH ins AS(
			INSERT INTO public.information (name, email)
			VALUES ($1, $2)
			RETURNING id
		)
		INSERT INTO public.department (name, manager,employee_id)
		VALUES ($3, $4, (SELECT id FROM ins))
		RETURNING id, $1 AS employee_name, $2 AS employee_email, $3 AS department_name, $4 AS department_manager, (SELECT id FROM ins) AS employee_id
	`

	err := repo.db.QueryRowContext(ctx, qry, employee_dept.Employee_Name, employee_dept.Employee_Email, employee_dept.Department_Name, employee_dept.Department_Manager).
		Scan(&response.Department_Id, &response.Employee_Name, &response.Employee_Email, &response.Department_Name, &response.Department_Manager, &response.EmployeeId)

	fmt.Println("Repo response:", response)
	return response, err
}

func (repo *EmployeeRepo) GetAll(ctx context.Context) ([]entity.Employee, error) {
	qry := `SELECT id, name, email FROM public.information`

	rows, err := repo.db.QueryContext(ctx, qry)
	if err != nil {
		log.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	employees := []entity.Employee{} // Slice to hold all employees

	for rows.Next() {
		var id int64
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		employee := entity.Employee{Id: id, Name: name, Email: email}
		employees = append(employees, employee)
	}

	return employees, err
}

func (repo *EmployeeRepo) Update(ctx context.Context, id string, employee entity.CreateEmployee) (entity.Employee, error) {

	var response entity.Employee

	qryUpdate := `UPDATE public.information SET name=$1, email=$2 WHERE id=$3 RETURNING *`

	err := repo.db.QueryRowContext(ctx, qryUpdate, employee.Name, employee.Email, id).Scan(&response.Name, &response.Email, &response.Id)
	if err != nil {
		log.Println("Error updating employee:", err)
		return entity.Employee{}, err
	}

	return response, err
}

func (repo *EmployeeRepo) UpdateFull(ctx context.Context, employee entity.CreateEmployee, id string) (entity.UpdateEmployee, error) {
	var response entity.UpdateEmployee

	qryUpdate := `UPDATE public.information SET name=$1, email=$2 WHERE id=$3 RETURNING *`

	err := repo.db.QueryRowContext(ctx, qryUpdate, employee.Name, employee.Email, id).Scan(&response.Name, &response.Email, &response.Id)
	if err != nil {
		log.Println("Error updating employee:", err)
		return entity.UpdateEmployee{}, err
	}

	return response, err
}

func (repo *EmployeeRepo) Delete(ctx context.Context, id string, employee entity.Employee) error {

	qry := `DELETE from public.information WHERE id=$1`

	_, err := repo.db.ExecContext(ctx, qry, id)
	if err != nil {
		log.Println("Error deleting employee:", err)
		return nil
	}
	return nil
}
