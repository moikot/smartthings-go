// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListForAutomationAction list for automation action
//
// swagger:model listForAutomationAction
type ListForAutomationAction struct {
	ListBase

	// The name of the command that is called when an item is chosen from the list
	Command string `json:"command,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ListForAutomationAction) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ListBase
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ListBase = aO0

	// AO1
	var dataAO1 struct {
		Command string `json:"command,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Command = dataAO1.Command

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ListForAutomationAction) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ListBase)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Command string `json:"command,omitempty"`
	}

	dataAO1.Command = m.Command

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this list for automation action
func (m *ListForAutomationAction) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ListBase
	if err := m.ListBase.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ListForAutomationAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListForAutomationAction) UnmarshalBinary(b []byte) error {
	var res ListForAutomationAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
