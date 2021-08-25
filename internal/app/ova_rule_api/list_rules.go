package ova_rule_api

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rs/zerolog/log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *APIServer) ListRules(ctx context.Context, req *empty.Empty) (*emptypb.Empty, error) {
	log.Info().Msg(fmt.Sprintf("ListRulesRequest: %+v", req))
	return &emptypb.Empty{}, nil
}
