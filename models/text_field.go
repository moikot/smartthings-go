// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TextField This display type creates textfield that the user enters a new value and that the new value is passed as an argument to the command
//
// swagger:model textField
type TextField struct {

	// command
	// Required: true
	Command *string `json:"command"`

	// range
	Range Range `json:"range,omitempty"`

	// value
	Value Value `json:"value,omitempty"`
}

// Validate validates this text field
func (m *TextField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

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

func (m *TextField) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

func (m *TextField) validateRange(formats strfmt.Registry) error {

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

func (m *TextField) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := m.Value.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("value")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextField) UnmarshalBinary(b []byte) error {
	var res TextField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
