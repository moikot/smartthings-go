// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SceneDeviceGroup scene device group
//
// swagger:model SceneDeviceGroup
type SceneDeviceGroup struct {

	// capability
	Capability *SceneCapability `json:"capability,omitempty"`

	// the id of the device
	DeviceGroupID string `json:"deviceGroupId,omitempty"`
}

// Validate validates this scene device group
func (m *SceneDeviceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SceneDeviceGroup) validateCapability(formats strfmt.Registry) error {

	if swag.IsZero(m.Capability) { // not required
		return nil
	}

	if m.Capability != nil {
		if err := m.Capability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capability")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SceneDeviceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SceneDeviceGroup) UnmarshalBinary(b []byte) error {
	var res SceneDeviceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
