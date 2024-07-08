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
