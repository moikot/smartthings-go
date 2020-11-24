// Code generated by go-swagger; DO NOT EDIT.

package installedapps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// CreateInstalledAppEventsReader is a Reader for the CreateInstalledAppEvents structure.
type CreateInstalledAppEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInstalledAppEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateInstalledAppEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInstalledAppEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInstalledAppEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInstalledAppEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateInstalledAppEventsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateInstalledAppEventsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateInstalledAppEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateInstalledAppEventsOK creates a CreateInstalledAppEventsOK with default headers values
func NewCreateInstalledAppEventsOK() *CreateInstalledAppEventsOK {
	return &CreateInstalledAppEventsOK{}
}

/*CreateInstalledAppEventsOK handles this case with default header values.

Created events.
*/
type CreateInstalledAppEventsOK struct {
	Payload models.CreateInstalledAppEventsResponse
}

func (o *CreateInstalledAppEventsOK) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/events][%d] createInstalledAppEventsOK  %+v", 200, o.Payload)
}

func (o *CreateInstalledAppEventsOK) GetPayload() models.CreateInstalledAppEventsResponse {
	return o.Payload
}

func (o *CreateInstalledAppEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstalledAppEventsBadRequest creates a CreateInstalledAppEventsBadRequest with default headers values
func NewCreateInstalledAppEventsBadRequest() *CreateInstalledAppEventsBadRequest {
	return &CreateInstalledAppEventsBadRequest{}
}

/*CreateInstalledAppEventsBadRequest handles this case with default header values.

Bad request
*/
type CreateInstalledAppEventsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateInstalledAppEventsBadRequest) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/events][%d] createInstalledAppEventsBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInstalledAppEventsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInstalledAppEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstalledAppEventsUnauthorized creates a CreateInstalledAppEventsUnauthorized with default headers values
func NewCreateInstalledAppEventsUnauthorized() *CreateInstalledAppEventsUnauthorized {
	return &CreateInstalledAppEventsUnauthorized{}
}

/*CreateInstalledAppEventsUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateInstalledAppEventsUnauthorized struct {
}

func (o *CreateInstalledAppEventsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/events][%d] createInstalledAppEventsUnauthorized ", 401)
}

func (o *CreateInstalledAppEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstalledAppEventsForbidden creates a CreateInstalledAppEventsForbidden with default headers values
func NewCreateInstalledAppEventsForbidden() *CreateInstalledAppEventsForbidden {
	return &CreateInstalledAppEventsForbidden{}
}

/*CreateInstalledAppEventsForbidden handles this case with default header values.

Forbidden
*/
type CreateInstalledAppEventsForbidden struct {
}

func (o *CreateInstalledAppEventsForbidden) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/events][%d] createInstalledAppEventsForbidden ", 403)
}

func (o *CreateInstalledAppEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateInstalledAppEventsUnprocessableEntity creates a CreateInstalledAppEventsUnprocessableEntity with default headers values
func NewCreateInstalledAppEventsUnprocessableEntity() *CreateInstalledAppEventsUnprocessableEntity {
	return &CreateInstalledAppEventsUnprocessableEntity{}
}

/*CreateInstalledAppEventsUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type CreateInstalledAppEventsUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *CreateInstalledAppEventsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/events][%d] createInstalledAppEventsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateInstalledAppEventsUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInstalledAppEventsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstalledAppEventsTooManyRequests creates a CreateInstalledAppEventsTooManyRequests with default headers values
func NewCreateInstalledAppEventsTooManyRequests() *CreateInstalledAppEventsTooManyRequests {
	return &CreateInstalledAppEventsTooManyRequests{}
}

/*CreateInstalledAppEventsTooManyRequests handles this case with default header values.

Too many requests
*/
type CreateInstalledAppEventsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CreateInstalledAppEventsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/events][%d] createInstalledAppEventsTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateInstalledAppEventsTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInstalledAppEventsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInstalledAppEventsDefault creates a CreateInstalledAppEventsDefault with default headers values
func NewCreateInstalledAppEventsDefault(code int) *CreateInstalledAppEventsDefault {
	return &CreateInstalledAppEventsDefault{
		_statusCode: code,
	}
}

/*CreateInstalledAppEventsDefault handles this case with default header values.

Unexpected error
*/
type CreateInstalledAppEventsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create installed app events default response
func (o *CreateInstalledAppEventsDefault) Code() int {
	return o._statusCode
}

func (o *CreateInstalledAppEventsDefault) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/events][%d] createInstalledAppEvents default  %+v", o._statusCode, o.Payload)
}

func (o *CreateInstalledAppEventsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateInstalledAppEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
