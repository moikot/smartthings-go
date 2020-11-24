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

// DeviceInstallRequest device install request
//
// swagger:model DeviceInstallRequest
type DeviceInstallRequest struct {

	// app
	// Required: true
	App *DeviceInstallRequestApp `json:"app"`

	// The label for the device.
	Label string `json:"label,omitempty"`

	// The ID of the Location with which the device is associated.
	// Required: true
	LocationID *string `json:"locationId"`
}

// Validate validates this device install request
func (m *DeviceInstallRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceInstallRequest) validateApp(formats strfmt.Registry) error {

	if err := validate.Required("app", "body", m.App); err != nil {
		return err
	}

	if m.App != nil {
		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceInstallRequest) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("locationId", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInstallRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInstallRequest) UnmarshalBinary(b []byte) error {
	var res DeviceInstallRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DeviceInstallRequestApp device install request app
//
// swagger:model DeviceInstallRequestApp
type DeviceInstallRequestApp struct {

	// A field to store an ID from a system external to SmartThings.
	// Max Length: 64
	ExternalID string `json:"externalId,omitempty"`

	// The ID of the installed application
	// Required: true
	InstalledAppID *string `json:"installedAppId"`

	// The device profile Id
	// Required: true
	ProfileID *string `json:"profileId"`
}

// Validate validates this device install request app
func (m *DeviceInstallRequestApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstalledAppID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceInstallRequestApp) validateExternalID(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalID) { // not required
		return nil
	}

	if err := validate.MaxLength("app"+"."+"externalId", "body", string(m.ExternalID), 64); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInstallRequestApp) validateInstalledAppID(formats strfmt.Registry) error {

	if err := validate.Required("app"+"."+"installedAppId", "body", m.InstalledAppID); err != nil {
		return err
	}

	return nil
}

func (m *DeviceInstallRequestApp) validateProfileID(formats strfmt.Registry) error {

	if err := validate.Required("app"+"."+"profileId", "body", m.ProfileID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceInstallRequestApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceInstallRequestApp) UnmarshalBinary(b []byte) error {
	var res DeviceInstallRequestApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
