package ova_rule_api

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

func (a *apiServer) DescribeRule(
	ctx context.Context,
	req *desc.DescribeRuleRequest,
) (*desc.DescribeRuleResponse, error) {
	log.Info().Msgf("DescribeRuleRequest: %+v", req)

	if err := validateDescribeRuleRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	rule, err := a.repo.DescribeRule(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	result := desc.Rule{
		Id:     rule.ID,
		Name:   rule.Name,
		UserId: rule.UserID,
	}

	return &desc.DescribeRuleResponse{Result: &result}, nil
}
