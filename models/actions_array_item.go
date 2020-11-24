// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActionsArrayItem actions array item
//
// swagger:model actionsArrayItem
type ActionsArrayItem struct {
	ActionItem

	CapabilityKey

	// component
	Component Component `json:"component,omitempty"`

	// visible condition
	VisibleCondition *VisibleCondition `json:"visibleCondition,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ActionsArrayItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ActionItem
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ActionItem = aO0

	// AO1
	var aO1 CapabilityKey
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CapabilityKey = aO1

	// AO2
	var dataAO2 struct {
		Component Component `json:"component,omitempty"`

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
func (m ActionsArrayItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.ActionItem)
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

// Validate validates this actions array item
func (m *ActionsArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ActionItem
	if err := m.ActionItem.Validate(formats); err != nil {
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

func (m *ActionsArrayItem) validateComponent(formats strfmt.Registry) error {

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

func (m *ActionsArrayItem) validateVisibleCondition(formats strfmt.Registry) error {

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
func (m *ActionsArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionsArrayItem) UnmarshalBinary(b []byte) error {
	var res ActionsArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
