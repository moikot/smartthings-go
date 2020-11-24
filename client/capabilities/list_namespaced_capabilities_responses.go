// Code generated by go-swagger; DO NOT EDIT.

package capabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/moikot/smartthings-go/models"
)

// ListNamespacedCapabilitiesReader is a Reader for the ListNamespacedCapabilities structure.
type ListNamespacedCapabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNamespacedCapabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNamespacedCapabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNamespacedCapabilitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListNamespacedCapabilitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListNamespacedCapabilitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListNamespacedCapabilitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNamespacedCapabilitiesOK creates a ListNamespacedCapabilitiesOK with default headers values
func NewListNamespacedCapabilitiesOK() *ListNamespacedCapabilitiesOK {
	return &ListNamespacedCapabilitiesOK{}
}

/*ListNamespacedCapabilitiesOK handles this case with default header values.

successful return of a namespace's capabilities
*/
type ListNamespacedCapabilitiesOK struct {
	/*Maximum requests allowed within the rate limit window.
	 */
	XRateLimitLimit int64
	/*Remaining requests available within the window.
	 */
	XRateLimitRemaining int64
	/*Time in milliseconds until the current window expires.
	 */
	XRateLimitReset int64

	Payload *models.PagedCapabilities
}

func (o *ListNamespacedCapabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /capabilities/namespaces/{namespace}][%d] listNamespacedCapabilitiesOK  %+v", 200, o.Payload)
}

func (o *ListNamespacedCapabilitiesOK) GetPayload() *models.PagedCapabilities {
	return o.Payload
}

func (o *ListNamespacedCapabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-RateLimit-Limit
	xRateLimitLimit, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Limit"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Limit", "header", "int64", response.GetHeader("X-RateLimit-Limit"))
	}
	o.XRateLimitLimit = xRateLimitLimit

	// response header X-RateLimit-Remaining
	xRateLimitRemaining, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Remaining"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", response.GetHeader("X-RateLimit-Remaining"))
	}
	o.XRateLimitRemaining = xRateLimitRemaining

	// response header X-RateLimit-Reset
	xRateLimitReset, err := swag.ConvertInt64(response.GetHeader("X-RateLimit-Reset"))
	if err != nil {
		return errors.InvalidType("X-RateLimit-Reset", "header", "int64", response.GetHeader("X-RateLimit-Reset"))
	}
	o.XRateLimitReset = xRateLimitReset

	o.Payload = new(models.PagedCapabilities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNamespacedCapabilitiesUnauthorized creates a ListNamespacedCapabilitiesUnauthorized with default headers values
func NewListNamespacedCapabilitiesUnauthorized() *ListNamespacedCapabilitiesUnauthorized {
	return &ListNamespacedCapabilitiesUnauthorized{}
}

/*ListNamespacedCapabilitiesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListNamespacedCapabilitiesUnauthorized struct {
}

func (o *ListNamespacedCapabilitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /capabilities/namespaces/{namespace}][%d] listNamespacedCapabilitiesUnauthorized ", 401)
}

func (o *ListNamespacedCapabilitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNamespacedCapabilitiesForbidden creates a ListNamespacedCapabilitiesForbidden with default headers values
func NewListNamespacedCapabilitiesForbidden() *ListNamespacedCapabilitiesForbidden {
	return &ListNamespacedCapabilitiesForbidden{}
}

/*ListNamespacedCapabilitiesForbidden handles this case with default header values.

Forbidden
*/
type ListNamespacedCapabilitiesForbidden struct {
}

func (o *ListNamespacedCapabilitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /capabilities/namespaces/{namespace}][%d] listNamespacedCapabilitiesForbidden ", 403)
}

func (o *ListNamespacedCapabilitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListNamespacedCapabilitiesTooManyRequests creates a ListNamespacedCapabilitiesTooManyRequests with default headers values
func NewListNamespacedCapabilitiesTooManyRequests() *ListNamespacedCapabilitiesTooManyRequests {
	return &ListNamespacedCapabilitiesTooManyRequests{}
}

/*ListNamespacedCapabilitiesTooManyRequests handles this case with default header values.

Too many requests
*/
type ListNamespacedCapabilitiesTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *ListNamespacedCapabilitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /capabilities/namespaces/{namespace}][%d] listNamespacedCapabilitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListNamespacedCapabilitiesTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNamespacedCapabilitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNamespacedCapabilitiesDefault creates a ListNamespacedCapabilitiesDefault with default headers values
func NewListNamespacedCapabilitiesDefault(code int) *ListNamespacedCapabilitiesDefault {
	return &ListNamespacedCapabilitiesDefault{
		_statusCode: code,
	}
}

/*ListNamespacedCapabilitiesDefault handles this case with default header values.

Unexpected error
*/
type ListNamespacedCapabilitiesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list namespaced capabilities default response
func (o *ListNamespacedCapabilitiesDefault) Code() int {
	return o._statusCode
}

func (o *ListNamespacedCapabilitiesDefault) Error() string {
	return fmt.Sprintf("[GET /capabilities/namespaces/{namespace}][%d] listNamespacedCapabilities default  %+v", o._statusCode, o.Payload)
}

func (o *ListNamespacedCapabilitiesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNamespacedCapabilitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}