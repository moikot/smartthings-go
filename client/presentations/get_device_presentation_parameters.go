// Code generated by go-swagger; DO NOT EDIT.

package presentations

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

// NewGetDevicePresentationParams creates a new GetDevicePresentationParams object
// with the default values initialized.
func NewGetDevicePresentationParams() *GetDevicePresentationParams {
	var ()
	return &GetDevicePresentationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDevicePresentationParamsWithTimeout creates a new GetDevicePresentationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDevicePresentationParamsWithTimeout(timeout time.Duration) *GetDevicePresentationParams {
	var ()
	return &GetDevicePresentationParams{

		timeout: timeout,
	}
}

// NewGetDevicePresentationParamsWithContext creates a new GetDevicePresentationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDevicePresentationParamsWithContext(ctx context.Context) *GetDevicePresentationParams {
	var ()
	return &GetDevicePresentationParams{

		Context: ctx,
	}
}

// NewGetDevicePresentationParamsWithHTTPClient creates a new GetDevicePresentationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDevicePresentationParamsWithHTTPClient(client *http.Client) *GetDevicePresentationParams {
	var ()
	return &GetDevicePresentationParams{
		HTTPClient: client,
	}
}

/*GetDevicePresentationParams contains all the parameters to send to the API endpoint
for the get device presentation operation typically these are written to a http.Request
*/
type GetDevicePresentationParams struct {

	/*AcceptLanguage
	  Language header representing the client's preferred language. The format of the `Accept-Language` header follows what is defined in [RFC 7231, section 5.3.5](https://tools.ietf.org/html/rfc7231#section-5.3.5)

	*/
	AcceptLanguage *string
	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*IfNoneMatch
	  The ETag for the request.

	*/
	IfNoneMatch *string
	/*DeviceID
	  The ID of a device for which we want to load the device presentation.
	If the device ID is provided, no other fields are required.

	*/
	DeviceID *string
	/*ManufacturerName
	  Secondary namespacing key for grouping presentations (formerly `mnmn`)

	*/
	ManufacturerName *string
	/*PresentationID
	  System generated identifier that corresponds to a device presentation (formerly `vid`)

	*/
	PresentationID string
	/*View
	  view type

	*/
	View *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device presentation params
func (o *GetDevicePresentationParams) WithTimeout(timeout time.Duration) *GetDevicePresentationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device presentation params
func (o *GetDevicePresentationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device presentation params
func (o *GetDevicePresentationParams) WithContext(ctx context.Context) *GetDevicePresentationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device presentation params
func (o *GetDevicePresentationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device presentation params
func (o *GetDevicePresentationParams) WithHTTPClient(client *http.Client) *GetDevicePresentationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device presentation params
func (o *GetDevicePresentationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the get device presentation params
func (o *GetDevicePresentationParams) WithAcceptLanguage(acceptLanguage *string) *GetDevicePresentationParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the get device presentation params
func (o *GetDevicePresentationParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithAuthorization adds the authorization to the get device presentation params
func (o *GetDevicePresentationParams) WithAuthorization(authorization string) *GetDevicePresentationParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get device presentation params
func (o *GetDevicePresentationParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithIfNoneMatch adds the ifNoneMatch to the get device presentation params
func (o *GetDevicePresentationParams) WithIfNoneMatch(ifNoneMatch *string) *GetDevicePresentationParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get device presentation params
func (o *GetDevicePresentationParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDeviceID adds the deviceID to the get device presentation params
func (o *GetDevicePresentationParams) WithDeviceID(deviceID *string) *GetDevicePresentationParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device presentation params
func (o *GetDevicePresentationParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithManufacturerName adds the manufacturerName to the get device presentation params
func (o *GetDevicePresentationParams) WithManufacturerName(manufacturerName *string) *GetDevicePresentationParams {
	o.SetManufacturerName(manufacturerName)
	return o
}

// SetManufacturerName adds the manufacturerName to the get device presentation params
func (o *GetDevicePresentationParams) SetManufacturerName(manufacturerName *string) {
	o.ManufacturerName = manufacturerName
}

// WithPresentationID adds the presentationID to the get device presentation params
func (o *GetDevicePresentationParams) WithPresentationID(presentationID string) *GetDevicePresentationParams {
	o.SetPresentationID(presentationID)
	return o
}

// SetPresentationID adds the presentationId to the get device presentation params
func (o *GetDevicePresentationParams) SetPresentationID(presentationID string) {
	o.PresentationID = presentationID
}

// WithView adds the view to the get device presentation params
func (o *GetDevicePresentationParams) WithView(view *string) *GetDevicePresentationParams {
	o.SetView(view)
	return o
}

// SetView adds the view to the get device presentation params
func (o *GetDevicePresentationParams) SetView(view *string) {
	o.View = view
}

// WriteToRequest writes these params to a swagger request
func (o *GetDevicePresentationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}

	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}

	}

	if o.DeviceID != nil {

		// query param deviceId
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("deviceId", qDeviceID); err != nil {
				return err
			}
		}

	}

	if o.ManufacturerName != nil {

		// query param manufacturerName
		var qrManufacturerName string
		if o.ManufacturerName != nil {
			qrManufacturerName = *o.ManufacturerName
		}
		qManufacturerName := qrManufacturerName
		if qManufacturerName != "" {
			if err := r.SetQueryParam("manufacturerName", qManufacturerName); err != nil {
				return err
			}
		}

	}

	// query param presentationId
	qrPresentationID := o.PresentationID
	qPresentationID := qrPresentationID
	if qPresentationID != "" {
		if err := r.SetQueryParam("presentationId", qPresentationID); err != nil {
			return err
		}
	}

	if o.View != nil {

		// query param view
		var qrView string
		if o.View != nil {
			qrView = *o.View
		}
		qView := qrView
		if qView != "" {
			if err := r.SetQueryParam("view", qView); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}