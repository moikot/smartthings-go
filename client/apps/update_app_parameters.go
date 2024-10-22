// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/moikot/smartthings-go/models"
)

// NewUpdateAppParams creates a new UpdateAppParams object
// with the default values initialized.
func NewUpdateAppParams() *UpdateAppParams {
	var ()
	return &UpdateAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppParamsWithTimeout creates a new UpdateAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppParamsWithTimeout(timeout time.Duration) *UpdateAppParams {
	var ()
	return &UpdateAppParams{

		timeout: timeout,
	}
}

// NewUpdateAppParamsWithContext creates a new UpdateAppParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAppParamsWithContext(ctx context.Context) *UpdateAppParams {
	var ()
	return &UpdateAppParams{

		Context: ctx,
	}
}

// NewUpdateAppParamsWithHTTPClient creates a new UpdateAppParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAppParamsWithHTTPClient(client *http.Client) *UpdateAppParams {
	var ()
	return &UpdateAppParams{
		HTTPClient: client,
	}
}

/*UpdateAppParams contains all the parameters to send to the API endpoint
for the update app operation typically these are written to a http.Request
*/
type UpdateAppParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*AppNameOrID
	  The appName or appId field of an app.

	*/
	AppNameOrID string
	/*RequireConfirmation
	  Override default configuration to use either PING or CONFIRMATION lifecycle. For WEBHOOK_SMART_APP only.


	*/
	RequireConfirmation *bool
	/*SignatureType
	  The Signature Type of the application. For WEBHOOK_SMART_APP only.


	*/
	SignatureType *string
	/*UpdateAppRequest*/
	UpdateAppRequest *models.UpdateAppRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update app params
func (o *UpdateAppParams) WithTimeout(timeout time.Duration) *UpdateAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app params
func (o *UpdateAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app params
func (o *UpdateAppParams) WithContext(ctx context.Context) *UpdateAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app params
func (o *UpdateAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app params
func (o *UpdateAppParams) WithHTTPClient(client *http.Client) *UpdateAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app params
func (o *UpdateAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update app params
func (o *UpdateAppParams) WithAuthorization(authorization string) *UpdateAppParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update app params
func (o *UpdateAppParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAppNameOrID adds the appNameOrID to the update app params
func (o *UpdateAppParams) WithAppNameOrID(appNameOrID string) *UpdateAppParams {
	o.SetAppNameOrID(appNameOrID)
	return o
}

// SetAppNameOrID adds the appNameOrId to the update app params
func (o *UpdateAppParams) SetAppNameOrID(appNameOrID string) {
	o.AppNameOrID = appNameOrID
}

// WithRequireConfirmation adds the requireConfirmation to the update app params
func (o *UpdateAppParams) WithRequireConfirmation(requireConfirmation *bool) *UpdateAppParams {
	o.SetRequireConfirmation(requireConfirmation)
	return o
}

// SetRequireConfirmation adds the requireConfirmation to the update app params
func (o *UpdateAppParams) SetRequireConfirmation(requireConfirmation *bool) {
	o.RequireConfirmation = requireConfirmation
}

// WithSignatureType adds the signatureType to the update app params
func (o *UpdateAppParams) WithSignatureType(signatureType *string) *UpdateAppParams {
	o.SetSignatureType(signatureType)
	return o
}

// SetSignatureType adds the signatureType to the update app params
func (o *UpdateAppParams) SetSignatureType(signatureType *string) {
	o.SignatureType = signatureType
}

// WithUpdateAppRequest adds the updateAppRequest to the update app params
func (o *UpdateAppParams) WithUpdateAppRequest(updateAppRequest *models.UpdateAppRequest) *UpdateAppParams {
	o.SetUpdateAppRequest(updateAppRequest)
	return o
}

// SetUpdateAppRequest adds the updateAppRequest to the update app params
func (o *UpdateAppParams) SetUpdateAppRequest(updateAppRequest *models.UpdateAppRequest) {
	o.UpdateAppRequest = updateAppRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param appNameOrId
	if err := r.SetPathParam("appNameOrId", o.AppNameOrID); err != nil {
		return err
	}

	if o.RequireConfirmation != nil {

		// query param requireConfirmation
		var qrRequireConfirmation bool
		if o.RequireConfirmation != nil {
			qrRequireConfirmation = *o.RequireConfirmation
		}
		qRequireConfirmation := swag.FormatBool(qrRequireConfirmation)
		if qRequireConfirmation != "" {
			if err := r.SetQueryParam("requireConfirmation", qRequireConfirmation); err != nil {
				return err
			}
		}

	}

	if o.SignatureType != nil {

		// query param signatureType
		var qrSignatureType string
		if o.SignatureType != nil {
			qrSignatureType = *o.SignatureType
		}
		qSignatureType := qrSignatureType
		if qSignatureType != "" {
			if err := r.SetQueryParam("signatureType", qSignatureType); err != nil {
				return err
			}
		}

	}

	if o.UpdateAppRequest != nil {
		if err := r.SetBodyParam(o.UpdateAppRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
