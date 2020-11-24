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

// CreateAppRequest create app request
//
// swagger:model CreateAppRequest
type CreateAppRequest struct {

	// A globally unique, developer-defined identifier for an app. It is alpha-numeric, may contain dashes,
	// underscores, periods, and must be less then 250 characters long.
	//
	// Required: true
	AppName *string `json:"appName"`

	// app type
	// Required: true
	AppType AppType `json:"appType"`

	// An App maybe associated to many classifications.  A classification drives how the integration is presented
	// to the user in the SmartThings mobile clients.  These classifications include:
	// * AUTOMATION - Denotes an integration that should display under the "Automation" tab in mobile clients.
	// * SERVICE - Denotes an integration that is classified as a "Service".
	// * DEVICE - Denotes an integration that should display under the "Device" tab in mobile clients.
	// * CONNECTED_SERVICE - Denotes an integration that should display under the "Connected Services" menu in mobile clients.
	// * HIDDEN - Denotes an integration that should not display in mobile clients
	//
	// Required: true
	Classifications []AppClassification `json:"classifications"`

	// A default description for an app.
	//
	// Required: true
	// Max Length: 250
	Description *string `json:"description"`

	// A default display name for an app.
	//
	// Required: true
	// Max Length: 75
	DisplayName *string `json:"displayName"`

	// icon image
	IconImage *IconImage `json:"iconImage,omitempty"`

	// lambda smart app
	LambdaSmartApp *CreateOrUpdateLambdaSmartAppRequest `json:"lambdaSmartApp,omitempty"`

	// oauth
	Oauth *AppOAuth `json:"oauth,omitempty"`

	// principal type
	PrincipalType PrincipalType `json:"principalType,omitempty"`

	// Inform the installation systems that a particular app can only be installed once within a user's account.
	//
	SingleInstance *bool `json:"singleInstance,omitempty"`

	// ui
	UI *AppUISettings `json:"ui,omitempty"`

	// webhook smart app
	WebhookSmartApp *CreateOrUpdateWebhookSmartAppRequest `json:"webhookSmartApp,omitempty"`
}

// Validate validates this create app request
func (m *CreateAppRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIconImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLambdaSmartApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOauth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipalType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookSmartApp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAppRequest) validateAppName(formats strfmt.Registry) error {

	if err := validate.Required("appName", "body", m.AppName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateAppType(formats strfmt.Registry) error {

	if err := m.AppType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("appType")
		}
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateClassifications(formats strfmt.Registry) error {

	if err := validate.Required("classifications", "body", m.Classifications); err != nil {
		return err
	}

	for i := 0; i < len(m.Classifications); i++ {

		if err := m.Classifications[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classifications" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *CreateAppRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 250); err != nil {
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(*m.DisplayName), 75); err != nil {
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateIconImage(formats strfmt.Registry) error {

	if swag.IsZero(m.IconImage) { // not required
		return nil
	}

	if m.IconImage != nil {
		if err := m.IconImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iconImage")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAppRequest) validateLambdaSmartApp(formats strfmt.Registry) error {

	if swag.IsZero(m.LambdaSmartApp) { // not required
		return nil
	}

	if m.LambdaSmartApp != nil {
		if err := m.LambdaSmartApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lambdaSmartApp")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAppRequest) validateOauth(formats strfmt.Registry) error {

	if swag.IsZero(m.Oauth) { // not required
		return nil
	}

	if m.Oauth != nil {
		if err := m.Oauth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAppRequest) validatePrincipalType(formats strfmt.Registry) error {

	if swag.IsZero(m.PrincipalType) { // not required
		return nil
	}

	if err := m.PrincipalType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("principalType")
		}
		return err
	}

	return nil
}

func (m *CreateAppRequest) validateUI(formats strfmt.Registry) error {

	if swag.IsZero(m.UI) { // not required
		return nil
	}

	if m.UI != nil {
		if err := m.UI.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ui")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAppRequest) validateWebhookSmartApp(formats strfmt.Registry) error {

	if swag.IsZero(m.WebhookSmartApp) { // not required
		return nil
	}

	if m.WebhookSmartApp != nil {
		if err := m.WebhookSmartApp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhookSmartApp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAppRequest) UnmarshalBinary(b []byte) error {
	var res CreateAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
