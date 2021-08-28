package ova_rule_api

import (
	"fmt"
	"log"
	"net"

	"github.com/ozonva/ova-rule-api/configs"
	"github.com/ozonva/ova-rule-api/internal/repo"
	"github.com/ozonva/ova-rule-api/internal/saver"
	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"google.golang.org/grpc"
)

type apiServer struct {
	desc.UnimplementedAPIServer
	repo  repo.Repo
	saver saver.Saver
}

func NewAPIServer(repo repo.Repo, saver saver.Saver) desc.APIServer {
	return &apiServer{
		repo:  repo,
		saver: saver,
	}
}

func Run(apiServer *desc.APIServer) error {
	address := fmt.Sprintf("%s:%s", configs.ServerConfig.Host, configs.ServerConfig.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterAPIServer(s, *apiServer)

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
