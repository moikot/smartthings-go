// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Switch switch
//
// swagger:model switch
type Switch struct {
	SwitchAllOf0

	SwitchCommand

	SwitchState

	SwitchAllOf3
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Switch) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SwitchAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SwitchAllOf0 = aO0

	// AO1
	var aO1 SwitchCommand
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SwitchCommand = aO1

	// AO2
	var aO2 SwitchState
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.SwitchState = aO2

	// AO3
	var aO3 SwitchAllOf3
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.SwitchAllOf3 = aO3

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Switch) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 4)

	aO0, err := swag.WriteJSON(m.SwitchAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SwitchCommand)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.SwitchState)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.SwitchAllOf3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this switch
func (m *Switch) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SwitchAllOf0
	// validation for a type composition with SwitchCommand
	if err := m.SwitchCommand.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SwitchState
	if err := m.SwitchState.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SwitchAllOf3

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Switch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Switch) UnmarshalBinary(b []byte) error {
	var res Switch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SwitchAllOf0 This switch will display as a slider style on/off switch in the UI
//
// swagger:model SwitchAllOf0
type SwitchAllOf0 interface{}

// SwitchAllOf3 switch all of3
//
// swagger:model SwitchAllOf3
type SwitchAllOf3 interface{}
