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

// ScheduleRequest schedule request
//
// swagger:model ScheduleRequest
type ScheduleRequest struct {

	// cron
	Cron *CronSchedule `json:"cron,omitempty"`

	// The unique per installed app name of the schedule.
	// Required: true
	// Max Length: 36
	// Min Length: 1
	Name *string `json:"name"`

	// once
	Once *OnceSchedule `json:"once,omitempty"`
}

// Validate validates this schedule request
func (m *ScheduleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCron(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnce(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleRequest) validateCron(formats strfmt.Registry) error {

	if swag.IsZero(m.Cron) { // not required
		return nil
	}

	if m.Cron != nil {
		if err := m.Cron.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cron")
			}
			return err
		}
	}

	return nil
}

func (m *ScheduleRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 36); err != nil {
		return err
	}

	return nil
}

func (m *ScheduleRequest) validateOnce(formats strfmt.Registry) error {

	if swag.IsZero(m.Once) { // not required
		return nil
	}

	if m.Once != nil {
		if err := m.Once.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("once")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleRequest) UnmarshalBinary(b []byte) error {
	var res ScheduleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
