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

// ArrayOperand array operand
//
// swagger:model ArrayOperand
type ArrayOperand struct {

	// aggregation
	Aggregation OperandAggregationMode `json:"aggregation,omitempty"`

	// operands
	// Required: true
	Operands []*Operand `json:"operands"`
}

// Validate validates this array operand
func (m *ArrayOperand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperands(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArrayOperand) validateAggregation(formats strfmt.Registry) error {

	if swag.IsZero(m.Aggregation) { // not required
		return nil
	}

	if err := m.Aggregation.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("aggregation")
		}
		return err
	}

	return nil
}

func (m *ArrayOperand) validateOperands(formats strfmt.Registry) error {

	if err := validate.Required("operands", "body", m.Operands); err != nil {
		return err
	}

	for i := 0; i < len(m.Operands); i++ {
		if swag.IsZero(m.Operands[i]) { // not required
			continue
		}

		if m.Operands[i] != nil {
			if err := m.Operands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArrayOperand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArrayOperand) UnmarshalBinary(b []byte) error {
	var res ArrayOperand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
