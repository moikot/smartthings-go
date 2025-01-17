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

// CreateOrUpdateLambdaSmartAppRequest Details related to a Lambda Smart App implementation.  This model should only be specified for apps of type LAMBDA_SMART_APP.
//
//
// swagger:model CreateOrUpdateLambdaSmartAppRequest
type CreateOrUpdateLambdaSmartAppRequest struct {

	// A list of AWS arns referencing a Lambda function.
	// Required: true
	Functions []string `json:"functions"`
}

// Validate validates this create or update lambda smart app request
func (m *CreateOrUpdateLambdaSmartAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunctions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrUpdateLambdaSmartAppRequest) validateFunctions(formats strfmt.Registry) error {

	if err := validate.Required("functions", "body", m.Functions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrUpdateLambdaSmartAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrUpdateLambdaSmartAppRequest) UnmarshalBinary(b []byte) error {
	var res CreateOrUpdateLambdaSmartAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
