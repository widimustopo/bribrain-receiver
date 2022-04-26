package main

import (
	"bribrain-receiver/repository"
	"bribrain-receiver/rpc"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Printf("Error while get environment file %v", err)
	}

	//CONNECT TO MYSQL DB
	connDB := new(repository.ConnectionSQL)
	repository.ConnectionDB = connDB.InitializationDB(5)

	forever := make(chan bool)

	go func() {
		// RUNNING GRPC SERVER
		rpc := new(rpc.QoinRpc)
		rpc.GrpcStartServer()
	}()

	<-forever
}
