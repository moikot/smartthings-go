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

// SliderBase slider base
//
// swagger:model sliderBase
type SliderBase struct {

	// range
	// Required: true
	Range Range `json:"range"`

	// step
	Step Step `json:"step,omitempty"`

	// unit
	Unit Unit `json:"unit,omitempty"`
}

// Validate validates this slider base
func (m *SliderBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SliderBase) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	if err := m.Range.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("range")
		}
		return err
	}

	return nil
}

func (m *SliderBase) validateStep(formats strfmt.Registry) error {

	if swag.IsZero(m.Step) { // not required
		return nil
	}

	if err := m.Step.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("step")
		}
		return err
	}

	return nil
}

func (m *SliderBase) validateUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if err := m.Unit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("unit")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SliderBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SliderBase) UnmarshalBinary(b []byte) error {
	var res SliderBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
