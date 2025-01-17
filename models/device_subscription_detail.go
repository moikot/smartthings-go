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

// DeviceSubscriptionDetail Details of a subscription of source type DEVICE. The combination of subscribed values must be unique per installed app.
//
// swagger:model DeviceSubscriptionDetail
type DeviceSubscriptionDetail struct {

	// Name of the capabilities attribute or * for all.
	// Max Length: 256
	// Min Length: 1
	Attribute string `json:"attribute,omitempty"`

	// Name of the capability that is subscribed to or * for all.
	// Max Length: 128
	// Min Length: 1
	Capability string `json:"capability,omitempty"`

	// The component ID on the device that is subscribed to or * for all.
	ComponentID *string `json:"componentId,omitempty"`

	// The GUID of the device that is subscribed to.
	// Required: true
	DeviceID *string `json:"deviceId"`

	// List of mode ID's that the subscription will execute for. If not provided then all modes will be supported.
	Modes []string `json:"modes"`

	// Only execute the subscription if the subscribed event is a state change from previous events.
	StateChangeOnly bool `json:"stateChangeOnly,omitempty"`

	// A name for the subscription that will be passed to the installed app. Must be unique per installed app.
	SubscriptionName string `json:"subscriptionName,omitempty"`

	// A particular value for the attribute that will trigger the subscription or * for all.
	// Max Length: 256
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this device subscription detail
func (m *DeviceSubscriptionDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttribute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceSubscriptionDetail) validateAttribute(formats strfmt.Registry) error {

	if swag.IsZero(m.Attribute) { // not required
		return nil
	}

	if err := validate.MinLength("attribute", "body", string(m.Attribute), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("attribute", "body", string(m.Attribute), 256); err != nil {
		return err
	}

	return nil
}

func (m *DeviceSubscriptionDetail) validateCapability(formats strfmt.Registry) error {

	if swag.IsZero(m.Capability) { // not required
		return nil
	}

	if err := validate.MinLength("capability", "body", string(m.Capability), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("capability", "body", string(m.Capability), 128); err != nil {
		return err
	}

	return nil
}

func (m *DeviceSubscriptionDetail) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceSubscriptionDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceSubscriptionDetail) UnmarshalBinary(b []byte) error {
	var res DeviceSubscriptionDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
