// Code generated by go-swagger; DO NOT EDIT.

package scenes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// ExecuteSceneReader is a Reader for the ExecuteScene structure.
type ExecuteSceneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteSceneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteSceneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExecuteSceneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExecuteSceneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewExecuteSceneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExecuteSceneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExecuteSceneOK creates a ExecuteSceneOK with default headers values
func NewExecuteSceneOK() *ExecuteSceneOK {
	return &ExecuteSceneOK{}
}

/*ExecuteSceneOK handles this case with default header values.

The Scene has been executed
*/
type ExecuteSceneOK struct {
	Payload *models.StandardSuccessResponse
}

func (o *ExecuteSceneOK) Error() string {
	return fmt.Sprintf("[POST /scenes/{sceneId}/execute][%d] executeSceneOK  %+v", 200, o.Payload)
}

func (o *ExecuteSceneOK) GetPayload() *models.StandardSuccessResponse {
	return o.Payload
}

func (o *ExecuteSceneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardSuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteSceneUnauthorized creates a ExecuteSceneUnauthorized with default headers values
func NewExecuteSceneUnauthorized() *ExecuteSceneUnauthorized {
	return &ExecuteSceneUnauthorized{}
}

/*ExecuteSceneUnauthorized handles this case with default header values.

Unauthorized
*/
type ExecuteSceneUnauthorized struct {
}

func (o *ExecuteSceneUnauthorized) Error() string {
	return fmt.Sprintf("[POST /scenes/{sceneId}/execute][%d] executeSceneUnauthorized ", 401)
}

func (o *ExecuteSceneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecuteSceneForbidden creates a ExecuteSceneForbidden with default headers values
func NewExecuteSceneForbidden() *ExecuteSceneForbidden {
	return &ExecuteSceneForbidden{}
}

/*ExecuteSceneForbidden handles this case with default header values.

Forbidden
*/
type ExecuteSceneForbidden struct {
}

func (o *ExecuteSceneForbidden) Error() string {
	return fmt.Sprintf("[POST /scenes/{sceneId}/execute][%d] executeSceneForbidden ", 403)
}

func (o *ExecuteSceneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecuteSceneTooManyRequests creates a ExecuteSceneTooManyRequests with default headers values
func NewExecuteSceneTooManyRequests() *ExecuteSceneTooManyRequests {
	return &ExecuteSceneTooManyRequests{}
}

/*ExecuteSceneTooManyRequests handles this case with default header values.

Too many requests
*/
type ExecuteSceneTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *ExecuteSceneTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /scenes/{sceneId}/execute][%d] executeSceneTooManyRequests  %+v", 429, o.Payload)
}

func (o *ExecuteSceneTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecuteSceneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteSceneDefault creates a ExecuteSceneDefault with default headers values
func NewExecuteSceneDefault(code int) *ExecuteSceneDefault {
	return &ExecuteSceneDefault{
		_statusCode: code,
	}
}

/*ExecuteSceneDefault handles this case with default header values.

Unexpected error
*/
type ExecuteSceneDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the execute scene default response
func (o *ExecuteSceneDefault) Code() int {
	return o._statusCode
}

func (o *ExecuteSceneDefault) Error() string {
	return fmt.Sprintf("[POST /scenes/{sceneId}/execute][%d] executeScene default  %+v", o._statusCode, o.Payload)
}

func (o *ExecuteSceneDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ExecuteSceneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
