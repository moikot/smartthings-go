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

// DateOperand date operand
//
// swagger:model DateOperand
type DateOperand struct {

	// day
	// Maximum: 31
	// Minimum: 1
	Day int32 `json:"day,omitempty"`

	// days of week
	DaysOfWeek []DayOfWeek `json:"daysOfWeek"`

	// month
	// Maximum: 12
	// Minimum: 1
	Month int32 `json:"month,omitempty"`

	// reference
	Reference DateReference `json:"reference,omitempty"`

	// A java time zone ID reference
	TimeZoneID string `json:"timeZoneId,omitempty"`

	// year
	Year int32 `json:"year,omitempty"`
}

// Validate validates this date operand
func (m *DateOperand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDaysOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DateOperand) validateDay(formats strfmt.Registry) error {

	if swag.IsZero(m.Day) { // not required
		return nil
	}

	if err := validate.MinimumInt("day", "body", int64(m.Day), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("day", "body", int64(m.Day), 31, false); err != nil {
		return err
	}

	return nil
}

func (m *DateOperand) validateDaysOfWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.DaysOfWeek) { // not required
		return nil
	}

	for i := 0; i < len(m.DaysOfWeek); i++ {

		if err := m.DaysOfWeek[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("daysOfWeek" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *DateOperand) validateMonth(formats strfmt.Registry) error {

	if swag.IsZero(m.Month) { // not required
		return nil
	}

	if err := validate.MinimumInt("month", "body", int64(m.Month), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("month", "body", int64(m.Month), 12, false); err != nil {
		return err
	}

	return nil
}

func (m *DateOperand) validateReference(formats strfmt.Registry) error {

	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if err := m.Reference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reference")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DateOperand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DateOperand) UnmarshalBinary(b []byte) error {
	var res DateOperand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
