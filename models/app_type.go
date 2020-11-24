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

// AppType Denotes the type of app.
//
// swagger:model AppType
type AppType string

const (

	// AppTypeLAMBDASMARTAPP captures enum value "LAMBDA_SMART_APP"
	AppTypeLAMBDASMARTAPP AppType = "LAMBDA_SMART_APP"

	// AppTypeWEBHOOKSMARTAPP captures enum value "WEBHOOK_SMART_APP"
	AppTypeWEBHOOKSMARTAPP AppType = "WEBHOOK_SMART_APP"
)

// for schema
var appTypeEnum []interface{}

func init() {
	var res []AppType
	if err := json.Unmarshal([]byte(`["LAMBDA_SMART_APP","WEBHOOK_SMART_APP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appTypeEnum = append(appTypeEnum, v)
	}
}

func (m AppType) validateAppTypeEnum(path, location string, value AppType) error {
	if err := validate.EnumCase(path, location, value, appTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this app type
func (m AppType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAppTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}