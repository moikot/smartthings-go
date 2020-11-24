// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebhookSmartApp Details related to a Webhook Smart App implementation.  This model will only be available for apps of type WEBHOOK_SMART_APP.
//
//
// swagger:model WebhookSmartApp
type WebhookSmartApp struct {

	// The public half of an RSA key pair.  Useful for verifying a Webhook execution request signature to
	// ensure it came from SmartThings.
	//
	PublicKey string `json:"publicKey,omitempty"`

	// signature type
	SignatureType SignatureType `json:"signatureType,omitempty"`

	// target status
	TargetStatus AppTargetStatus `json:"targetStatus,omitempty"`

	// A URL that should be invoked during execution.
	TargetURL string `json:"targetUrl,omitempty"`
}

// Validate validates this webhook smart app
func (m *WebhookSmartApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSignatureType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookSmartApp) validateSignatureType(formats strfmt.Registry) error {

	if swag.IsZero(m.SignatureType) { // not required
		return nil
	}

	if err := m.SignatureType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("signatureType")
		}
		return err
	}

	return nil
}

func (m *WebhookSmartApp) validateTargetStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetStatus) { // not required
		return nil
	}

	if err := m.TargetStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetStatus")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebhookSmartApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookSmartApp) UnmarshalBinary(b []byte) error {
	var res WebhookSmartApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
