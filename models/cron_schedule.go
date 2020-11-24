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

// CronSchedule cron schedule
//
// swagger:model CronSchedule
type CronSchedule struct {

	// The cron expression for the schedule for CRON schedules.
	// The format matches that specified by the [Quartz scheduler](http://www.quartz-scheduler.org/documentation/quartz-2.x/tutorials/crontrigger.html) but should not include the seconds (1st)
	// field. The exact second will be chosen at random but will remain consistent. The years part must be les than 2 years from now.
	//
	// Required: true
	// Max Length: 256
	Expression *string `json:"expression"`

	// The timezone id for CRON schedules.
	// Required: true
	Timezone *string `json:"timezone"`
}

// Validate validates this cron schedule
func (m *CronSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CronSchedule) validateExpression(formats strfmt.Registry) error {

	if err := validate.Required("expression", "body", m.Expression); err != nil {
		return err
	}

	if err := validate.MaxLength("expression", "body", string(*m.Expression), 256); err != nil {
		return err
	}

	return nil
}

func (m *CronSchedule) validateTimezone(formats strfmt.Registry) error {

	if err := validate.Required("timezone", "body", m.Timezone); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CronSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CronSchedule) UnmarshalBinary(b []byte) error {
	var res CronSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}