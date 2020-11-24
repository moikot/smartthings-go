// Code generated by go-swagger; DO NOT EDIT.

package capabilities

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
)

// NewListCapabilitiesParams creates a new ListCapabilitiesParams object
// with the default values initialized.
func NewListCapabilitiesParams() *ListCapabilitiesParams {
	var ()
	return &ListCapabilitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCapabilitiesParamsWithTimeout creates a new ListCapabilitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCapabilitiesParamsWithTimeout(timeout time.Duration) *ListCapabilitiesParams {
	var ()
	return &ListCapabilitiesParams{

		timeout: timeout,
	}
}

// NewListCapabilitiesParamsWithContext creates a new ListCapabilitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCapabilitiesParamsWithContext(ctx context.Context) *ListCapabilitiesParams {
	var ()
	return &ListCapabilitiesParams{

		Context: ctx,
	}
}

// NewListCapabilitiesParamsWithHTTPClient creates a new ListCapabilitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCapabilitiesParamsWithHTTPClient(client *http.Client) *ListCapabilitiesParams {
	var ()
	return &ListCapabilitiesParams{
		HTTPClient: client,
	}
}

/*ListCapabilitiesParams contains all the parameters to send to the API endpoint
for the list capabilities operation typically these are written to a http.Request
*/
type ListCapabilitiesParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list capabilities params
func (o *ListCapabilitiesParams) WithTimeout(timeout time.Duration) *ListCapabilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list capabilities params
func (o *ListCapabilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list capabilities params
func (o *ListCapabilitiesParams) WithContext(ctx context.Context) *ListCapabilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list capabilities params
func (o *ListCapabilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list capabilities params
func (o *ListCapabilitiesParams) WithHTTPClient(client *http.Client) *ListCapabilitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list capabilities params
func (o *ListCapabilitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the list capabilities params
func (o *ListCapabilitiesParams) WithAuthorization(authorization string) *ListCapabilitiesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list capabilities params
func (o *ListCapabilitiesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WriteToRequest writes these params to a swagger request
func (o *ListCapabilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
