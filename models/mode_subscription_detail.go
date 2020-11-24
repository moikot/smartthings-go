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

// ModeSubscriptionDetail Details of a subscription of source type MODE. This will subscribe to all mode changes for the given location. The installed app can then act on the resulting mode change event accordingly.
//
// swagger:model ModeSubscriptionDetail
type ModeSubscriptionDetail struct {

	// The GUID for the location to subscribe to mode changes.
	// Required: true
	LocationID *string `json:"locationId"`
}

// Validate validates this mode subscription detail
func (m *ModeSubscriptionDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModeSubscriptionDetail) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("locationId", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModeSubscriptionDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModeSubscriptionDetail) UnmarshalBinary(b []byte) error {
	var res ModeSubscriptionDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
