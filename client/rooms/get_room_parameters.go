// Code generated by go-swagger; DO NOT EDIT.

package rooms

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

// NewGetRoomParams creates a new GetRoomParams object
// with the default values initialized.
func NewGetRoomParams() *GetRoomParams {
	var ()
	return &GetRoomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoomParamsWithTimeout creates a new GetRoomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoomParamsWithTimeout(timeout time.Duration) *GetRoomParams {
	var ()
	return &GetRoomParams{

		timeout: timeout,
	}
}

// NewGetRoomParamsWithContext creates a new GetRoomParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoomParamsWithContext(ctx context.Context) *GetRoomParams {
	var ()
	return &GetRoomParams{

		Context: ctx,
	}
}

// NewGetRoomParamsWithHTTPClient creates a new GetRoomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoomParamsWithHTTPClient(client *http.Client) *GetRoomParams {
	var ()
	return &GetRoomParams{
		HTTPClient: client,
	}
}

/*GetRoomParams contains all the parameters to send to the API endpoint
for the get room operation typically these are written to a http.Request
*/
type GetRoomParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*LocationID
	  The ID of the location.

	*/
	LocationID string
	/*RoomID
	  The ID of the room.

	*/
	RoomID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get room params
func (o *GetRoomParams) WithTimeout(timeout time.Duration) *GetRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get room params
func (o *GetRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get room params
func (o *GetRoomParams) WithContext(ctx context.Context) *GetRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get room params
func (o *GetRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get room params
func (o *GetRoomParams) WithHTTPClient(client *http.Client) *GetRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get room params
func (o *GetRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get room params
func (o *GetRoomParams) WithAuthorization(authorization string) *GetRoomParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get room params
func (o *GetRoomParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocationID adds the locationID to the get room params
func (o *GetRoomParams) WithLocationID(locationID string) *GetRoomParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the get room params
func (o *GetRoomParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithRoomID adds the roomID to the get room params
func (o *GetRoomParams) WithRoomID(roomID string) *GetRoomParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the get room params
func (o *GetRoomParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param roomId
	if err := r.SetPathParam("roomId", o.RoomID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
