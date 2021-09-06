package ova_rule_api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

// Status возвращает статус приложения (healthcheck).
func (a *apiServer) Status(ctx context.Context, req *emptypb.Empty) (*desc.StatusResponse, error) {
	return &desc.StatusResponse{Status: "ok"}, nil
}
