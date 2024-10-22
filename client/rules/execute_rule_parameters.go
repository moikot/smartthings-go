// Code generated by go-swagger; DO NOT EDIT.

package rules

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

// NewExecuteRuleParams creates a new ExecuteRuleParams object
// with the default values initialized.
func NewExecuteRuleParams() *ExecuteRuleParams {
	var ()
	return &ExecuteRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteRuleParamsWithTimeout creates a new ExecuteRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExecuteRuleParamsWithTimeout(timeout time.Duration) *ExecuteRuleParams {
	var ()
	return &ExecuteRuleParams{

		timeout: timeout,
	}
}

// NewExecuteRuleParamsWithContext creates a new ExecuteRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewExecuteRuleParamsWithContext(ctx context.Context) *ExecuteRuleParams {
	var ()
	return &ExecuteRuleParams{

		Context: ctx,
	}
}

// NewExecuteRuleParamsWithHTTPClient creates a new ExecuteRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExecuteRuleParamsWithHTTPClient(client *http.Client) *ExecuteRuleParams {
	var ()
	return &ExecuteRuleParams{
		HTTPClient: client,
	}
}

/*ExecuteRuleParams contains all the parameters to send to the API endpoint
for the execute rule operation typically these are written to a http.Request
*/
type ExecuteRuleParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*LocationID
	  The ID of the location that both the installed smart app and source are associated with.

	*/
	LocationID string
	/*RuleID
	  The rule ID

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the execute rule params
func (o *ExecuteRuleParams) WithTimeout(timeout time.Duration) *ExecuteRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute rule params
func (o *ExecuteRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute rule params
func (o *ExecuteRuleParams) WithContext(ctx context.Context) *ExecuteRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute rule params
func (o *ExecuteRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute rule params
func (o *ExecuteRuleParams) WithHTTPClient(client *http.Client) *ExecuteRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute rule params
func (o *ExecuteRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the execute rule params
func (o *ExecuteRuleParams) WithAuthorization(authorization string) *ExecuteRuleParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the execute rule params
func (o *ExecuteRuleParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocationID adds the locationID to the execute rule params
func (o *ExecuteRuleParams) WithLocationID(locationID string) *ExecuteRuleParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the execute rule params
func (o *ExecuteRuleParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithRuleID adds the ruleID to the execute rule params
func (o *ExecuteRuleParams) WithRuleID(ruleID string) *ExecuteRuleParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the execute rule params
func (o *ExecuteRuleParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param locationId
	qrLocationID := o.LocationID
	qLocationID := qrLocationID
	if qLocationID != "" {
		if err := r.SetQueryParam("locationId", qLocationID); err != nil {
			return err
		}
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
