// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CapabilityStatus A map of Attribute name to Attribute state.
//
// swagger:model CapabilityStatus
type CapabilityStatus map[string]AttributeState

// Validate validates this capability status
func (m CapabilityStatus) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := validate.Required(k, "body", m[k]); err != nil {
			return err
		}
		if val, ok := m[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
