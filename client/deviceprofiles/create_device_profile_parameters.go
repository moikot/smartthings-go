// Code generated by go-swagger; DO NOT EDIT.

package deviceprofiles

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

// NewCreateDeviceProfileParams creates a new CreateDeviceProfileParams object
// with the default values initialized.
func NewCreateDeviceProfileParams() *CreateDeviceProfileParams {
	var ()
	return &CreateDeviceProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDeviceProfileParamsWithTimeout creates a new CreateDeviceProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDeviceProfileParamsWithTimeout(timeout time.Duration) *CreateDeviceProfileParams {
	var ()
	return &CreateDeviceProfileParams{

		timeout: timeout,
	}
}

// NewCreateDeviceProfileParamsWithContext creates a new CreateDeviceProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDeviceProfileParamsWithContext(ctx context.Context) *CreateDeviceProfileParams {
	var ()
	return &CreateDeviceProfileParams{

		Context: ctx,
	}
}

// NewCreateDeviceProfileParamsWithHTTPClient creates a new CreateDeviceProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDeviceProfileParamsWithHTTPClient(client *http.Client) *CreateDeviceProfileParams {
	var ()
	return &CreateDeviceProfileParams{
		HTTPClient: client,
	}
}

/*CreateDeviceProfileParams contains all the parameters to send to the API endpoint
for the create device profile operation typically these are written to a http.Request
*/
type CreateDeviceProfileParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*Request
	  The device profile to be created.

	*/
	Request *models.CreateDeviceProfileRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create device profile params
func (o *CreateDeviceProfileParams) WithTimeout(timeout time.Duration) *CreateDeviceProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create device profile params
func (o *CreateDeviceProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create device profile params
func (o *CreateDeviceProfileParams) WithContext(ctx context.Context) *CreateDeviceProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create device profile params
func (o *CreateDeviceProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create device profile params
func (o *CreateDeviceProfileParams) WithHTTPClient(client *http.Client) *CreateDeviceProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create device profile params
func (o *CreateDeviceProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create device profile params
func (o *CreateDeviceProfileParams) WithAuthorization(authorization string) *CreateDeviceProfileParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create device profile params
func (o *CreateDeviceProfileParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithRequest adds the request to the create device profile params
func (o *CreateDeviceProfileParams) WithRequest(request *models.CreateDeviceProfileRequest) *CreateDeviceProfileParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create device profile params
func (o *CreateDeviceProfileParams) SetRequest(request *models.CreateDeviceProfileRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDeviceProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
