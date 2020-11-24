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

// DeleteRoomReader is a Reader for the DeleteRoom structure.
type DeleteRoomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoomBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoomUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoomForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoomTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRoomDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRoomOK creates a DeleteRoomOK with default headers values
func NewDeleteRoomOK() *DeleteRoomOK {
	return &DeleteRoomOK{}
}

/*DeleteRoomOK handles this case with default header values.

An empty object response.
*/
type DeleteRoomOK struct {
	Payload models.DeleteRoomResponse
}

func (o *DeleteRoomOK) Error() string {
	return fmt.Sprintf("[DELETE /locations/{locationId}/rooms/{roomId}][%d] deleteRoomOK  %+v", 200, o.Payload)
}

func (o *DeleteRoomOK) GetPayload() models.DeleteRoomResponse {
	return o.Payload
}

func (o *DeleteRoomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoomBadRequest creates a DeleteRoomBadRequest with default headers values
func NewDeleteRoomBadRequest() *DeleteRoomBadRequest {
	return &DeleteRoomBadRequest{}
}

/*DeleteRoomBadRequest handles this case with default header values.

Bad request
*/
type DeleteRoomBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DeleteRoomBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /locations/{locationId}/rooms/{roomId}][%d] deleteRoomBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoomBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRoomBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoomUnauthorized creates a DeleteRoomUnauthorized with default headers values
func NewDeleteRoomUnauthorized() *DeleteRoomUnauthorized {
	return &DeleteRoomUnauthorized{}
}

/*DeleteRoomUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRoomUnauthorized struct {
}

func (o *DeleteRoomUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /locations/{locationId}/rooms/{roomId}][%d] deleteRoomUnauthorized ", 401)
}

func (o *DeleteRoomUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoomForbidden creates a DeleteRoomForbidden with default headers values
func NewDeleteRoomForbidden() *DeleteRoomForbidden {
	return &DeleteRoomForbidden{}
}

/*DeleteRoomForbidden handles this case with default header values.

Forbidden
*/
type DeleteRoomForbidden struct {
}

func (o *DeleteRoomForbidden) Error() string {
	return fmt.Sprintf("[DELETE /locations/{locationId}/rooms/{roomId}][%d] deleteRoomForbidden ", 403)
}

func (o *DeleteRoomForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoomTooManyRequests creates a DeleteRoomTooManyRequests with default headers values
func NewDeleteRoomTooManyRequests() *DeleteRoomTooManyRequests {
	return &DeleteRoomTooManyRequests{}
}

/*DeleteRoomTooManyRequests handles this case with default header values.

Too many requests
*/
type DeleteRoomTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *DeleteRoomTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /locations/{locationId}/rooms/{roomId}][%d] deleteRoomTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoomTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRoomTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoomDefault creates a DeleteRoomDefault with default headers values
func NewDeleteRoomDefault(code int) *DeleteRoomDefault {
	return &DeleteRoomDefault{
		_statusCode: code,
	}
}

/*DeleteRoomDefault handles this case with default header values.

Unexpected error
*/
type DeleteRoomDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete room default response
func (o *DeleteRoomDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRoomDefault) Error() string {
	return fmt.Sprintf("[DELETE /locations/{locationId}/rooms/{roomId}][%d] deleteRoom default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRoomDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteRoomDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}