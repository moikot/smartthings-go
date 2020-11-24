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

// AdhocMessageTemplate A message template definition, representing a message in a specific locale and it's variables.
//
// swagger:model AdhocMessageTemplate
type AdhocMessageTemplate struct {

	// locale tag
	// Required: true
	LocaleTag LocaleTag `json:"localeTag"`

	// A message template string.  Specify variables using the double curly braces convention.
	// i.e. "Hello, {{ firstName }}!"
	//
	// Required: true
	Template *string `json:"template"`

	// A map<string,string> with the key representing the variable name, and the value representing the verbiage
	// to be replaced in template string.
	//
	Variables map[string]string `json:"variables,omitempty"`
}

// Validate validates this adhoc message template
func (m *AdhocMessageTemplate) Validate(formats strfmt.Registry) error {
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

func (m *AdhocMessageTemplate) validateLocaleTag(formats strfmt.Registry) error {

	if err := m.LocaleTag.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("localeTag")
		}
		return err
	}

	return nil
}

func (m *AdhocMessageTemplate) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdhocMessageTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdhocMessageTemplate) UnmarshalBinary(b []byte) error {
	var res AdhocMessageTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
