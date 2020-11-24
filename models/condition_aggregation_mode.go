// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ConditionAggregationMode condition aggregation mode
//
// swagger:model ConditionAggregationMode
type ConditionAggregationMode string

const (

	// ConditionAggregationModeAny captures enum value "Any"
	ConditionAggregationModeAny ConditionAggregationMode = "Any"

	// ConditionAggregationModeAll captures enum value "All"
	ConditionAggregationModeAll ConditionAggregationMode = "All"
)

// for schema
var conditionAggregationModeEnum []interface{}

func init() {
	var res []ConditionAggregationMode
	if err := json.Unmarshal([]byte(`["Any","All"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionAggregationModeEnum = append(conditionAggregationModeEnum, v)
	}
}

func (m ConditionAggregationMode) validateConditionAggregationModeEnum(path, location string, value ConditionAggregationMode) error {
	if err := validate.EnumCase(path, location, value, conditionAggregationModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this condition aggregation mode
func (m ConditionAggregationMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConditionAggregationModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}