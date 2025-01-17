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

// UpdateDeviceReader is a Reader for the UpdateDevice structure.
type UpdateDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateDeviceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateDeviceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateDeviceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDeviceOK creates a UpdateDeviceOK with default headers values
func NewUpdateDeviceOK() *UpdateDeviceOK {
	return &UpdateDeviceOK{}
}

/*UpdateDeviceOK handles this case with default header values.

Updated Device.
*/
type UpdateDeviceOK struct {
	Payload *models.Device
}

func (o *UpdateDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] updateDeviceOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceOK) GetPayload() *models.Device {
	return o.Payload
}

func (o *UpdateDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Device)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceBadRequest creates a UpdateDeviceBadRequest with default headers values
func NewUpdateDeviceBadRequest() *UpdateDeviceBadRequest {
	return &UpdateDeviceBadRequest{}
}

/*UpdateDeviceBadRequest handles this case with default header values.

Bad request
*/
type UpdateDeviceBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *UpdateDeviceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] updateDeviceBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDeviceBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceUnauthorized creates a UpdateDeviceUnauthorized with default headers values
func NewUpdateDeviceUnauthorized() *UpdateDeviceUnauthorized {
	return &UpdateDeviceUnauthorized{}
}

/*UpdateDeviceUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateDeviceUnauthorized struct {
}

func (o *UpdateDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] updateDeviceUnauthorized ", 401)
}

func (o *UpdateDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDeviceForbidden creates a UpdateDeviceForbidden with default headers values
func NewUpdateDeviceForbidden() *UpdateDeviceForbidden {
	return &UpdateDeviceForbidden{}
}

/*UpdateDeviceForbidden handles this case with default header values.

Forbidden
*/
type UpdateDeviceForbidden struct {
}

func (o *UpdateDeviceForbidden) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] updateDeviceForbidden ", 403)
}

func (o *UpdateDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDeviceUnprocessableEntity creates a UpdateDeviceUnprocessableEntity with default headers values
func NewUpdateDeviceUnprocessableEntity() *UpdateDeviceUnprocessableEntity {
	return &UpdateDeviceUnprocessableEntity{}
}

/*UpdateDeviceUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type UpdateDeviceUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *UpdateDeviceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] updateDeviceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateDeviceUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceTooManyRequests creates a UpdateDeviceTooManyRequests with default headers values
func NewUpdateDeviceTooManyRequests() *UpdateDeviceTooManyRequests {
	return &UpdateDeviceTooManyRequests{}
}

/*UpdateDeviceTooManyRequests handles this case with default header values.

Too many requests
*/
type UpdateDeviceTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *UpdateDeviceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] updateDeviceTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateDeviceTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDeviceDefault creates a UpdateDeviceDefault with default headers values
func NewUpdateDeviceDefault(code int) *UpdateDeviceDefault {
	return &UpdateDeviceDefault{
		_statusCode: code,
	}
}

/*UpdateDeviceDefault handles this case with default header values.

Unexpected error
*/
type UpdateDeviceDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device default response
func (o *UpdateDeviceDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDeviceDefault) Error() string {
	return fmt.Sprintf("[PUT /devices/{deviceId}][%d] updateDevice default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDeviceDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateDeviceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
