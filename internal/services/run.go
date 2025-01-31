package services

import (
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func RunService(service string, repo interface{}) {
	lsn, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		log.Fatalln("can't listen on the port", err)
	}

	server := grpc.NewServer()
	err = RegisterService(server, service, repo)
	if err != nil {
		log.Fatalf("error from `RegisterService` function, package `services`: %v", err)
	}

	err = server.Serve(lsn)
	if err != nil {
		log.Fatalln("can't accept incoming connections", err)
	}
}
