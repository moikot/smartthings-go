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

// UpdateCapabilityRequest update capability request
//
// swagger:model UpdateCapabilityRequest
type UpdateCapabilityRequest struct {

	// A mapping of attribute names to their definitions. All attribute names are lower camelcase.
	Attributes map[string]CapabilityAttribute `json:"attributes,omitempty"`

	// A mapping of command names to their definitions. All command names are lower camelcase.
	Commands map[string]CapabilityCommand `json:"commands,omitempty"`
}

// Validate validates this update capability request
func (m *UpdateCapabilityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCapabilityRequest) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for k := range m.Attributes {

		if err := validate.Required("attributes"+"."+k, "body", m.Attributes[k]); err != nil {
			return err
		}
		if val, ok := m.Attributes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *UpdateCapabilityRequest) validateCommands(formats strfmt.Registry) error {

	if swag.IsZero(m.Commands) { // not required
		return nil
	}

	for k := range m.Commands {

		if err := validate.Required("commands"+"."+k, "body", m.Commands[k]); err != nil {
			return err
		}
		if val, ok := m.Commands[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCapabilityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCapabilityRequest) UnmarshalBinary(b []byte) error {
	var res UpdateCapabilityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
