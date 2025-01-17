// Code generated by go-swagger; DO NOT EDIT.

package installedapps

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

// NewCreateInstalledAppEventsParams creates a new CreateInstalledAppEventsParams object
// with the default values initialized.
func NewCreateInstalledAppEventsParams() *CreateInstalledAppEventsParams {
	var ()
	return &CreateInstalledAppEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstalledAppEventsParamsWithTimeout creates a new CreateInstalledAppEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInstalledAppEventsParamsWithTimeout(timeout time.Duration) *CreateInstalledAppEventsParams {
	var ()
	return &CreateInstalledAppEventsParams{

		timeout: timeout,
	}
}

// NewCreateInstalledAppEventsParamsWithContext creates a new CreateInstalledAppEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInstalledAppEventsParamsWithContext(ctx context.Context) *CreateInstalledAppEventsParams {
	var ()
	return &CreateInstalledAppEventsParams{

		Context: ctx,
	}
}

// NewCreateInstalledAppEventsParamsWithHTTPClient creates a new CreateInstalledAppEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInstalledAppEventsParamsWithHTTPClient(client *http.Client) *CreateInstalledAppEventsParams {
	var ()
	return &CreateInstalledAppEventsParams{
		HTTPClient: client,
	}
}

/*CreateInstalledAppEventsParams contains all the parameters to send to the API endpoint
for the create installed app events operation typically these are written to a http.Request
*/
type CreateInstalledAppEventsParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*CreateInstalledAppEventsRequest*/
	CreateInstalledAppEventsRequest *models.CreateInstalledAppEventsRequest
	/*InstalledAppID
	  The ID of the installed application.

	*/
	InstalledAppID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create installed app events params
func (o *CreateInstalledAppEventsParams) WithTimeout(timeout time.Duration) *CreateInstalledAppEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create installed app events params
func (o *CreateInstalledAppEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create installed app events params
func (o *CreateInstalledAppEventsParams) WithContext(ctx context.Context) *CreateInstalledAppEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create installed app events params
func (o *CreateInstalledAppEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create installed app events params
func (o *CreateInstalledAppEventsParams) WithHTTPClient(client *http.Client) *CreateInstalledAppEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create installed app events params
func (o *CreateInstalledAppEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create installed app events params
func (o *CreateInstalledAppEventsParams) WithAuthorization(authorization string) *CreateInstalledAppEventsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create installed app events params
func (o *CreateInstalledAppEventsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithCreateInstalledAppEventsRequest adds the createInstalledAppEventsRequest to the create installed app events params
func (o *CreateInstalledAppEventsParams) WithCreateInstalledAppEventsRequest(createInstalledAppEventsRequest *models.CreateInstalledAppEventsRequest) *CreateInstalledAppEventsParams {
	o.SetCreateInstalledAppEventsRequest(createInstalledAppEventsRequest)
	return o
}

// SetCreateInstalledAppEventsRequest adds the createInstalledAppEventsRequest to the create installed app events params
func (o *CreateInstalledAppEventsParams) SetCreateInstalledAppEventsRequest(createInstalledAppEventsRequest *models.CreateInstalledAppEventsRequest) {
	o.CreateInstalledAppEventsRequest = createInstalledAppEventsRequest
}

// WithInstalledAppID adds the installedAppID to the create installed app events params
func (o *CreateInstalledAppEventsParams) WithInstalledAppID(installedAppID string) *CreateInstalledAppEventsParams {
	o.SetInstalledAppID(installedAppID)
	return o
}

// SetInstalledAppID adds the installedAppId to the create installed app events params
func (o *CreateInstalledAppEventsParams) SetInstalledAppID(installedAppID string) {
	o.InstalledAppID = installedAppID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstalledAppEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.CreateInstalledAppEventsRequest != nil {
		if err := r.SetBodyParam(o.CreateInstalledAppEventsRequest); err != nil {
			return err
		}
	}

	// path param installedAppId
	if err := r.SetPathParam("installedAppId", o.InstalledAppID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
