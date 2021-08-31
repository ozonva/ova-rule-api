package ova_rule_api

import (
	"context"
	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"
)

func (a *apiServer) ListRules(
	ctx context.Context,
	req *desc.ListRulesRequest,
) (*desc.ListRulesResponse, error) {
	log.Info().Msgf("ListRulesRequest: %+v", req)

	rules, err := a.repo.ListRules(req.Limit, req.Offset)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	result := make([]*desc.Rule, 0, len(rules))
	for _, rule := range rules {
		result = append(result, &desc.Rule{
			Id:     rule.ID,
			Name:   rule.Name,
			UserId: rule.UserID,
		})
	}

	return &desc.ListRulesResponse{Result: result}, nil
}
