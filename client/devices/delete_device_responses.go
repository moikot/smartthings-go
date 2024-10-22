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

// DeleteDeviceReader is a Reader for the DeleteDevice structure.
type DeleteDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteDeviceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDeviceOK creates a DeleteDeviceOK with default headers values
func NewDeleteDeviceOK() *DeleteDeviceOK {
	return &DeleteDeviceOK{}
}

/*DeleteDeviceOK handles this case with default header values.

Device deleted.
*/
type DeleteDeviceOK struct {
	Payload models.DeleteDeviceCommandsResponse
}

func (o *DeleteDeviceOK) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}][%d] deleteDeviceOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceOK) GetPayload() models.DeleteDeviceCommandsResponse {
	return o.Payload
}

func (o *DeleteDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceBadRequest creates a DeleteDeviceBadRequest with default headers values
func NewDeleteDeviceBadRequest() *DeleteDeviceBadRequest {
	return &DeleteDeviceBadRequest{}
}

/*DeleteDeviceBadRequest handles this case with default header values.

Bad request
*/
type DeleteDeviceBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DeleteDeviceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}][%d] deleteDeviceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDeviceBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceUnauthorized creates a DeleteDeviceUnauthorized with default headers values
func NewDeleteDeviceUnauthorized() *DeleteDeviceUnauthorized {
	return &DeleteDeviceUnauthorized{}
}

/*DeleteDeviceUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteDeviceUnauthorized struct {
}

func (o *DeleteDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}][%d] deleteDeviceUnauthorized ", 401)
}

func (o *DeleteDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceForbidden creates a DeleteDeviceForbidden with default headers values
func NewDeleteDeviceForbidden() *DeleteDeviceForbidden {
	return &DeleteDeviceForbidden{}
}

/*DeleteDeviceForbidden handles this case with default header values.

Forbidden
*/
type DeleteDeviceForbidden struct {
}

func (o *DeleteDeviceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}][%d] deleteDeviceForbidden ", 403)
}

func (o *DeleteDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDeviceTooManyRequests creates a DeleteDeviceTooManyRequests with default headers values
func NewDeleteDeviceTooManyRequests() *DeleteDeviceTooManyRequests {
	return &DeleteDeviceTooManyRequests{}
}

/*DeleteDeviceTooManyRequests handles this case with default header values.

Too many requests
*/
type DeleteDeviceTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *DeleteDeviceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}][%d] deleteDeviceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteDeviceTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDeviceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceDefault creates a DeleteDeviceDefault with default headers values
func NewDeleteDeviceDefault(code int) *DeleteDeviceDefault {
	return &DeleteDeviceDefault{
		_statusCode: code,
	}
}

/*DeleteDeviceDefault handles this case with default header values.

Unexpected error
*/
type DeleteDeviceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete device default response
func (o *DeleteDeviceDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDeviceDefault) Error() string {
	return fmt.Sprintf("[DELETE /devices/{deviceId}][%d] deleteDevice default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeviceDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
