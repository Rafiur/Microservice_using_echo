package entity

type Employee struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SalaryEvaluation struct {
	Salary_amount int    `json:"salary_amount"`
	Project       string `json:"project"`
	Joining_Date  string `json:"joining_date"`
}

type CreateEmployee struct {
	//Id    int64  `json:"id"`
	Name             string           `json:"name"`
	Email            string           `json:"email"`
	SalaryEvaluation SalaryEvaluation `json:"salary_evaluation"`
}

type GetAllEmployee struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Salary_amount int    `json:"salary_amount"`
	Project       string `json:"project"`
	Joining_Date  string `json:"joining_date"`
}
type UpdateEmployee struct {
	Id               int64            `json:"id"`
	Name             string           `json:"name"`
	Email            string           `json:"email"`
	SalaryEvaluation SalaryEvaluation `json:"salary_evaluation"`
}

type Department struct {
	Id         int    `json:"id"`
	Name       string `json:"department_name"`
	Manager    string `json:"department_manager"`
	EmployeeId int64  `json:"employeeId"`
}

type CreateEmployeeWithDepartment struct {
	Employee_Name      string `json:"employee_name"`
	Employee_Email     string `json:"employee_email"`
	Department_Name    string `json:"department_name"`
	Department_Manager string `json:"department_manager"`
	//EmployeeId         int64  `json:"employeeId"`
}

type ResponseCreatewithDept struct {
	Department_Id      int    `json:"department_id"`
	Employee_Name      string `json:"employee_name"`
	Employee_Email     string `json:"employee_email"`
	Department_Name    string `json:"department_name"`
	Department_Manager string `json:"department_manager"`
	EmployeeId         int64  `json:"employeeId"`
}
