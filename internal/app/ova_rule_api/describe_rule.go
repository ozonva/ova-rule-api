package ova_rule_api

import (
	"context"
	"github.com/rs/zerolog/log"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *APIServer) DescribeRule(ctx context.Context, req *desc.DescribeRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msgf("DescribeRuleRequest: %+v", req)
	return &emptypb.Empty{}, nil
}