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

// DeviceLifecycleDetail Details of a subscription of source type DEVICE_LIFECYCLE. Only one of deviceIds or locationId should be supplied.
//
// swagger:model DeviceLifecycleDetail
type DeviceLifecycleDetail struct {

	// An array of GUIDs of devices being subscribed to. A max of 20 GUIDs are allowed.
	// Max Items: 20
	DeviceIds []string `json:"deviceIds"`

	// The id of the location that both the app and source device are in.
	LocationID string `json:"locationId,omitempty"`

	// A name for the subscription that will be passed to the installed app.
	SubscriptionName string `json:"subscriptionName,omitempty"`
}

// Validate validates this device lifecycle detail
func (m *DeviceLifecycleDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceLifecycleDetail) validateDeviceIds(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceIds) { // not required
		return nil
	}

	iDeviceIdsSize := int64(len(m.DeviceIds))

	if err := validate.MaxItems("deviceIds", "body", iDeviceIdsSize, 20); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceLifecycleDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceLifecycleDetail) UnmarshalBinary(b []byte) error {
	var res DeviceLifecycleDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
