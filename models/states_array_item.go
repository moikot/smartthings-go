// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatesArrayItem states array item
//
// swagger:model statesArrayItem
type StatesArrayItem struct {
	StateItem

	CapabilityKey

	// component
	// Required: true
	Component Component `json:"component"`

	// visible condition
	VisibleCondition *VisibleCondition `json:"visibleCondition,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StatesArrayItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 StateItem
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.StateItem = aO0

	// AO1
	var aO1 CapabilityKey
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CapabilityKey = aO1

	// AO2
	var dataAO2 struct {
		Component Component `json:"component"`

		VisibleCondition *VisibleCondition `json:"visibleCondition,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Component = dataAO2.Component

	m.VisibleCondition = dataAO2.VisibleCondition

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StatesArrayItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.StateItem)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.CapabilityKey)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	var dataAO2 struct {
		Component Component `json:"component"`

		VisibleCondition *VisibleCondition `json:"visibleCondition,omitempty"`
	}

	dataAO2.Component = m.Component

	dataAO2.VisibleCondition = m.VisibleCondition

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this states array item
func (m *StatesArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with StateItem
	if err := m.StateItem.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CapabilityKey
	if err := m.CapabilityKey.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibleCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatesArrayItem) validateComponent(formats strfmt.Registry) error {

	if err := m.Component.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("component")
		}
		return err
	}

	return nil
}

func (m *StatesArrayItem) validateVisibleCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.VisibleCondition) { // not required
		return nil
	}

	if m.VisibleCondition != nil {
		if err := m.VisibleCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibleCondition")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatesArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatesArrayItem) UnmarshalBinary(b []byte) error {
	var res StatesArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
