package ova_rule_api

import (
	"context"
	"github.com/ozonva/ova-rule-api/internal/models"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

func (a *apiServer) MultiCreateRule(ctx context.Context, req *desc.MultiCreateRuleRequest) (*empty.Empty, error) {
	log.Info().Msgf("CreateRuleRequest: %+v", req)

	if err := validateMultiCreateRuleRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	for _, rule := range req.Rules {
		err := a.saver.Save(models.Rule{
			ID:     rule.Id,
			Name:   rule.Name,
			UserID: rule.UserId,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}
