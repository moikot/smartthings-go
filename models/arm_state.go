// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ArmState arm state
//
// swagger:model ArmState
type ArmState string

const (

	// ArmStateArmedStay captures enum value "ArmedStay"
	ArmStateArmedStay ArmState = "ArmedStay"

	// ArmStateArmedAway captures enum value "ArmedAway"
	ArmStateArmedAway ArmState = "ArmedAway"

	// ArmStateDisarmed captures enum value "Disarmed"
	ArmStateDisarmed ArmState = "Disarmed"
)

// for schema
var armStateEnum []interface{}

func init() {
	var res []ArmState
	if err := json.Unmarshal([]byte(`["ArmedStay","ArmedAway","Disarmed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		armStateEnum = append(armStateEnum, v)
	}
}

func (m ArmState) validateArmStateEnum(path, location string, value ArmState) error {
	if err := validate.EnumCase(path, location, value, armStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this arm state
func (m ArmState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateArmStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
