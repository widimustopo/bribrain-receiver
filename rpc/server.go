package rpc

import (
	"bribrain-receiver/rpc/rka"
	"bribrain-receiver/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var RoutesServer *grpc.Server

func init() {
	logrus.Info("Initialization rpc")
}

func RouteGrpcServer(s *grpc.Server) {
	rka.RegisterRkaServiceServer(s, service.RkaService{})
}
