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

// SwitchStateForDashboard switch state for dashboard
//
// swagger:model switchStateForDashboard
type SwitchStateForDashboard struct {

	// state
	State *SwitchStateForDashboardState `json:"state,omitempty"`
}

// Validate validates this switch state for dashboard
func (m *SwitchStateForDashboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwitchStateForDashboard) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwitchStateForDashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwitchStateForDashboard) UnmarshalBinary(b []byte) error {
	var res SwitchStateForDashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SwitchStateForDashboardState To describe "on" and "off" state of a switch
//
// swagger:model SwitchStateForDashboardState
type SwitchStateForDashboardState struct {

	// off
	// Required: true
	Off *string `json:"off"`

	// on
	// Required: true
	On *string `json:"on"`

	// value
	Value Value `json:"value,omitempty"`
}

// Validate validates this switch state for dashboard state
func (m *SwitchStateForDashboardState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOff(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOn(formats); err != nil {
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

func (m *SwitchStateForDashboardState) validateOff(formats strfmt.Registry) error {

	if err := validate.Required("state"+"."+"off", "body", m.Off); err != nil {
		return err
	}

	return nil
}

func (m *SwitchStateForDashboardState) validateOn(formats strfmt.Registry) error {

	if err := validate.Required("state"+"."+"on", "body", m.On); err != nil {
		return err
	}

	return nil
}

func (m *SwitchStateForDashboardState) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := m.Value.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state" + "." + "value")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwitchStateForDashboardState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwitchStateForDashboardState) UnmarshalBinary(b []byte) error {
	var res SwitchStateForDashboardState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
