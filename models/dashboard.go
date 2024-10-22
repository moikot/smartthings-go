// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Dashboard dashboard
//
// swagger:model dashboard
type Dashboard struct {

	// actions
	Actions []*ActionsArrayItem `json:"actions"`

	// basic plus
	BasicPlus []*BasicPlusArrayItem `json:"basicPlus"`

	// states
	States []*StatesArrayItem `json:"states"`
}

// Validate validates this dashboard
func (m *Dashboard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBasicPlus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Dashboard) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Dashboard) validateBasicPlus(formats strfmt.Registry) error {

	if swag.IsZero(m.BasicPlus) { // not required
		return nil
	}

	for i := 0; i < len(m.BasicPlus); i++ {
		if swag.IsZero(m.BasicPlus[i]) { // not required
			continue
		}

		if m.BasicPlus[i] != nil {
			if err := m.BasicPlus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("basicPlus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Dashboard) validateStates(formats strfmt.Registry) error {

	if swag.IsZero(m.States) { // not required
		return nil
	}

	for i := 0; i < len(m.States); i++ {
		if swag.IsZero(m.States[i]) { // not required
			continue
		}

		if m.States[i] != nil {
			if err := m.States[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("states" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Dashboard) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Dashboard) UnmarshalBinary(b []byte) error {
	var res Dashboard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
