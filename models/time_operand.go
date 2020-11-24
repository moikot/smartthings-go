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

// TimeOperand time operand
//
// swagger:model TimeOperand
type TimeOperand struct {

	// days of week
	DaysOfWeek []DayOfWeek `json:"daysOfWeek"`

	// offset
	Offset *Interval `json:"offset,omitempty"`

	// reference
	// Required: true
	Reference TimeReference `json:"reference"`

	// A java time zone ID reference
	TimeZoneID string `json:"timeZoneId,omitempty"`
}

// Validate validates this time operand
func (m *TimeOperand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaysOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffset(formats); err != nil {
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

func (m *TimeOperand) validateDaysOfWeek(formats strfmt.Registry) error {

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

func (m *TimeOperand) validateOffset(formats strfmt.Registry) error {

	if swag.IsZero(m.Offset) { // not required
		return nil
	}

	if m.Offset != nil {
		if err := m.Offset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("offset")
			}
			return err
		}
	}

	return nil
}

func (m *TimeOperand) validateReference(formats strfmt.Registry) error {

	if err := m.Reference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reference")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeOperand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeOperand) UnmarshalBinary(b []byte) error {
	var res TimeOperand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}