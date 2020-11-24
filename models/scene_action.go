// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SceneAction Possible scene actions, mutually exclusive
//
// swagger:model SceneAction
type SceneAction struct {

	// device group request
	DeviceGroupRequest *SceneDeviceGroupRequest `json:"deviceGroupRequest,omitempty"`

	// device request
	DeviceRequest *SceneDeviceRequest `json:"deviceRequest,omitempty"`

	// mode request
	ModeRequest *SceneModeRequest `json:"modeRequest,omitempty"`

	// sleep request
	SleepRequest *SceneSleepRequest `json:"sleepRequest,omitempty"`
}

// Validate validates this scene action
func (m *SceneAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceGroupRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSleepRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SceneAction) validateDeviceGroupRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceGroupRequest) { // not required
		return nil
	}

	if m.DeviceGroupRequest != nil {
		if err := m.DeviceGroupRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceGroupRequest")
			}
			return err
		}
	}

	return nil
}

func (m *SceneAction) validateDeviceRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceRequest) { // not required
		return nil
	}

	if m.DeviceRequest != nil {
		if err := m.DeviceRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deviceRequest")
			}
			return err
		}
	}

	return nil
}

func (m *SceneAction) validateModeRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.ModeRequest) { // not required
		return nil
	}

	if m.ModeRequest != nil {
		if err := m.ModeRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modeRequest")
			}
			return err
		}
	}

	return nil
}

func (m *SceneAction) validateSleepRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.SleepRequest) { // not required
		return nil
	}

	if m.SleepRequest != nil {
		if err := m.SleepRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sleepRequest")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SceneAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SceneAction) UnmarshalBinary(b []byte) error {
	var res SceneAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
