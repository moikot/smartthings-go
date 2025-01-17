// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateInstalledAppEventsRequest create installed app events request
//
// swagger:model CreateInstalledAppEventsRequest
type CreateInstalledAppEventsRequest struct {

	// An array of smartapp dashboard card events used to trigger client behavior for dashboard cards.
	// Dashboard card events are directives to a SmartThings client to take actions in relation to lifecycle changes
	// to a specific dashboard card.  These events are not delivered to the web plugin event handler.
	//
	// Max Items: 8
	// Min Items: 1
	SmartAppDashboardCardEvents []*SmartAppDashboardCardEventRequest `json:"smartAppDashboardCardEvents"`

	// An array of smartapp events used to trigger client behavior in loaded web plugin detail pages.  Events will
	// be delivered to JavaScript event handler of all active client processes related to parameterized installed app.
	//
	// Max Items: 8
	// Min Items: 1
	SmartAppEvents []*SmartAppEventRequest `json:"smartAppEvents"`
}

// Validate validates this create installed app events request
func (m *CreateInstalledAppEventsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSmartAppDashboardCardEvents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmartAppEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateInstalledAppEventsRequest) validateSmartAppDashboardCardEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.SmartAppDashboardCardEvents) { // not required
		return nil
	}

	iSmartAppDashboardCardEventsSize := int64(len(m.SmartAppDashboardCardEvents))

	if err := validate.MinItems("smartAppDashboardCardEvents", "body", iSmartAppDashboardCardEventsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("smartAppDashboardCardEvents", "body", iSmartAppDashboardCardEventsSize, 8); err != nil {
		return err
	}

	for i := 0; i < len(m.SmartAppDashboardCardEvents); i++ {
		if swag.IsZero(m.SmartAppDashboardCardEvents[i]) { // not required
			continue
		}

		if m.SmartAppDashboardCardEvents[i] != nil {
			if err := m.SmartAppDashboardCardEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("smartAppDashboardCardEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateInstalledAppEventsRequest) validateSmartAppEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.SmartAppEvents) { // not required
		return nil
	}

	iSmartAppEventsSize := int64(len(m.SmartAppEvents))

	if err := validate.MinItems("smartAppEvents", "body", iSmartAppEventsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("smartAppEvents", "body", iSmartAppEventsSize, 8); err != nil {
		return err
	}

	for i := 0; i < len(m.SmartAppEvents); i++ {
		if swag.IsZero(m.SmartAppEvents[i]) { // not required
			continue
		}

		if m.SmartAppEvents[i] != nil {
			if err := m.SmartAppEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("smartAppEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateInstalledAppEventsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateInstalledAppEventsRequest) UnmarshalBinary(b []byte) error {
	var res CreateInstalledAppEventsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
