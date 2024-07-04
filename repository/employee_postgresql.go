package repository

import (
	"context"
	"database/sql"
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

	_, err := repo.db.ExecContext(ctx, qryUpdate, employee.Name, employee.Email, id)
	if err != nil {
		log.Println("Error updating employee:", err)
		return response, err
	}

	// idInt, _ := strconv.Atoi(id)

	// response.Id = int64(idInt)

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
