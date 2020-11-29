// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package models

import (
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes"
	iav1beta1 "github.com/percona/pmm/api/managementpb/ia"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	reform "gopkg.in/reform.v1"
)

// SaveRule persists alert rule.
func SaveRule(q *reform.Querier, r *Rule) error {
	if err := ValidateRule(r); err != nil {
		return err
	}

	nc, err := ruleToAlertRule(r)
	if err != nil {
		return err
	}

	err = q.Insert(nc)
	if err != nil {
		return errors.Wrap(err, "failed to create alert rule")
	}

	return nil
}

// UpdateRule updates existing alert rule.
func UpdateRule(q *reform.Querier, r *Rule) error {
	if err := ValidateRule(r); err != nil {
		return errors.Wrap(err, "channel validation failed")
	}

	nc, err := ruleToAlertRule(r)
	if err != nil {
		return err
	}

	err = q.Update(nc)
	if err != nil {
		return errors.Wrap(err, "failed to create alert rule")
	}

	return nil
}

// RemoveRule removes alert rule with specified id.
func RemoveRule(q *reform.Querier, id string) error {
	err := q.Delete(&alertRule{ID: id})
	if err != nil {
		return errors.Wrap(err, "failed to delete alert rule")
	}
	return nil
}

// GetRules returns saved alert rules configuration.
func GetRules(q *reform.Querier) ([]Rule, error) {
	structs, err := q.SelectAllFrom(alertRuleTable, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to select alert rules")

	}

	rules := make([]Rule, len(structs))
	for i, s := range structs {
		c, err := alertRuleToRule(s.(*alertRule))
		if err != nil {
			return nil, err
		}
		rules[i] = *c
	}

	return rules, nil
}

// ValidateRule validates alert rule.
func ValidateRule(r *Rule) error {
	if r.ID == "" {
		return status.Error(codes.InvalidArgument, "alert rule id is empty")
	}
	return nil
}

func ruleToAlertRule(r *Rule) (*alertRule, error) {
	ar := &alertRule{
		ID:        r.ID,
		Disabled:  r.Disabled,
		For:       r.For.String(),
		Severity:  r.Severity.String(),
		CreatedAt: r.CreatedAt,
	}

	t, err := json.Marshal(r.Template)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall template")
	}
	ar.Template = &t

	p, err := json.Marshal(r.Params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall params")
	}
	ar.Params = &p

	cl, err := json.Marshal(r.CustomLabels)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall custom labels")
	}
	ar.CustomLabels = &cl

	f, err := json.Marshal(r.Filters)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall filters")
	}
	ar.Filters = &f

	c, err := json.Marshal(r.Channels)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall channels")
	}
	ar.Channels = &c

	return ar, nil
}

func alertRuleToRule(ar *alertRule) (*Rule, error) {
	r := &Rule{
		ID:       ar.ID,
		Disabled: ar.Disabled,
	}

	dur, err := time.ParseDuration(ar.For)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert duration")
	}
	r.For = dur

	createdAt, err := ptypes.TimestampProto(ar.CreatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert timestamp")
	}
	r.CreatedAt = createdAt.AsTime()

	r.Template = &iav1beta1.Template{}
	err = json.Unmarshal(*ar.Template, r.Template)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall template")
	}

	var params []*iav1beta1.RuleParam
	err = json.Unmarshal(*ar.Params, &params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall params")
	}
	r.Params = params

	var filters []*Filter
	err = json.Unmarshal(*ar.Filters, &filters)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall filters")
	}
	r.Filters = filters

	var channels []*iav1beta1.Channel
	err = json.Unmarshal(*ar.Channels, &channels)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall channels")
	}
	r.Channels = channels

	return r, nil
}