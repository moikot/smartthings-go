// Code generated by go-swagger; DO NOT EDIT.

package schedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// CreateScheduleReader is a Reader for the CreateSchedule structure.
type CreateScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateScheduleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateScheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateScheduleOK creates a CreateScheduleOK with default headers values
func NewCreateScheduleOK() *CreateScheduleOK {
	return &CreateScheduleOK{}
}

/*CreateScheduleOK handles this case with default header values.

The created schedule.
*/
type CreateScheduleOK struct {
	Payload *models.Schedule
}

func (o *CreateScheduleOK) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/schedules][%d] createScheduleOK  %+v", 200, o.Payload)
}

func (o *CreateScheduleOK) GetPayload() *models.Schedule {
	return o.Payload
}

func (o *CreateScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Schedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduleBadRequest creates a CreateScheduleBadRequest with default headers values
func NewCreateScheduleBadRequest() *CreateScheduleBadRequest {
	return &CreateScheduleBadRequest{}
}

/*CreateScheduleBadRequest handles this case with default header values.

Bad request
*/
type CreateScheduleBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/schedules][%d] createScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *CreateScheduleBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduleUnauthorized creates a CreateScheduleUnauthorized with default headers values
func NewCreateScheduleUnauthorized() *CreateScheduleUnauthorized {
	return &CreateScheduleUnauthorized{}
}

/*CreateScheduleUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateScheduleUnauthorized struct {
}

func (o *CreateScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/schedules][%d] createScheduleUnauthorized ", 401)
}

func (o *CreateScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateScheduleForbidden creates a CreateScheduleForbidden with default headers values
func NewCreateScheduleForbidden() *CreateScheduleForbidden {
	return &CreateScheduleForbidden{}
}

/*CreateScheduleForbidden handles this case with default header values.

Forbidden
*/
type CreateScheduleForbidden struct {
}

func (o *CreateScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/schedules][%d] createScheduleForbidden ", 403)
}

func (o *CreateScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateScheduleUnprocessableEntity creates a CreateScheduleUnprocessableEntity with default headers values
func NewCreateScheduleUnprocessableEntity() *CreateScheduleUnprocessableEntity {
	return &CreateScheduleUnprocessableEntity{}
}

/*CreateScheduleUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type CreateScheduleUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *CreateScheduleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/schedules][%d] createScheduleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateScheduleUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateScheduleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduleTooManyRequests creates a CreateScheduleTooManyRequests with default headers values
func NewCreateScheduleTooManyRequests() *CreateScheduleTooManyRequests {
	return &CreateScheduleTooManyRequests{}
}

/*CreateScheduleTooManyRequests handles this case with default header values.

Too many requests
*/
type CreateScheduleTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CreateScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/schedules][%d] createScheduleTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateScheduleTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateScheduleDefault creates a CreateScheduleDefault with default headers values
func NewCreateScheduleDefault(code int) *CreateScheduleDefault {
	return &CreateScheduleDefault{
		_statusCode: code,
	}
}

/*CreateScheduleDefault handles this case with default header values.

Unexpected error
*/
type CreateScheduleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create schedule default response
func (o *CreateScheduleDefault) Code() int {
	return o._statusCode
}

func (o *CreateScheduleDefault) Error() string {
	return fmt.Sprintf("[POST /installedapps/{installedAppId}/schedules][%d] createSchedule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateScheduleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateScheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
