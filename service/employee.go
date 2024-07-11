package service

import (
	"context"
	"fmt"
	"log"

	"example.com/PROJECT_1/entity"
	"example.com/PROJECT_1/entity/proto"
)

type EmployeeService struct {
	repository Repository
	grpcRepo   GrpcRepo
}

func NewEmployeeService(repository Repository, grpcRepo GrpcRepo) *EmployeeService {
	return &EmployeeService{
		repository: repository,
		grpcRepo:   grpcRepo,
	}
}

type Repository interface {
	Insert(ctx context.Context, employee entity.CreateEmployee) (entity.Employee, error)
	AddWithDeptRepo(ctx context.Context, employee_dept entity.CreateEmployeeWithDepartment) (entity.ResponseCreatewithDept, error)
	GetAll(ctx context.Context) ([]entity.Employee, error)
	Update(ctx context.Context, id string, employee entity.CreateEmployee) (entity.Employee, error)
	UpdateFull(ctx context.Context, employee entity.CreateEmployee, id string) (entity.UpdateEmployee, error)
	Delete(ctx context.Context, id string, employee entity.Employee) error
}

type GrpcRepo interface {
	CreateSalary(ctx context.Context, in *proto.CreateSalaryRequest) (*proto.CreateSalaryResponse, error)
	GetAllSalary(ctx context.Context, in *proto.GetAllSalaryRequest) ([]*proto.EmployeeSalary, error)
	UpdateSalary(ctx context.Context, in *proto.UpdateSalaryRequest) (*proto.UpdateSalaryResponse, error)
}

func (s *EmployeeService) Insert(ctx context.Context, employee entity.CreateEmployee) (entity.Employee, error) {

	res, err := s.repository.Insert(ctx, employee)
	fmt.Println(err)

	payload := &proto.CreateSalaryRequest{
		EmployeeId:   int32(res.Id),
		SalaryAmount: int32(employee.SalaryEvaluation.Salary_amount),
		Project:      employee.SalaryEvaluation.Project,
		JoiningDate:  employee.SalaryEvaluation.Joining_Date,
	}

	_, err = s.grpcRepo.CreateSalary(ctx, payload)

	fmt.Println("err.........", err)

	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *EmployeeService) GetAll(ctx context.Context) ([]entity.Employee, error) {
	employees, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *EmployeeService) GetAllWithSalary(ctx context.Context) ([]entity.GetAllEmployee, error) {
	employee_salaries, err := s.grpcRepo.GetAllSalary(ctx, &proto.GetAllSalaryRequest{})
	if err != nil {
		return nil, err
	}

	var employees []entity.GetAllEmployee
	for _, salary := range employee_salaries {
		employees = append(employees, entity.GetAllEmployee{
			Id:            int64(salary.EmployeeId),
			Name:          salary.Name,
			Email:         salary.Email,
			Joining_Date:  salary.JoiningDate,
			Project:       salary.Project,
			Salary_amount: int(salary.SalaryAmount),
		})
	}
	fmt.Println(employees)
	return employees, nil
}
//same databse different table insertion
func (s *EmployeeService) AddWithDeptService(ctx context.Context, employee_dept entity.CreateEmployeeWithDepartment) (entity.ResponseCreatewithDept, error) {
	res, err := s.repository.AddWithDeptRepo(ctx, employee_dept)
	if err != nil {
		log.Println("Service function failure:", err)
	}
	fmt.Println("Service res:", res)
	return res, nil
}

func (s *EmployeeService) Update(ctx context.Context, employee entity.CreateEmployee, id string) (entity.Employee, error) {

	res, err := s.repository.Update(ctx, id, employee)
	if err != nil {
		return res, nil
	}

	return res, nil
}

func (s *EmployeeService) UpdateFull(ctx context.Context, employee entity.CreateEmployee, id string) (*proto.UpdateSalaryResponse, error) {

	res, err := s.repository.UpdateFull(ctx, employee, id)

	fmt.Println(err)

	payload := &proto.UpdateSalaryRequest{
		EmployeeId:   int32(res.Id),
		SalaryAmount: int32(employee.SalaryEvaluation.Salary_amount),
		Project:      employee.SalaryEvaluation.Project,
		JoiningDate:  employee.SalaryEvaluation.Joining_Date,
	}

	resp, err := s.grpcRepo.UpdateSalary(ctx, payload)

	//fmt.Println(resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *EmployeeService) Delete(ctx context.Context, id string, employee entity.Employee) error {
	err := s.repository.Delete(ctx, id, employee)
	if err != nil {
		return err
	}
	return nil
}
