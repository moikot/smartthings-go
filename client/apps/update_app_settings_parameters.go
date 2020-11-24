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

	"github.com/moikot/smartthings-go/models"
)

// NewUpdateAppSettingsParams creates a new UpdateAppSettingsParams object
// with the default values initialized.
func NewUpdateAppSettingsParams() *UpdateAppSettingsParams {
	var ()
	return &UpdateAppSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppSettingsParamsWithTimeout creates a new UpdateAppSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppSettingsParamsWithTimeout(timeout time.Duration) *UpdateAppSettingsParams {
	var ()
	return &UpdateAppSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateAppSettingsParamsWithContext creates a new UpdateAppSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAppSettingsParamsWithContext(ctx context.Context) *UpdateAppSettingsParams {
	var ()
	return &UpdateAppSettingsParams{

		Context: ctx,
	}
}

// NewUpdateAppSettingsParamsWithHTTPClient creates a new UpdateAppSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAppSettingsParamsWithHTTPClient(client *http.Client) *UpdateAppSettingsParams {
	var ()
	return &UpdateAppSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateAppSettingsParams contains all the parameters to send to the API endpoint
for the update app settings operation typically these are written to a http.Request
*/
type UpdateAppSettingsParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*AppNameOrID
	  The appName or appId  field of an app.

	*/
	AppNameOrID string
	/*UpdateAppSettingsRequest*/
	UpdateAppSettingsRequest *models.UpdateAppSettingsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update app settings params
func (o *UpdateAppSettingsParams) WithTimeout(timeout time.Duration) *UpdateAppSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app settings params
func (o *UpdateAppSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app settings params
func (o *UpdateAppSettingsParams) WithContext(ctx context.Context) *UpdateAppSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app settings params
func (o *UpdateAppSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app settings params
func (o *UpdateAppSettingsParams) WithHTTPClient(client *http.Client) *UpdateAppSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app settings params
func (o *UpdateAppSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update app settings params
func (o *UpdateAppSettingsParams) WithAuthorization(authorization string) *UpdateAppSettingsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update app settings params
func (o *UpdateAppSettingsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAppNameOrID adds the appNameOrID to the update app settings params
func (o *UpdateAppSettingsParams) WithAppNameOrID(appNameOrID string) *UpdateAppSettingsParams {
	o.SetAppNameOrID(appNameOrID)
	return o
}

// SetAppNameOrID adds the appNameOrId to the update app settings params
func (o *UpdateAppSettingsParams) SetAppNameOrID(appNameOrID string) {
	o.AppNameOrID = appNameOrID
}

// WithUpdateAppSettingsRequest adds the updateAppSettingsRequest to the update app settings params
func (o *UpdateAppSettingsParams) WithUpdateAppSettingsRequest(updateAppSettingsRequest *models.UpdateAppSettingsRequest) *UpdateAppSettingsParams {
	o.SetUpdateAppSettingsRequest(updateAppSettingsRequest)
	return o
}

// SetUpdateAppSettingsRequest adds the updateAppSettingsRequest to the update app settings params
func (o *UpdateAppSettingsParams) SetUpdateAppSettingsRequest(updateAppSettingsRequest *models.UpdateAppSettingsRequest) {
	o.UpdateAppSettingsRequest = updateAppSettingsRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UpdateAppSettingsRequest != nil {
		if err := r.SetBodyParam(o.UpdateAppSettingsRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
