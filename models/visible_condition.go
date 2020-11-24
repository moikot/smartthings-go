// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VisibleCondition visible condition
//
// swagger:model visibleCondition
type VisibleCondition struct {
	CapabilityKey

	// The component that controls the visibility of this component. This can be another component or this one.
	Component Component `json:"component,omitempty"`

	// The value that the visible condition evaluates against.
	Operand string `json:"operand,omitempty"`

	// operator
	Operator Operator `json:"operator,omitempty"`

	// value
	Value Value `json:"value,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VisibleCondition) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CapabilityKey
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CapabilityKey = aO0

	// AO1
	var dataAO1 struct {
		Component Component `json:"component,omitempty"`

		Operand string `json:"operand,omitempty"`

		Operator Operator `json:"operator,omitempty"`

		Value Value `json:"value,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Component = dataAO1.Component

	m.Operand = dataAO1.Operand

	m.Operator = dataAO1.Operator

	m.Value = dataAO1.Value

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VisibleCondition) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CapabilityKey)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Component Component `json:"component,omitempty"`

		Operand string `json:"operand,omitempty"`

		Operator Operator `json:"operator,omitempty"`

		Value Value `json:"value,omitempty"`
	}

	dataAO1.Component = m.Component

	dataAO1.Operand = m.Operand

	dataAO1.Operator = m.Operator

	dataAO1.Value = m.Value

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this visible condition
func (m *VisibleCondition) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CapabilityKey
	if err := m.CapabilityKey.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
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

func (m *VisibleCondition) validateComponent(formats strfmt.Registry) error {

	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if err := m.Component.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("component")
		}
		return err
	}

	return nil
}

func (m *VisibleCondition) validateOperator(formats strfmt.Registry) error {

	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	if err := m.Operator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator")
		}
		return err
	}

	return nil
}

func (m *VisibleCondition) validateValue(formats strfmt.Registry) error {

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
func (m *VisibleCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisibleCondition) UnmarshalBinary(b []byte) error {
	var res VisibleCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
