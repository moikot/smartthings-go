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

// SubscriptionFilterTypes The type of filter
//
// swagger:model SubscriptionFilterTypes
type SubscriptionFilterTypes string

const (

	// SubscriptionFilterTypesLOCATIONIDS captures enum value "LOCATIONIDS"
	SubscriptionFilterTypesLOCATIONIDS SubscriptionFilterTypes = "LOCATIONIDS"

	// SubscriptionFilterTypesROOMIDS captures enum value "ROOMIDS"
	SubscriptionFilterTypesROOMIDS SubscriptionFilterTypes = "ROOMIDS"

	// SubscriptionFilterTypesDEVICEIDS captures enum value "DEVICEIDS"
	SubscriptionFilterTypesDEVICEIDS SubscriptionFilterTypes = "DEVICEIDS"

	// SubscriptionFilterTypesINSTALLEDSMARTAPPIDS captures enum value "INSTALLEDSMARTAPPIDS"
	SubscriptionFilterTypesINSTALLEDSMARTAPPIDS SubscriptionFilterTypes = "INSTALLEDSMARTAPPIDS"

	// SubscriptionFilterTypesSMARTAPPIDS captures enum value "SMARTAPPIDS"
	SubscriptionFilterTypesSMARTAPPIDS SubscriptionFilterTypes = "SMARTAPPIDS"
)

// for schema
var subscriptionFilterTypesEnum []interface{}

func init() {
	var res []SubscriptionFilterTypes
	if err := json.Unmarshal([]byte(`["LOCATIONIDS","ROOMIDS","DEVICEIDS","INSTALLEDSMARTAPPIDS","SMARTAPPIDS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		subscriptionFilterTypesEnum = append(subscriptionFilterTypesEnum, v)
	}
}

func (m SubscriptionFilterTypes) validateSubscriptionFilterTypesEnum(path, location string, value SubscriptionFilterTypes) error {
	if err := validate.EnumCase(path, location, value, subscriptionFilterTypesEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this subscription filter types
func (m SubscriptionFilterTypes) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSubscriptionFilterTypesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
