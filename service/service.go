package service

import (
	"bribrain-receiver/repository"
	"bribrain-receiver/rpc/rka"
	"context"

	"github.com/sirupsen/logrus"
)

type RkaService struct {
	rka.UnimplementedRkaServiceServer
}

func (RkaService) Adduser(ctx context.Context, in *rka.User) (*rka.ResponseUser, error) {
	logrus.Info("Adduser data in : ", in)
	query := repository.Queries{}
	result, err := query.Adduser(in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (RkaService) AddRKA(ctx context.Context, in *rka.RKA) (*rka.ResponseRKA, error) {
	logrus.Info("AddRKA data in : ", in)
	query := repository.Queries{}
	result, err := query.AddRka(in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (RkaService) AddUserRKA(ctx context.Context, in *rka.RequestUserRKA) (*rka.ResponseUserRKA, error) {
	logrus.Info("AddUserRKA data in : ", in)
	query := repository.Queries{}
	result, err := query.AddRkaUser(in)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (RkaService) AccumulationRKA(ctx context.Context, in *rka.RequestAccumulation) (*rka.ResponseAccumulationRKA, error) {
	logrus.Info("AccumulationRKA data in : ", in)
	query := repository.Queries{}
	result, err := query.AccumulationRKA(in)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return result, nil
}
