// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StateItem state item
//
// swagger:model stateItem
type StateItem struct {

	// alternatives
	Alternatives Alternatives `json:"alternatives,omitempty"`

	// group
	Group Group `json:"group,omitempty"`

	// label
	// Required: true
	Label FormattedLabel `json:"label"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *StateItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Alternatives Alternatives `json:"alternatives,omitempty"`

		Group Group `json:"group,omitempty"`

		Label FormattedLabel `json:"label"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Alternatives = dataAO0.Alternatives

	m.Group = dataAO0.Group

	m.Label = dataAO0.Label

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m StateItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Alternatives Alternatives `json:"alternatives,omitempty"`

		Group Group `json:"group,omitempty"`

		Label FormattedLabel `json:"label"`
	}

	dataAO0.Alternatives = m.Alternatives

	dataAO0.Group = m.Group

	dataAO0.Label = m.Label

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this state item
func (m *StateItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlternatives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StateItem) validateAlternatives(formats strfmt.Registry) error {

	if swag.IsZero(m.Alternatives) { // not required
		return nil
	}

	if err := m.Alternatives.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("alternatives")
		}
		return err
	}

	return nil
}

func (m *StateItem) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if err := m.Group.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("group")
		}
		return err
	}

	return nil
}

func (m *StateItem) validateLabel(formats strfmt.Registry) error {

	if err := m.Label.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("label")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StateItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StateItem) UnmarshalBinary(b []byte) error {
	var res StateItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}