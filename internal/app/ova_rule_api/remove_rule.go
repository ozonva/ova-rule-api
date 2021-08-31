package ova_rule_api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *apiServer) RemoveRule(ctx context.Context, req *desc.RemoveRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msgf("RemoveRuleRequest: %+v", req)

	err := a.repo.RemoveRule(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	log.Info().Msgf("Правило с id=%d удалено", req.Id)

	return &emptypb.Empty{}, nil
}
