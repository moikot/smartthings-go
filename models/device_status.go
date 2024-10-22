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

// DeviceStatus The status of a Device.
//
// swagger:model DeviceStatus
type DeviceStatus struct {

	// A map of componentId to Component status.
	Components map[string]ComponentStatus `json:"components,omitempty"`

	// health state
	HealthState *DeviceStatusHealthState `json:"healthState,omitempty"`
}

// Validate validates this device status
func (m *DeviceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatus) validateComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for k := range m.Components {

		if val, ok := m.Components[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DeviceStatus) validateHealthState(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthState) { // not required
		return nil
	}

	if m.HealthState != nil {
		if err := m.HealthState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthState")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStatus) UnmarshalBinary(b []byte) error {
	var res DeviceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceStatusHealthState device status health state
//
// swagger:model DeviceStatusHealthState
type DeviceStatusHealthState struct {

	// Last reported date/time in UTC
	// Format: date-time
	LastUpdatedDate strfmt.DateTime `json:"lastUpdatedDate,omitempty"`

	// Current state of the device
	// Enum: [ONLINE OFFLINE]
	State string `json:"state,omitempty"`
}

// Validate validates this device status health state
func (m *DeviceStatusHealthState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceStatusHealthState) validateLastUpdatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("healthState"+"."+"lastUpdatedDate", "body", "date-time", m.LastUpdatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var deviceStatusHealthStateTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONLINE","OFFLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceStatusHealthStateTypeStatePropEnum = append(deviceStatusHealthStateTypeStatePropEnum, v)
	}
}

const (

	// DeviceStatusHealthStateStateONLINE captures enum value "ONLINE"
	DeviceStatusHealthStateStateONLINE string = "ONLINE"

	// DeviceStatusHealthStateStateOFFLINE captures enum value "OFFLINE"
	DeviceStatusHealthStateStateOFFLINE string = "OFFLINE"
)

// prop value enum
func (m *DeviceStatusHealthState) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deviceStatusHealthStateTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeviceStatusHealthState) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("healthState"+"."+"state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceStatusHealthState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceStatusHealthState) UnmarshalBinary(b []byte) error {
	var res DeviceStatusHealthState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
