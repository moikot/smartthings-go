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

// Argument argument
//
// swagger:model Argument
type Argument struct {

	// A name that is unique within the command. Used for i18n and named argument command execution.
	// Required: true
	// Pattern: ^[[a-z]*([A-Z][a-z]*)*]{1,36}$
	Name *string `json:"name"`

	// Whether or not the argument must be supplied.
	// If the argument is at the end of the arguments array then it can be completely ignored.
	// If the argument is followed by another argument `null` must be supplied.
	//
	Optional *bool `json:"optional,omitempty"`

	// [JSON schema](http://json-schema.org/specification-links.html#draft-4) for the argument. The API implements JSON schema version 4. For more info regarding JSON schema, please read [Understanding JSON Schema](https://json-schema.org/understanding-json-schema/index.html).
	//
	// Required: true
	Schema interface{} `json:"schema"`
}

// Validate validates this argument
func (m *Argument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Argument) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[[a-z]*([A-Z][a-z]*)*]{1,36}$`); err != nil {
		return err
	}

	return nil
}

func (m *Argument) validateSchema(formats strfmt.Registry) error {

	if err := validate.Required("schema", "body", m.Schema); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Argument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Argument) UnmarshalBinary(b []byte) error {
	var res Argument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
