package ova_rule_api

import (
	"fmt"
	"log"
	"net"

	"github.com/ozonva/ova-rule-api/configs"
	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"google.golang.org/grpc"
)

type APIServer struct {
	desc.UnimplementedAPIServer
}

func NewAPIServer() desc.APIServer {
	return &APIServer{}
}

func Run() error {
	address := fmt.Sprintf("%s:%s", configs.ServerConfig.Host, configs.ServerConfig.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterAPIServer(s, NewAPIServer())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
