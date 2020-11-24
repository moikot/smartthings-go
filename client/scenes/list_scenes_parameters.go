// Code generated by go-swagger; DO NOT EDIT.

package scenes

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

// NewListScenesParams creates a new ListScenesParams object
// with the default values initialized.
func NewListScenesParams() *ListScenesParams {
	var ()
	return &ListScenesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListScenesParamsWithTimeout creates a new ListScenesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListScenesParamsWithTimeout(timeout time.Duration) *ListScenesParams {
	var ()
	return &ListScenesParams{

		timeout: timeout,
	}
}

// NewListScenesParamsWithContext creates a new ListScenesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListScenesParamsWithContext(ctx context.Context) *ListScenesParams {
	var ()
	return &ListScenesParams{

		Context: ctx,
	}
}

// NewListScenesParamsWithHTTPClient creates a new ListScenesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListScenesParamsWithHTTPClient(client *http.Client) *ListScenesParams {
	var ()
	return &ListScenesParams{
		HTTPClient: client,
	}
}

/*ListScenesParams contains all the parameters to send to the API endpoint
for the list scenes operation typically these are written to a http.Request
*/
type ListScenesParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*LocationID
	  The location of a scene.

	*/
	LocationID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list scenes params
func (o *ListScenesParams) WithTimeout(timeout time.Duration) *ListScenesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list scenes params
func (o *ListScenesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list scenes params
func (o *ListScenesParams) WithContext(ctx context.Context) *ListScenesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list scenes params
func (o *ListScenesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list scenes params
func (o *ListScenesParams) WithHTTPClient(client *http.Client) *ListScenesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list scenes params
func (o *ListScenesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the list scenes params
func (o *ListScenesParams) WithAuthorization(authorization string) *ListScenesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list scenes params
func (o *ListScenesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocationID adds the locationID to the list scenes params
func (o *ListScenesParams) WithLocationID(locationID *string) *ListScenesParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list scenes params
func (o *ListScenesParams) SetLocationID(locationID *string) {
	o.LocationID = locationID
}

// WriteToRequest writes these params to a swagger request
func (o *ListScenesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.LocationID != nil {

		// query param locationId
		var qrLocationID string
		if o.LocationID != nil {
			qrLocationID = *o.LocationID
		}
		qLocationID := qrLocationID
		if qLocationID != "" {
			if err := r.SetQueryParam("locationId", qLocationID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
