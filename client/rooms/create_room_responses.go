// Code generated by go-swagger; DO NOT EDIT.

package rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// CreateRoomReader is a Reader for the CreateRoom structure.
type CreateRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRoomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRoomUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateRoomUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateRoomTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateRoomDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRoomOK creates a CreateRoomOK with default headers values
func NewCreateRoomOK() *CreateRoomOK {
	return &CreateRoomOK{}
}

/*CreateRoomOK handles this case with default header values.

Created successfully.
*/
type CreateRoomOK struct {
	Payload *models.Room
}

func (o *CreateRoomOK) Error() string {
	return fmt.Sprintf("[POST /locations/{locationId}/rooms][%d] createRoomOK  %+v", 200, o.Payload)
}

func (o *CreateRoomOK) GetPayload() *models.Room {
	return o.Payload
}

func (o *CreateRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Room)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoomBadRequest creates a CreateRoomBadRequest with default headers values
func NewCreateRoomBadRequest() *CreateRoomBadRequest {
	return &CreateRoomBadRequest{}
}

/*CreateRoomBadRequest handles this case with default header values.

Bad request
*/
type CreateRoomBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateRoomBadRequest) Error() string {
	return fmt.Sprintf("[POST /locations/{locationId}/rooms][%d] createRoomBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRoomBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRoomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoomUnauthorized creates a CreateRoomUnauthorized with default headers values
func NewCreateRoomUnauthorized() *CreateRoomUnauthorized {
	return &CreateRoomUnauthorized{}
}

/*CreateRoomUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRoomUnauthorized struct {
}

func (o *CreateRoomUnauthorized) Error() string {
	return fmt.Sprintf("[POST /locations/{locationId}/rooms][%d] createRoomUnauthorized ", 401)
}

func (o *CreateRoomUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoomForbidden creates a CreateRoomForbidden with default headers values
func NewCreateRoomForbidden() *CreateRoomForbidden {
	return &CreateRoomForbidden{}
}

/*CreateRoomForbidden handles this case with default header values.

Forbidden
*/
type CreateRoomForbidden struct {
}

func (o *CreateRoomForbidden) Error() string {
	return fmt.Sprintf("[POST /locations/{locationId}/rooms][%d] createRoomForbidden ", 403)
}

func (o *CreateRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRoomUnprocessableEntity creates a CreateRoomUnprocessableEntity with default headers values
func NewCreateRoomUnprocessableEntity() *CreateRoomUnprocessableEntity {
	return &CreateRoomUnprocessableEntity{}
}

/*CreateRoomUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type CreateRoomUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *CreateRoomUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /locations/{locationId}/rooms][%d] createRoomUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateRoomUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRoomUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoomTooManyRequests creates a CreateRoomTooManyRequests with default headers values
func NewCreateRoomTooManyRequests() *CreateRoomTooManyRequests {
	return &CreateRoomTooManyRequests{}
}

/*CreateRoomTooManyRequests handles this case with default header values.

Too many requests
*/
type CreateRoomTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CreateRoomTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /locations/{locationId}/rooms][%d] createRoomTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateRoomTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRoomTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRoomDefault creates a CreateRoomDefault with default headers values
func NewCreateRoomDefault(code int) *CreateRoomDefault {
	return &CreateRoomDefault{
		_statusCode: code,
	}
}

/*CreateRoomDefault handles this case with default header values.

Unexpected error
*/
type CreateRoomDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create room default response
func (o *CreateRoomDefault) Code() int {
	return o._statusCode
}

func (o *CreateRoomDefault) Error() string {
	return fmt.Sprintf("[POST /locations/{locationId}/rooms][%d] createRoom default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRoomDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateRoomDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}