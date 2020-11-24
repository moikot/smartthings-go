// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlternativeItem alternative item
//
// swagger:model alternativeItem
type AlternativeItem struct {

	// Additional description for each value. This description is shown in detail view or automation under this particular key.
	Description string `json:"description,omitempty"`

	// icon Url
	IconURL IconURL `json:"iconUrl,omitempty"`

	// The label string to be displayed when the 'key' is evaluated.
	// Required: true
	Key *string `json:"key"`

	// This shows the active or inactive state of the value. Active is shown in color, while inactive value is shown dimmed in the UI view. For example, Motion sensor capability would have `detected` as active and in color and `clear` as inactive and dimmed so that a user can see the `detected` event easily.
	// Enum: [inactive active]
	Type *string `json:"type,omitempty"`

	// The alternative string to be displayed when the `key` is selected.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this alternative item
func (m *AlternativeItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIconURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlternativeItem) validateIconURL(formats strfmt.Registry) error {

	if swag.IsZero(m.IconURL) { // not required
		return nil
	}

	if err := m.IconURL.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("iconUrl")
		}
		return err
	}

	return nil
}

func (m *AlternativeItem) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

var alternativeItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["inactive","active"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alternativeItemTypeTypePropEnum = append(alternativeItemTypeTypePropEnum, v)
	}
}

const (

	// AlternativeItemTypeInactive captures enum value "inactive"
	AlternativeItemTypeInactive string = "inactive"

	// AlternativeItemTypeActive captures enum value "active"
	AlternativeItemTypeActive string = "active"
)

// prop value enum
func (m *AlternativeItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alternativeItemTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AlternativeItem) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *AlternativeItem) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlternativeItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlternativeItem) UnmarshalBinary(b []byte) error {
	var res AlternativeItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}