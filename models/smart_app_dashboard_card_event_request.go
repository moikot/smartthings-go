// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SmartAppDashboardCardEventRequest An event used to trigger a client action on a SmartApp dashboard card.
//
//
// swagger:model SmartAppDashboardCardEventRequest
type SmartAppDashboardCardEventRequest struct {

	// A developer defined dashboard card ID.
	CardID string `json:"cardId,omitempty"`

	// lifecycle
	Lifecycle DashboardCardLifecycle `json:"lifecycle,omitempty"`
}

// Validate validates this smart app dashboard card event request
func (m *SmartAppDashboardCardEventRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLifecycle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmartAppDashboardCardEventRequest) validateLifecycle(formats strfmt.Registry) error {

	if swag.IsZero(m.Lifecycle) { // not required
		return nil
	}

	if err := m.Lifecycle.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lifecycle")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmartAppDashboardCardEventRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmartAppDashboardCardEventRequest) UnmarshalBinary(b []byte) error {
	var res SmartAppDashboardCardEventRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
