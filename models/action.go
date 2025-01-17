// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Action action
//
// swagger:model Action
type Action struct {

	// command
	Command *CommandAction `json:"command,omitempty"`

	// every
	Every *EveryAction `json:"every,omitempty"`

	// if
	If *IfAction `json:"if,omitempty"`

	// location
	Location *LocationAction `json:"location,omitempty"`

	// sleep
	Sleep *SleepAction `json:"sleep,omitempty"`
}

// Validate validates this action
func (m *Action) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSleep(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Action) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.Command) { // not required
		return nil
	}

	if m.Command != nil {
		if err := m.Command.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("command")
			}
			return err
		}
	}

	return nil
}

func (m *Action) validateEvery(formats strfmt.Registry) error {

	if swag.IsZero(m.Every) { // not required
		return nil
	}

	if m.Every != nil {
		if err := m.Every.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("every")
			}
			return err
		}
	}

	return nil
}

func (m *Action) validateIf(formats strfmt.Registry) error {

	if swag.IsZero(m.If) { // not required
		return nil
	}

	if m.If != nil {
		if err := m.If.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("if")
			}
			return err
		}
	}

	return nil
}

func (m *Action) validateLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *Action) validateSleep(formats strfmt.Registry) error {

	if swag.IsZero(m.Sleep) { // not required
		return nil
	}

	if m.Sleep != nil {
		if err := m.Sleep.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sleep")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Action) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Action) UnmarshalBinary(b []byte) error {
	var res Action
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
