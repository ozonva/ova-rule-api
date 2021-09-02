package ova_rule_api

import (
	"context"
	"github.com/ozonva/ova-rule-api/internal/models"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

func (a *apiServer) UpdateRule(ctx context.Context, req *desc.UpdateRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msgf("UpdateRuleRequest: %+v", req)

	if err := validateUpdateRuleRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	rule := models.Rule{
		ID: req.Rule.Id,
		Name: req.Rule.Name,
		UserID: req.Rule.UserId,
	}
	err := a.repo.UpdateRule(rule)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	log.Info().Msgf("Правило с id=%d обновлено", req.Rule.Id)

	return &emptypb.Empty{}, nil
}
