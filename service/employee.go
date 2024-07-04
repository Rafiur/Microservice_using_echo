package service

import (
	"context"

	"example.com/PROJECT_1/entity"
)

type EmployeeService struct {
	repository Repository
}

func NewEmployeeService(repository Repository) *EmployeeService {
	return &EmployeeService{
		repository: repository,
	}
}

type Repository interface {
	Insert(ctx context.Context, employee entity.CreateEmployee) (entity.Employee, error)
	GetAll(ctx context.Context) ([]entity.Employee, error)
	Update(ctx context.Context, id string, employee entity.CreateEmployee) (entity.Employee, error)
	Delete(ctx context.Context, id string, employee entity.Employee) error
}

func (s *EmployeeService) GetAll(ctx context.Context) ([]entity.Employee, error) {
	employees, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *EmployeeService) Insert(ctx context.Context, employee entity.CreateEmployee) (entity.Employee, error) {

	res, err := s.repository.Insert(ctx, employee)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *EmployeeService) Update(ctx context.Context, employee entity.CreateEmployee, id string) (entity.Employee, error) {

	res, err := s.repository.Update(ctx, id, employee)
	if err != nil {
		return res, nil
	}
	return res, nil
}

func (s *EmployeeService) Delete(ctx context.Context, id string, employee entity.Employee) error {
	err := s.repository.Delete(ctx, id, employee)
	if err != nil {
		return err
	}
	return nil
}
