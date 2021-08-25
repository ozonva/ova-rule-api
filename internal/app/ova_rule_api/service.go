package ova_rule_api

import (
	"log"
	"net"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"google.golang.org/grpc"
	"github.com/ozonva/ova-rule-api/configs"
)

type APIServer struct {
	desc.UnimplementedAPIServer
}

func NewAPIServer() desc.APIServer {
	return &APIServer{}
}

func Run() error {
	listen, err := net.Listen("tcp", configs.ServerConfig.Port)
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
