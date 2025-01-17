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

// MessageTemplate A message template definition, representing a message in a specific locale.
//
// swagger:model MessageTemplate
type MessageTemplate struct {

	// locale tag
	// Required: true
	LocaleTag LocaleTag `json:"localeTag"`

	// A message template string. Specify variables using the double curly braces convention.
	// i.e. "Hello, {{ firstName }}!"
	//
	// Required: true
	Template *string `json:"template"`
}

// Validate validates this message template
func (m *MessageTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocaleTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageTemplate) validateLocaleTag(formats strfmt.Registry) error {

	if err := m.LocaleTag.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("localeTag")
		}
		return err
	}

	return nil
}

func (m *MessageTemplate) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageTemplate) UnmarshalBinary(b []byte) error {
	var res MessageTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
