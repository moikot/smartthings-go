// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DetailViewListItem detail view list item
//
// swagger:model detailViewListItem
type DetailViewListItem struct {
	DetailViewItem

	// component
	Component Component `json:"component,omitempty"`

	// visible condition
	VisibleCondition *VisibleCondition `json:"visibleCondition,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DetailViewListItem) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DetailViewItem
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DetailViewItem = aO0

	// AO1
	var dataAO1 struct {
		Component Component `json:"component,omitempty"`

		VisibleCondition *VisibleCondition `json:"visibleCondition,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Component = dataAO1.Component

	m.VisibleCondition = dataAO1.VisibleCondition

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DetailViewListItem) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.DetailViewItem)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Component Component `json:"component,omitempty"`

		VisibleCondition *VisibleCondition `json:"visibleCondition,omitempty"`
	}

	dataAO1.Component = m.Component

	dataAO1.VisibleCondition = m.VisibleCondition

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this detail view list item
func (m *DetailViewListItem) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DetailViewItem
	if err := m.DetailViewItem.Validate(formats); err != nil {
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

func (m *DetailViewListItem) validateComponent(formats strfmt.Registry) error {

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

func (m *DetailViewListItem) validateVisibleCondition(formats strfmt.Registry) error {

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
func (m *DetailViewListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailViewListItem) UnmarshalBinary(b []byte) error {
	var res DetailViewListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
