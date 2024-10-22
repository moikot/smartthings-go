// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActionExecutionResult The result of an action execution
//
// swagger:model ActionExecutionResult
type ActionExecutionResult struct {

	// action Id
	// Required: true
	ActionID *string `json:"actionId"`

	// command
	Command []*CommandActionExecutionResult `json:"command"`

	// if
	If *IfActionExecutionResult `json:"if,omitempty"`

	// location
	Location *LocationActionExecutionResult `json:"location,omitempty"`

	// sleep
	Sleep *SleepActionExecutionResult `json:"sleep,omitempty"`
}

// Validate validates this action execution result
func (m *ActionExecutionResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommand(formats); err != nil {
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

func (m *ActionExecutionResult) validateActionID(formats strfmt.Registry) error {

	if err := validate.Required("actionId", "body", m.ActionID); err != nil {
		return err
	}

	return nil
}

func (m *ActionExecutionResult) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.Command) { // not required
		return nil
	}

	for i := 0; i < len(m.Command); i++ {
		if swag.IsZero(m.Command[i]) { // not required
			continue
		}

		if m.Command[i] != nil {
			if err := m.Command[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("command" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActionExecutionResult) validateIf(formats strfmt.Registry) error {

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

func (m *ActionExecutionResult) validateLocation(formats strfmt.Registry) error {

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

func (m *ActionExecutionResult) validateSleep(formats strfmt.Registry) error {

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
func (m *ActionExecutionResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionExecutionResult) UnmarshalBinary(b []byte) error {
	var res ActionExecutionResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
