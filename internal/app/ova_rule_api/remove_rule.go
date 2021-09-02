package ova_rule_api

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-rule-api/internal/kafka"
	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

func (a *apiServer) RemoveRule(ctx context.Context, req *desc.RemoveRuleRequest) (*emptypb.Empty, error) {
	log.Info().Msgf("RemoveRuleRequest: %+v", req)

	if err := validateRemoveRuleRequest(req); err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	err := a.repo.RemoveRule(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	log.Info().Msgf("Правило с id=%d удалено", req.Id)

	msg := encodeRemoveRuleRequestToJSON(req)
	preparedMsg := kafka.PrepareMessage("remove_rule", msg)
	a.producer.SendMessageWithContext(ctx, preparedMsg)

	log.Info().Msgf("Отправили в очередь событие про удаление правила с id=%d", req.Id)

	return &emptypb.Empty{}, nil
}

func encodeRemoveRuleRequestToJSON(req *desc.RemoveRuleRequest) string {
	body := struct {
		ID uint64 `json:"id"`
	}{
		ID: req.Id,
	}

	result, _ := json.Marshal(body)

	return string(result)
}
