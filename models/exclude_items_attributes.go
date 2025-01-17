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

// ExcludeItemsAttributes exclude items attributes
//
// swagger:model excludeItemsAttributes
type ExcludeItemsAttributes []*ExcludeItemsAttributesItems0

// Validate validates this exclude items attributes
func (m ExcludeItemsAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ExcludeItemsAttributesItems0 exclude items attributes items0
//
// swagger:model ExcludeItemsAttributesItems0
type ExcludeItemsAttributesItems0 struct {

	// To exclude only some values of the attribute
	ExcludedValues []string `json:"excludedValues"`

	// Attribute name to be excluded
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this exclude items attributes items0
func (m *ExcludeItemsAttributesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExcludeItemsAttributesItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExcludeItemsAttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExcludeItemsAttributesItems0) UnmarshalBinary(b []byte) error {
	var res ExcludeItemsAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
