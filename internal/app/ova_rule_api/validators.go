package ova_rule_api

import (
	"github.com/pkg/errors"

	desc "github.com/ozonva/ova-rule-api/pkg/api/github.com/ozonva/ova-rule-api/pkg/ova-rule-api"
)

var errNotPositiveID = errors.New("id must be positive")

func validateCreateRuleRequest(req *desc.CreateRuleRequest) error {
	if req.Id == 0 {
		return errors.Wrap(errNotPositiveID, "rule")
	}

	if req.UserId == 0 {
		return errors.Wrap(errNotPositiveID, "user")
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
		return errors.Wrap(errNotPositiveID, "rule")
	}

	return nil
}

func validateRemoveRuleRequest(req *desc.RemoveRuleRequest) error {
	if req.Id == 0 {
		return errors.Wrap(errNotPositiveID, "rule")
	}

	return nil
}

func validateMultiCreateRuleRequest(req *desc.MultiCreateRuleRequest) error {
	for _, rule := range req.Rules {
		if rule.Id == 0 {
			return errors.Wrap(errNotPositiveID, "rule")
		}
	}

	return nil
}

func validateUpdateRuleRequest(req *desc.UpdateRuleRequest) error {
	if req.Rule.Id == 0 {
		return errors.Wrap(errNotPositiveID, "rule")
	}

	return nil
}
