// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LocationActionExecutionResult The result of a location action execution
//
// swagger:model LocationActionExecutionResult
type LocationActionExecutionResult struct {

	// Location ID for location actions
	// Required: true
	LocationID *string `json:"locationId"`

	// result
	// Required: true
	Result ExecutionResult `json:"result"`
}

// Validate validates this location action execution result
func (m *LocationActionExecutionResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationActionExecutionResult) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("locationId", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *LocationActionExecutionResult) validateResult(formats strfmt.Registry) error {

	if err := m.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocationActionExecutionResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationActionExecutionResult) UnmarshalBinary(b []byte) error {
	var res LocationActionExecutionResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}