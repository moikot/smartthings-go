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
)

// NewListInstallationsParams creates a new ListInstallationsParams object
// with the default values initialized.
func NewListInstallationsParams() *ListInstallationsParams {
	var ()
	return &ListInstallationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListInstallationsParamsWithTimeout creates a new ListInstallationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListInstallationsParamsWithTimeout(timeout time.Duration) *ListInstallationsParams {
	var ()
	return &ListInstallationsParams{

		timeout: timeout,
	}
}

// NewListInstallationsParamsWithContext creates a new ListInstallationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListInstallationsParamsWithContext(ctx context.Context) *ListInstallationsParams {
	var ()
	return &ListInstallationsParams{

		Context: ctx,
	}
}

// NewListInstallationsParamsWithHTTPClient creates a new ListInstallationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListInstallationsParamsWithHTTPClient(client *http.Client) *ListInstallationsParams {
	var ()
	return &ListInstallationsParams{
		HTTPClient: client,
	}
}

/*ListInstallationsParams contains all the parameters to send to the API endpoint
for the list installations operation typically these are written to a http.Request
*/
type ListInstallationsParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*AppID
	  The ID of an App

	*/
	AppID *string
	/*DeviceID
	  The ID of the device

	*/
	DeviceID *string
	/*InstalledAppStatus
	  State of the Installed App.

	*/
	InstalledAppStatus *string
	/*InstalledAppType
	  Denotes the type of installed app.

	*/
	InstalledAppType *string
	/*LocationID
	  The ID of the location that both the installed smart app and source are associated with.

	*/
	LocationID *string
	/*ModeID
	  The ID of the mode

	*/
	ModeID *string
	/*Tag
	  May be used to filter a resource by it's assigned user-tags.  Multiple tag query params are automatically joined with OR.

	Example usage in query string:
	```
	?tag:key_name=value1&tag:key_name=value2
	```


	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list installations params
func (o *ListInstallationsParams) WithTimeout(timeout time.Duration) *ListInstallationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list installations params
func (o *ListInstallationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list installations params
func (o *ListInstallationsParams) WithContext(ctx context.Context) *ListInstallationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list installations params
func (o *ListInstallationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list installations params
func (o *ListInstallationsParams) WithHTTPClient(client *http.Client) *ListInstallationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list installations params
func (o *ListInstallationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the list installations params
func (o *ListInstallationsParams) WithAuthorization(authorization string) *ListInstallationsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the list installations params
func (o *ListInstallationsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAppID adds the appID to the list installations params
func (o *ListInstallationsParams) WithAppID(appID *string) *ListInstallationsParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the list installations params
func (o *ListInstallationsParams) SetAppID(appID *string) {
	o.AppID = appID
}

// WithDeviceID adds the deviceID to the list installations params
func (o *ListInstallationsParams) WithDeviceID(deviceID *string) *ListInstallationsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the list installations params
func (o *ListInstallationsParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithInstalledAppStatus adds the installedAppStatus to the list installations params
func (o *ListInstallationsParams) WithInstalledAppStatus(installedAppStatus *string) *ListInstallationsParams {
	o.SetInstalledAppStatus(installedAppStatus)
	return o
}

// SetInstalledAppStatus adds the installedAppStatus to the list installations params
func (o *ListInstallationsParams) SetInstalledAppStatus(installedAppStatus *string) {
	o.InstalledAppStatus = installedAppStatus
}

// WithInstalledAppType adds the installedAppType to the list installations params
func (o *ListInstallationsParams) WithInstalledAppType(installedAppType *string) *ListInstallationsParams {
	o.SetInstalledAppType(installedAppType)
	return o
}

// SetInstalledAppType adds the installedAppType to the list installations params
func (o *ListInstallationsParams) SetInstalledAppType(installedAppType *string) {
	o.InstalledAppType = installedAppType
}

// WithLocationID adds the locationID to the list installations params
func (o *ListInstallationsParams) WithLocationID(locationID *string) *ListInstallationsParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the list installations params
func (o *ListInstallationsParams) SetLocationID(locationID *string) {
	o.LocationID = locationID
}

// WithModeID adds the modeID to the list installations params
func (o *ListInstallationsParams) WithModeID(modeID *string) *ListInstallationsParams {
	o.SetModeID(modeID)
	return o
}

// SetModeID adds the modeId to the list installations params
func (o *ListInstallationsParams) SetModeID(modeID *string) {
	o.ModeID = modeID
}

// WithTag adds the tag to the list installations params
func (o *ListInstallationsParams) WithTag(tag *string) *ListInstallationsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the list installations params
func (o *ListInstallationsParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *ListInstallationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.AppID != nil {

		// query param appId
		var qrAppID string
		if o.AppID != nil {
			qrAppID = *o.AppID
		}
		qAppID := qrAppID
		if qAppID != "" {
			if err := r.SetQueryParam("appId", qAppID); err != nil {
				return err
			}
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

	if o.InstalledAppStatus != nil {

		// query param installedAppStatus
		var qrInstalledAppStatus string
		if o.InstalledAppStatus != nil {
			qrInstalledAppStatus = *o.InstalledAppStatus
		}
		qInstalledAppStatus := qrInstalledAppStatus
		if qInstalledAppStatus != "" {
			if err := r.SetQueryParam("installedAppStatus", qInstalledAppStatus); err != nil {
				return err
			}
		}

	}

	if o.InstalledAppType != nil {

		// query param installedAppType
		var qrInstalledAppType string
		if o.InstalledAppType != nil {
			qrInstalledAppType = *o.InstalledAppType
		}
		qInstalledAppType := qrInstalledAppType
		if qInstalledAppType != "" {
			if err := r.SetQueryParam("installedAppType", qInstalledAppType); err != nil {
				return err
			}
		}

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

	if o.ModeID != nil {

		// query param modeId
		var qrModeID string
		if o.ModeID != nil {
			qrModeID = *o.ModeID
		}
		qModeID := qrModeID
		if qModeID != "" {
			if err := r.SetQueryParam("modeId", qModeID); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
