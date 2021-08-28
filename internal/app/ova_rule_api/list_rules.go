package ova_rule_api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *apiServer) ListRules(ctx context.Context, req *empty.Empty) (*emptypb.Empty, error) {
	log.Info().Msgf("ListRulesRequest: %+v", req)
	return &emptypb.Empty{}, nil
}
