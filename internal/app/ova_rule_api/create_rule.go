package ova_rule_api

import (
	"context"
	"fmt"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *APIServer) CreateRule(ctx context.Context, req *desc.CreateRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msg(fmt.Sprintf("CreateRuleRequest: %+v", req))
	return &emptypb.Empty{}, nil
}
