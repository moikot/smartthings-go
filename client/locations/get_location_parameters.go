// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// NewGetLocationParams creates a new GetLocationParams object
// with the default values initialized.
func NewGetLocationParams() *GetLocationParams {
	var ()
	return &GetLocationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLocationParamsWithTimeout creates a new GetLocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLocationParamsWithTimeout(timeout time.Duration) *GetLocationParams {
	var ()
	return &GetLocationParams{

		timeout: timeout,
	}
}

// NewGetLocationParamsWithContext creates a new GetLocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLocationParamsWithContext(ctx context.Context) *GetLocationParams {
	var ()
	return &GetLocationParams{

		Context: ctx,
	}
}

// NewGetLocationParamsWithHTTPClient creates a new GetLocationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLocationParamsWithHTTPClient(client *http.Client) *GetLocationParams {
	var ()
	return &GetLocationParams{
		HTTPClient: client,
	}
}

/*GetLocationParams contains all the parameters to send to the API endpoint
for the get location operation typically these are written to a http.Request
*/
type GetLocationParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*LocationID
	  The ID of the location.

	*/
	LocationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get location params
func (o *GetLocationParams) WithTimeout(timeout time.Duration) *GetLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get location params
func (o *GetLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get location params
func (o *GetLocationParams) WithContext(ctx context.Context) *GetLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get location params
func (o *GetLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get location params
func (o *GetLocationParams) WithHTTPClient(client *http.Client) *GetLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get location params
func (o *GetLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get location params
func (o *GetLocationParams) WithAuthorization(authorization string) *GetLocationParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get location params
func (o *GetLocationParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocationID adds the locationID to the get location params
func (o *GetLocationParams) WithLocationID(locationID string) *GetLocationParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the get location params
func (o *GetLocationParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
