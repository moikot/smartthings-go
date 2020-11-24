// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceCommandsEvent An event that contains commands for devices that were created by this app.
//
// swagger:model DeviceCommandsEvent
type DeviceCommandsEvent struct {

	// commands
	Commands []*DeviceCommandsEventCommand `json:"commands"`

	// The guid of the device that the commands are for.
	DeviceID string `json:"deviceId,omitempty"`

	// The id of the event.
	EventID string `json:"eventId,omitempty"`

	// The external ID that was set during install of a device.
	ExternalID string `json:"externalId,omitempty"`

	// The device profile ID of the device instance.
	ProfileID string `json:"profileId,omitempty"`
}

// Validate validates this device commands event
func (m *DeviceCommandsEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceCommandsEvent) validateCommands(formats strfmt.Registry) error {

	if swag.IsZero(m.Commands) { // not required
		return nil
	}

	for i := 0; i < len(m.Commands); i++ {
		if swag.IsZero(m.Commands[i]) { // not required
			continue
		}

		if m.Commands[i] != nil {
			if err := m.Commands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceCommandsEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceCommandsEvent) UnmarshalBinary(b []byte) error {
	var res DeviceCommandsEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}