// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// GetDeviceReader is a Reader for the GetDevice structure.
type GetDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDeviceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceOK creates a GetDeviceOK with default headers values
func NewGetDeviceOK() *GetDeviceOK {
	return &GetDeviceOK{}
}

/*GetDeviceOK handles this case with default header values.

A Device
*/
type GetDeviceOK struct {
	Payload *models.Device
}

func (o *GetDeviceOK) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}][%d] getDeviceOK  %+v", 200, o.Payload)
}

func (o *GetDeviceOK) GetPayload() *models.Device {
	return o.Payload
}

func (o *GetDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Device)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceBadRequest creates a GetDeviceBadRequest with default headers values
func NewGetDeviceBadRequest() *GetDeviceBadRequest {
	return &GetDeviceBadRequest{}
}

/*GetDeviceBadRequest handles this case with default header values.

Bad request
*/
type GetDeviceBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceBadRequest) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}][%d] getDeviceBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeviceBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceUnauthorized creates a GetDeviceUnauthorized with default headers values
func NewGetDeviceUnauthorized() *GetDeviceUnauthorized {
	return &GetDeviceUnauthorized{}
}

/*GetDeviceUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeviceUnauthorized struct {
}

func (o *GetDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}][%d] getDeviceUnauthorized ", 401)
}

func (o *GetDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceForbidden creates a GetDeviceForbidden with default headers values
func NewGetDeviceForbidden() *GetDeviceForbidden {
	return &GetDeviceForbidden{}
}

/*GetDeviceForbidden handles this case with default header values.

Forbidden
*/
type GetDeviceForbidden struct {
}

func (o *GetDeviceForbidden) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}][%d] getDeviceForbidden ", 403)
}

func (o *GetDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceTooManyRequests creates a GetDeviceTooManyRequests with default headers values
func NewGetDeviceTooManyRequests() *GetDeviceTooManyRequests {
	return &GetDeviceTooManyRequests{}
}

/*GetDeviceTooManyRequests handles this case with default header values.

Too many requests
*/
type GetDeviceTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}][%d] getDeviceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeviceTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDefault creates a GetDeviceDefault with default headers values
func NewGetDeviceDefault(code int) *GetDeviceDefault {
	return &GetDeviceDefault{
		_statusCode: code,
	}
}

/*GetDeviceDefault handles this case with default header values.

Unexpected error
*/
type GetDeviceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device default response
func (o *GetDeviceDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDefault) Error() string {
	return fmt.Sprintf("[GET /devices/{deviceId}][%d] getDevice default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
