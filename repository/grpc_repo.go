package repository

import (
	"context"
	"time"

	"example.com/PROJECT_1/entity/proto"
)

type GrpcClientRepository struct {
	client proto.EmployeeToSalaryClient
}

func NewGrcpClientRepository(client proto.EmployeeToSalaryClient) *GrpcClientRepository {
	return &GrpcClientRepository{
		client: client,
	}
}

func (repo *GrpcClientRepository) CreateSalary(ctx context.Context, in *proto.CreateSalaryRequest) (*proto.CreateSalaryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := repo.client.CreateSalary(ctx, in)
	if err != nil {
		return &proto.CreateSalaryResponse{}, err
	}
	return res, nil
}

func (repo *GrpcClientRepository) UpdateSalary(ctx context.Context, in *proto.UpdateSalaryRequest) (*proto.UpdateSalaryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := repo.client.UpdateSalary(ctx, in)
	if err != nil {
		return &proto.UpdateSalaryResponse{}, err
	}
	return res, nil
}
