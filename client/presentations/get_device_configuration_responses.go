// Code generated by go-swagger; DO NOT EDIT.

package presentations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// GetDeviceConfigurationReader is a Reader for the GetDeviceConfiguration structure.
type GetDeviceConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeviceConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeviceConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeviceConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDeviceConfigurationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceConfigurationOK creates a GetDeviceConfigurationOK with default headers values
func NewGetDeviceConfigurationOK() *GetDeviceConfigurationOK {
	return &GetDeviceConfigurationOK{}
}

/*GetDeviceConfigurationOK handles this case with default header values.

The device configuration.
*/
type GetDeviceConfigurationOK struct {
	Payload *models.PublicDeviceConfiguration
}

func (o *GetDeviceConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /presentation/deviceconfig][%d] getDeviceConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetDeviceConfigurationOK) GetPayload() *models.PublicDeviceConfiguration {
	return o.Payload
}

func (o *GetDeviceConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicDeviceConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceConfigurationBadRequest creates a GetDeviceConfigurationBadRequest with default headers values
func NewGetDeviceConfigurationBadRequest() *GetDeviceConfigurationBadRequest {
	return &GetDeviceConfigurationBadRequest{}
}

/*GetDeviceConfigurationBadRequest handles this case with default header values.

Bad request
*/
type GetDeviceConfigurationBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /presentation/deviceconfig][%d] getDeviceConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeviceConfigurationBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceConfigurationUnauthorized creates a GetDeviceConfigurationUnauthorized with default headers values
func NewGetDeviceConfigurationUnauthorized() *GetDeviceConfigurationUnauthorized {
	return &GetDeviceConfigurationUnauthorized{}
}

/*GetDeviceConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeviceConfigurationUnauthorized struct {
}

func (o *GetDeviceConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /presentation/deviceconfig][%d] getDeviceConfigurationUnauthorized ", 401)
}

func (o *GetDeviceConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceConfigurationForbidden creates a GetDeviceConfigurationForbidden with default headers values
func NewGetDeviceConfigurationForbidden() *GetDeviceConfigurationForbidden {
	return &GetDeviceConfigurationForbidden{}
}

/*GetDeviceConfigurationForbidden handles this case with default header values.

Forbidden
*/
type GetDeviceConfigurationForbidden struct {
}

func (o *GetDeviceConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /presentation/deviceconfig][%d] getDeviceConfigurationForbidden ", 403)
}

func (o *GetDeviceConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceConfigurationTooManyRequests creates a GetDeviceConfigurationTooManyRequests with default headers values
func NewGetDeviceConfigurationTooManyRequests() *GetDeviceConfigurationTooManyRequests {
	return &GetDeviceConfigurationTooManyRequests{}
}

/*GetDeviceConfigurationTooManyRequests handles this case with default header values.

Too many requests
*/
type GetDeviceConfigurationTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceConfigurationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /presentation/deviceconfig][%d] getDeviceConfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeviceConfigurationTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceConfigurationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceConfigurationDefault creates a GetDeviceConfigurationDefault with default headers values
func NewGetDeviceConfigurationDefault(code int) *GetDeviceConfigurationDefault {
	return &GetDeviceConfigurationDefault{
		_statusCode: code,
	}
}

/*GetDeviceConfigurationDefault handles this case with default header values.

Unexpected error
*/
type GetDeviceConfigurationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device configuration default response
func (o *GetDeviceConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceConfigurationDefault) Error() string {
	return fmt.Sprintf("[GET /presentation/deviceconfig][%d] getDeviceConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceConfigurationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
