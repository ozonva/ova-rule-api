package ova_rule_api

import (
	"context"

	"github.com/ozonva/ova-rule-api/internal/models"
	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *apiServer) CreateRule(ctx context.Context, req *desc.CreateRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msgf("CreateRuleRequest: %+v", req)

	err := a.saver.Save(models.Rule{
		ID:     req.Id,
		Name:   req.Name,
		UserID: req.UserId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
