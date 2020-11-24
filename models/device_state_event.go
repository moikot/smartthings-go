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

// DeviceStateEvent device state event
//
// swagger:model DeviceStateEvent
type DeviceStateEvent struct {

	// Name of the capability attribute that this event relates to.
	Attribute string `json:"attribute,omitempty"`

	// Capability that this event relates to.
	Capability string `json:"capability,omitempty"`

	// The name of the component on this device, default is 'main'.
	Component string `json:"component,omitempty"`

	// Key value pairs about the state of the device. Valid values depend on the capability.
	Data map[string]interface{} `json:"data,omitempty"`

	// Unit of the value field.
	Unit string `json:"unit,omitempty"`

	// Value associated with the event. The valid value depends on the capability.
	// Required: true
	Value interface{} `json:"value"`
}

// Validate validates this device state event
func (m *DeviceStateEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStateEvent) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStateEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStateEvent) UnmarshalBinary(b []byte) error {
	var res DeviceStateEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}