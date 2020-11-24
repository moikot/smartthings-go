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

// ComponentTranslations A map of i18n translations.
//
// swagger:model ComponentTranslations
type ComponentTranslations struct {

	// UTF-8 text describing the component.
	// Max Length: 255
	Description string `json:"description,omitempty"`

	// Short UTF-8 text used when displaying the component.
	// Max Length: 25
	Label string `json:"label,omitempty"`
}

// Validate validates this component translations
func (m *ComponentTranslations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComponentTranslations) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *ComponentTranslations) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 25); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComponentTranslations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComponentTranslations) UnmarshalBinary(b []byte) error {
	var res ComponentTranslations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
