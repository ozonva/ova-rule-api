package ova_rule_api

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-rule-api/internal/models"
	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

func (a *apiServer) CreateRule(ctx context.Context, req *desc.CreateRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msgf("CreateRuleRequest: %+v", req)

	if err := validateCreateRuleRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

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
