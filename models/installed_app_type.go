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

// InstalledAppType Denotes the type of installed app.
//
// swagger:model InstalledAppType
type InstalledAppType string

const (

	// InstalledAppTypeLAMBDASMARTAPP captures enum value "LAMBDA_SMART_APP"
	InstalledAppTypeLAMBDASMARTAPP InstalledAppType = "LAMBDA_SMART_APP"

	// InstalledAppTypeWEBHOOKSMARTAPP captures enum value "WEBHOOK_SMART_APP"
	InstalledAppTypeWEBHOOKSMARTAPP InstalledAppType = "WEBHOOK_SMART_APP"
)

// for schema
var installedAppTypeEnum []interface{}

func init() {
	var res []InstalledAppType
	if err := json.Unmarshal([]byte(`["LAMBDA_SMART_APP","WEBHOOK_SMART_APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		installedAppTypeEnum = append(installedAppTypeEnum, v)
	}
}

func (m InstalledAppType) validateInstalledAppTypeEnum(path, location string, value InstalledAppType) error {
	if err := validate.EnumCase(path, location, value, installedAppTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this installed app type
func (m InstalledAppType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateInstalledAppTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}