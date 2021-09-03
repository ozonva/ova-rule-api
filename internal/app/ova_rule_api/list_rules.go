package ova_rule_api

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

func (a *apiServer) ListRules(
	ctx context.Context,
	req *desc.ListRulesRequest,
) (*desc.ListRulesResponse, error) {
	log.Info().Msgf("ListRulesRequest: %+v", req)

	if err := validateListRulesRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

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
