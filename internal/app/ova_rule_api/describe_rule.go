package ova_rule_api

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

func (a *apiServer) DescribeRule(
	ctx context.Context,
	req *desc.DescribeRuleRequest,
) (*desc.DescribeRuleResponse, error) {
	log.Info().Msg(fmt.Sprintf("DescribeRuleRequest: %+v", req))

	rule, err := a.repo.DescribeRule(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	result := desc.Rule{
		Id: rule.ID,
		Name: rule.Name,
		UserId: rule.UserID,
	}

	return &desc.DescribeRuleResponse{Result: &result}, nil
}
