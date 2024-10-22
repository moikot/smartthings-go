// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// ListAppsReader is a Reader for the ListApps structure.
type ListAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAppsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAppsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListAppsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAppsOK creates a ListAppsOK with default headers values
func NewListAppsOK() *ListAppsOK {
	return &ListAppsOK{}
}

/*ListAppsOK handles this case with default header values.

A paginated list of apps.
*/
type ListAppsOK struct {
	Payload *models.PagedApps
}

func (o *ListAppsOK) Error() string {
	return fmt.Sprintf("[GET /apps][%d] listAppsOK  %+v", 200, o.Payload)
}

func (o *ListAppsOK) GetPayload() *models.PagedApps {
	return o.Payload
}

func (o *ListAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagedApps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAppsUnauthorized creates a ListAppsUnauthorized with default headers values
func NewListAppsUnauthorized() *ListAppsUnauthorized {
	return &ListAppsUnauthorized{}
}

/*ListAppsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAppsUnauthorized struct {
}

func (o *ListAppsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps][%d] listAppsUnauthorized ", 401)
}

func (o *ListAppsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAppsForbidden creates a ListAppsForbidden with default headers values
func NewListAppsForbidden() *ListAppsForbidden {
	return &ListAppsForbidden{}
}

/*ListAppsForbidden handles this case with default header values.

Forbidden
*/
type ListAppsForbidden struct {
}

func (o *ListAppsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps][%d] listAppsForbidden ", 403)
}

func (o *ListAppsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAppsTooManyRequests creates a ListAppsTooManyRequests with default headers values
func NewListAppsTooManyRequests() *ListAppsTooManyRequests {
	return &ListAppsTooManyRequests{}
}

/*ListAppsTooManyRequests handles this case with default header values.

Too many requests
*/
type ListAppsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *ListAppsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps][%d] listAppsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListAppsTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAppsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAppsDefault creates a ListAppsDefault with default headers values
func NewListAppsDefault(code int) *ListAppsDefault {
	return &ListAppsDefault{
		_statusCode: code,
	}
}

/*ListAppsDefault handles this case with default header values.

Unexpected error
*/
type ListAppsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list apps default response
func (o *ListAppsDefault) Code() int {
	return o._statusCode
}

func (o *ListAppsDefault) Error() string {
	return fmt.Sprintf("[GET /apps][%d] listApps default  %+v", o._statusCode, o.Payload)
}

func (o *ListAppsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
