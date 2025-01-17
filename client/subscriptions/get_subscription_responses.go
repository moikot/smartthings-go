// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSubscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSubscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSubscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSubscriptionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/*GetSubscriptionOK handles this case with default header values.

The subscription
*/
type GetSubscriptionOK struct {
	Payload *models.Subscription
}

func (o *GetSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions/{subscriptionId}][%d] getSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionOK) GetPayload() *models.Subscription {
	return o.Payload
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionBadRequest creates a GetSubscriptionBadRequest with default headers values
func NewGetSubscriptionBadRequest() *GetSubscriptionBadRequest {
	return &GetSubscriptionBadRequest{}
}

/*GetSubscriptionBadRequest handles this case with default header values.

Bad request
*/
type GetSubscriptionBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions/{subscriptionId}][%d] getSubscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *GetSubscriptionBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionUnauthorized creates a GetSubscriptionUnauthorized with default headers values
func NewGetSubscriptionUnauthorized() *GetSubscriptionUnauthorized {
	return &GetSubscriptionUnauthorized{}
}

/*GetSubscriptionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetSubscriptionUnauthorized struct {
}

func (o *GetSubscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions/{subscriptionId}][%d] getSubscriptionUnauthorized ", 401)
}

func (o *GetSubscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionForbidden creates a GetSubscriptionForbidden with default headers values
func NewGetSubscriptionForbidden() *GetSubscriptionForbidden {
	return &GetSubscriptionForbidden{}
}

/*GetSubscriptionForbidden handles this case with default header values.

Forbidden
*/
type GetSubscriptionForbidden struct {
}

func (o *GetSubscriptionForbidden) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions/{subscriptionId}][%d] getSubscriptionForbidden ", 403)
}

func (o *GetSubscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionTooManyRequests creates a GetSubscriptionTooManyRequests with default headers values
func NewGetSubscriptionTooManyRequests() *GetSubscriptionTooManyRequests {
	return &GetSubscriptionTooManyRequests{}
}

/*GetSubscriptionTooManyRequests handles this case with default header values.

Too many requests
*/
type GetSubscriptionTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetSubscriptionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions/{subscriptionId}][%d] getSubscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSubscriptionTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSubscriptionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionDefault creates a GetSubscriptionDefault with default headers values
func NewGetSubscriptionDefault(code int) *GetSubscriptionDefault {
	return &GetSubscriptionDefault{
		_statusCode: code,
	}
}

/*GetSubscriptionDefault handles this case with default header values.

Unexpected error
*/
type GetSubscriptionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get subscription default response
func (o *GetSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetSubscriptionDefault) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}/subscriptions/{subscriptionId}][%d] getSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *GetSubscriptionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
