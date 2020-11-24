// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Unit unit
//
// swagger:model unit
type Unit string

// Validate validates this unit
func (m Unit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Pattern("", "body", string(m), `(^[[a-z]*([A-Z][a-z]*)*){1,36}(\.unit)+`); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}