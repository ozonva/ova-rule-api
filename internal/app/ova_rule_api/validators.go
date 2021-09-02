package ova_rule_api

import (
	"github.com/pkg/errors"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

var notPositiveIdError = errors.New("id must be positive")

func validateCreateRuleRequest(req *desc.CreateRuleRequest) error {
	if req.Id == 0 {
		return errors.Wrap(notPositiveIdError, "rule")
	}
	if req.UserId == 0 {
		return errors.Wrap(notPositiveIdError, "user")
	}
	if req.Name == "" {
		return errors.New("rule name must be not empty")
	}

	return nil
}

func validateListRulesRequest(req *desc.ListRulesRequest) error {
	if req.Limit == 0 {
		return errors.New("limit must be positive")
	}

	return nil
}

func validateDescribeRuleRequest(req *desc.DescribeRuleRequest) error {
	if req.Id == 0 {
		return errors.Wrap(notPositiveIdError, "rule")
	}

	return nil
}

func validateRemoveRuleRequest(req *desc.RemoveRuleRequest) error {
	if req.Id == 0 {
		return errors.Wrap(notPositiveIdError, "rule")
	}

	return nil
}

func validateMultiCreateRuleRequest(req *desc.MultiCreateRuleRequest) error {
	// TODO: имплементировать

	return nil
}
