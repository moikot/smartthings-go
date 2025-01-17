// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointApp endpoint app
//
// swagger:model EndpointApp
type EndpointApp struct {

	// The name of the endpoint app
	AppName string `json:"appName,omitempty"`

	// Possible values - "", "cst", "wwst", "review"
	CertificationStatus string `json:"certificationStatus,omitempty"`

	// Viper endpoint app id for the partner
	EndpointAppID string `json:"endpointAppId,omitempty"`

	// Possible values - "lambda" or "webhook"
	HostingType string `json:"hostingType,omitempty"`

	// url of partner icon
	Icon string `json:"icon,omitempty"`

	// url of partner icon in 2x dimensions
	Icon2x string `json:"icon2x,omitempty"`

	// url of partner icon in 3x dimensions
	Icon3x string `json:"icon3x,omitempty"`

	// lambda arn of the partner for US region (default)
	LambdaArn string `json:"lambdaArn,omitempty"`

	// lambda arn of the partner for AP region
	LambdaArnAP string `json:"lambdaArnAP,omitempty"`

	// lambda arn of the partner for CN region
	LambdaArnCN string `json:"lambdaArnCN,omitempty"`

	// lambda arn of the partner for EU region
	LambdaArnEU string `json:"lambdaArnEU,omitempty"`

	// oAuth authorization url of the partner
	OAuthAuthorizationURL string `json:"oAuthAuthorizationUrl,omitempty"`

	// Client id for the partner oAuth
	OAuthClientID string `json:"oAuthClientId,omitempty"`

	// Client secret for the partner oAuth
	OAuthClientSecret string `json:"oAuthClientSecret,omitempty"`

	// oAuth scope for the partner. Example "remote_control:all" for Lifx
	OAuthScope string `json:"oAuthScope,omitempty"`

	// oAuth token refresh url of the partner
	OAuthTokenURL string `json:"oAuthTokenUrl,omitempty"`

	// The name of the partner/brand
	PartnerName string `json:"partnerName,omitempty"`

	// Possible values - "alexa-schema", "st-schema", "google-schema"
	SchemaType string `json:"schemaType,omitempty"`

	// Email for the partner
	UserEmail string `json:"userEmail,omitempty"`

	// user id for the partner
	UserID string `json:"userId,omitempty"`

	// webhook url for the partner
	WebhookURL string `json:"webhookUrl,omitempty"`
}

// Validate validates this endpoint app
func (m *EndpointApp) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointApp) UnmarshalBinary(b []byte) error {
	var res EndpointApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
