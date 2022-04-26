package service

import (
	"bribrain-receiver/rpc/rka"
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type RkaService struct {
	rka.UnimplementedRkaServiceServer
}

func (RkaService) Adduser(ctx context.Context, in *rka.User) (*rka.ResponseUser, error) {
	logrus.Info("Adduser data in : ", in)
	return &rka.ResponseUser{
		Code:    http.StatusOK,
		Message: "success add user",
		Data:    nil,
	}, nil
}

func (RkaService) AddRKA(ctx context.Context, in *rka.RKA) (*rka.ResponseRKA, error) {
	logrus.Info("AddRKA data in : ", in)
	return &rka.ResponseRKA{
		Code:    http.StatusOK,
		Message: "success add RKA",
		Data:    nil,
	}, nil
}

func (RkaService) AddUserRKA(ctx context.Context, in *rka.RequestUserRKA) (*rka.ResponseUserRKA, error) {
	logrus.Info("AddUserRKA data in : ", in)
	return &rka.ResponseUserRKA{
		Code:    http.StatusOK,
		Message: "success add user RKA",
		Data:    nil,
	}, nil
}

func (RkaService) AccumulationRKA(ctx context.Context, in *rka.RequestAccumulation) (*rka.ResponseAccumulationRKA, error) {
	logrus.Info("data in : ", in)
	return &rka.ResponseAccumulationRKA{
		Code:            http.StatusOK,
		Message:         "success accumulation data",
		ProgressVisited: nil,
	}, nil
}
