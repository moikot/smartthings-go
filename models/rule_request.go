// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RuleRequest rule request
//
// swagger:model RuleRequest
type RuleRequest struct {

	// actions
	// Required: true
	Actions []*Action `json:"actions"`

	// Name for the rule
	// Required: true
	Name *string `json:"name"`

	// The sequence in which the actions are to be executed i.e. Serial (default) or Parallel
	Sequence *ActionSequence `json:"sequence,omitempty"`

	// Time zone ID for this rule. This overrides the location time zone ID, but is overridden by time zone ID provided by each operand individually
	TimeZoneID string `json:"timeZoneId,omitempty"`
}

// Validate validates this rule request
func (m *RuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleRequest) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("actions", "body", m.Actions); err != nil {
		return err
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RuleRequest) validateSequence(formats strfmt.Registry) error {

	if swag.IsZero(m.Sequence) { // not required
		return nil
	}

	if m.Sequence != nil {
		if err := m.Sequence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sequence")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleRequest) UnmarshalBinary(b []byte) error {
	var res RuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}