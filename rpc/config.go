package rpc

import (
	"log"
	"net"
	"os"

	grpc "google.golang.org/grpc"
)

// var RpcConfig *QoinRpc

type IRpc interface {
	GrpcStartServer()
}

type QoinRpc struct {
	Host    string
	Port    string
	Rpc     *IRpc
	RServer *grpc.Server
}

func (c *QoinRpc) GrpcStartServer() {

	// INIT ENVIRONMENT
	c.StartConnection()

	listen, err := net.Listen("tcp", ":"+c.Port)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	c.RServer = grpc.NewServer()
	RouteGrpcServer(c.RServer)

	// START GRPC SERVER
	log.Println("starting gRPC server ...")
	log.Println("gRpc Port :", c.Port)
	log.Println("gRpc Host :", c.Host)
	log.Println("⇨ grpc server started on [::]:" + c.Port)
	if err := c.RServer.Serve(listen); err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
}

func (c *QoinRpc) StartConnection() {
	c.Host = os.Getenv("GRPC_HOST")
	c.Port = os.Getenv("GRPC_PORT")
}
