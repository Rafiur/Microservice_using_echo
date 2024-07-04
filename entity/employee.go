package entity

type Employee struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateEmployee struct {
	//Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
