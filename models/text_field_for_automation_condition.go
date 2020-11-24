// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TextFieldForAutomationCondition text field for automation condition
//
// swagger:model textFieldForAutomationCondition
type TextFieldForAutomationCondition struct {

	// range
	Range Range `json:"range,omitempty"`

	// value
	// Required: true
	Value Value `json:"value"`
}

// Validate validates this text field for automation condition
func (m *TextFieldForAutomationCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextFieldForAutomationCondition) validateRange(formats strfmt.Registry) error {

	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if err := m.Range.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("range")
		}
		return err
	}

	return nil
}

func (m *TextFieldForAutomationCondition) validateValue(formats strfmt.Registry) error {

	if err := m.Value.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("value")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextFieldForAutomationCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextFieldForAutomationCondition) UnmarshalBinary(b []byte) error {
	var res TextFieldForAutomationCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
