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

	"github.com/moikot/smartthings-go/models"
)

// NewUpdateRoomParams creates a new UpdateRoomParams object
// with the default values initialized.
func NewUpdateRoomParams() *UpdateRoomParams {
	var ()
	return &UpdateRoomParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoomParamsWithTimeout creates a new UpdateRoomParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRoomParamsWithTimeout(timeout time.Duration) *UpdateRoomParams {
	var ()
	return &UpdateRoomParams{

		timeout: timeout,
	}
}

// NewUpdateRoomParamsWithContext creates a new UpdateRoomParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRoomParamsWithContext(ctx context.Context) *UpdateRoomParams {
	var ()
	return &UpdateRoomParams{

		Context: ctx,
	}
}

// NewUpdateRoomParamsWithHTTPClient creates a new UpdateRoomParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRoomParamsWithHTTPClient(client *http.Client) *UpdateRoomParams {
	var ()
	return &UpdateRoomParams{
		HTTPClient: client,
	}
}

/*UpdateRoomParams contains all the parameters to send to the API endpoint
for the update room operation typically these are written to a http.Request
*/
type UpdateRoomParams struct {

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
	/*UpdateRoomRequest*/
	UpdateRoomRequest *models.UpdateRoomRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update room params
func (o *UpdateRoomParams) WithTimeout(timeout time.Duration) *UpdateRoomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update room params
func (o *UpdateRoomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update room params
func (o *UpdateRoomParams) WithContext(ctx context.Context) *UpdateRoomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update room params
func (o *UpdateRoomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update room params
func (o *UpdateRoomParams) WithHTTPClient(client *http.Client) *UpdateRoomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update room params
func (o *UpdateRoomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update room params
func (o *UpdateRoomParams) WithAuthorization(authorization string) *UpdateRoomParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update room params
func (o *UpdateRoomParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocationID adds the locationID to the update room params
func (o *UpdateRoomParams) WithLocationID(locationID string) *UpdateRoomParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the update room params
func (o *UpdateRoomParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithRoomID adds the roomID to the update room params
func (o *UpdateRoomParams) WithRoomID(roomID string) *UpdateRoomParams {
	o.SetRoomID(roomID)
	return o
}

// SetRoomID adds the roomId to the update room params
func (o *UpdateRoomParams) SetRoomID(roomID string) {
	o.RoomID = roomID
}

// WithUpdateRoomRequest adds the updateRoomRequest to the update room params
func (o *UpdateRoomParams) WithUpdateRoomRequest(updateRoomRequest *models.UpdateRoomRequest) *UpdateRoomParams {
	o.SetUpdateRoomRequest(updateRoomRequest)
	return o
}

// SetUpdateRoomRequest adds the updateRoomRequest to the update room params
func (o *UpdateRoomParams) SetUpdateRoomRequest(updateRoomRequest *models.UpdateRoomRequest) {
	o.UpdateRoomRequest = updateRoomRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.UpdateRoomRequest != nil {
		if err := r.SetBodyParam(o.UpdateRoomRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
