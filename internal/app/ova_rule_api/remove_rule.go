package ova_rule_api

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

// RemoveRule удаляет правило из БД по идентификатору.
func (a *apiServer) RemoveRule(ctx context.Context, req *desc.RemoveRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msgf("RemoveRuleRequest: %+v", req)

	if err := validateRemoveRuleRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	err := a.repo.RemoveRule(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	a.metrics.RemoveRuleCounterInc()

	log.Info().Msgf("Правило с id=%d удалено", req.Id)

	return &emptypb.Empty{}, nil
}
