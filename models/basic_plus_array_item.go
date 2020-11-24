// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BasicPlusArrayItem basic plus array item
//
// swagger:model basicPlusArrayItem
type BasicPlusArrayItem struct {
	BasicPlusItem

	CapabilityKey

	// component
	Component Component `json:"component,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BasicPlusArrayItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BasicPlusItem
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BasicPlusItem = aO0

	// AO1
	var aO1 CapabilityKey
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CapabilityKey = aO1

	// AO2
	var dataAO2 struct {
		Component Component `json:"component,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.Component = dataAO2.Component

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BasicPlusArrayItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.BasicPlusItem)
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
		Component Component `json:"component,omitempty"`
	}

	dataAO2.Component = m.Component

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this basic plus array item
func (m *BasicPlusArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BasicPlusItem
	if err := m.BasicPlusItem.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with CapabilityKey
	if err := m.CapabilityKey.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BasicPlusArrayItem) validateComponent(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BasicPlusArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasicPlusArrayItem) UnmarshalBinary(b []byte) error {
	var res BasicPlusArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
