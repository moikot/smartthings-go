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

// ExcludedConditionItem excluded condition item
//
// swagger:model excludedConditionItem
type ExcludedConditionItem struct {

	// exclude
	// Required: true
	Exclude []*ExcludedConditionItemExcludeItems0 `json:"exclude"`

	// To express that a specific value is selected
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this excluded condition item
func (m *ExcludedConditionItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclude(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExcludedConditionItem) validateExclude(formats strfmt.Registry) error {

	if err := validate.Required("exclude", "body", m.Exclude); err != nil {
		return err
	}

	for i := 0; i < len(m.Exclude); i++ {
		if swag.IsZero(m.Exclude[i]) { // not required
			continue
		}

		if m.Exclude[i] != nil {
			if err := m.Exclude[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("exclude" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExcludedConditionItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExcludedConditionItem) UnmarshalBinary(b []byte) error {
	var res ExcludedConditionItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ExcludedConditionItemExcludeItems0 excluded condition item exclude items0
//
// swagger:model ExcludedConditionItemExcludeItems0
type ExcludedConditionItemExcludeItems0 struct {

	// attributes
	Attributes ExcludeItemsAttributes `json:"attributes,omitempty"`

	// capability
	// Required: true
	Capability *string `json:"capability"`

	// component
	Component string `json:"component,omitempty"`

	// version
	Version *int64 `json:"version,omitempty"`
}

// Validate validates this excluded condition item exclude items0
func (m *ExcludedConditionItemExcludeItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExcludedConditionItemExcludeItems0) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if err := m.Attributes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes")
		}
		return err
	}

	return nil
}

func (m *ExcludedConditionItemExcludeItems0) validateCapability(formats strfmt.Registry) error {

	if err := validate.Required("capability", "body", m.Capability); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExcludedConditionItemExcludeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExcludedConditionItemExcludeItems0) UnmarshalBinary(b []byte) error {
	var res ExcludedConditionItemExcludeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
