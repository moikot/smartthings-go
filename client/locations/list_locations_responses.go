// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// ListLocationsReader is a Reader for the ListLocations structure.
type ListLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListLocationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewListLocationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListLocationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListLocationsOK creates a ListLocationsOK with default headers values
func NewListLocationsOK() *ListLocationsOK {
	return &ListLocationsOK{}
}

/*ListLocationsOK handles this case with default header values.

An array of Locations
*/
type ListLocationsOK struct {
	Payload *models.PagedLocations
}

func (o *ListLocationsOK) Error() string {
	return fmt.Sprintf("[GET /locations][%d] listLocationsOK  %+v", 200, o.Payload)
}

func (o *ListLocationsOK) GetPayload() *models.PagedLocations {
	return o.Payload
}

func (o *ListLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PagedLocations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocationsUnauthorized creates a ListLocationsUnauthorized with default headers values
func NewListLocationsUnauthorized() *ListLocationsUnauthorized {
	return &ListLocationsUnauthorized{}
}

/*ListLocationsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListLocationsUnauthorized struct {
}

func (o *ListLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /locations][%d] listLocationsUnauthorized ", 401)
}

func (o *ListLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLocationsForbidden creates a ListLocationsForbidden with default headers values
func NewListLocationsForbidden() *ListLocationsForbidden {
	return &ListLocationsForbidden{}
}

/*ListLocationsForbidden handles this case with default header values.

Forbidden
*/
type ListLocationsForbidden struct {
}

func (o *ListLocationsForbidden) Error() string {
	return fmt.Sprintf("[GET /locations][%d] listLocationsForbidden ", 403)
}

func (o *ListLocationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListLocationsTooManyRequests creates a ListLocationsTooManyRequests with default headers values
func NewListLocationsTooManyRequests() *ListLocationsTooManyRequests {
	return &ListLocationsTooManyRequests{}
}

/*ListLocationsTooManyRequests handles this case with default header values.

Too many requests
*/
type ListLocationsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *ListLocationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /locations][%d] listLocationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *ListLocationsTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListLocationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocationsDefault creates a ListLocationsDefault with default headers values
func NewListLocationsDefault(code int) *ListLocationsDefault {
	return &ListLocationsDefault{
		_statusCode: code,
	}
}

/*ListLocationsDefault handles this case with default header values.

Unexpected error
*/
type ListLocationsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list locations default response
func (o *ListLocationsDefault) Code() int {
	return o._statusCode
}

func (o *ListLocationsDefault) Error() string {
	return fmt.Sprintf("[GET /locations][%d] listLocations default  %+v", o._statusCode, o.Payload)
}

func (o *ListLocationsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListLocationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
