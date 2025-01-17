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

// GetInstallationReader is a Reader for the GetInstallation structure.
type GetInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInstallationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInstallationOK creates a GetInstallationOK with default headers values
func NewGetInstallationOK() *GetInstallationOK {
	return &GetInstallationOK{}
}

/*GetInstallationOK handles this case with default header values.

An installed app.
*/
type GetInstallationOK struct {
	Payload *models.InstalledApp
}

func (o *GetInstallationOK) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}][%d] getInstallationOK  %+v", 200, o.Payload)
}

func (o *GetInstallationOK) GetPayload() *models.InstalledApp {
	return o.Payload
}

func (o *GetInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstalledApp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationBadRequest creates a GetInstallationBadRequest with default headers values
func NewGetInstallationBadRequest() *GetInstallationBadRequest {
	return &GetInstallationBadRequest{}
}

/*GetInstallationBadRequest handles this case with default header values.

Bad request
*/
type GetInstallationBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetInstallationBadRequest) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}][%d] getInstallationBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallationBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInstallationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationUnauthorized creates a GetInstallationUnauthorized with default headers values
func NewGetInstallationUnauthorized() *GetInstallationUnauthorized {
	return &GetInstallationUnauthorized{}
}

/*GetInstallationUnauthorized handles this case with default header values.

Unauthorized
*/
type GetInstallationUnauthorized struct {
}

func (o *GetInstallationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}][%d] getInstallationUnauthorized ", 401)
}

func (o *GetInstallationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstallationForbidden creates a GetInstallationForbidden with default headers values
func NewGetInstallationForbidden() *GetInstallationForbidden {
	return &GetInstallationForbidden{}
}

/*GetInstallationForbidden handles this case with default header values.

Forbidden
*/
type GetInstallationForbidden struct {
}

func (o *GetInstallationForbidden) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}][%d] getInstallationForbidden ", 403)
}

func (o *GetInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInstallationTooManyRequests creates a GetInstallationTooManyRequests with default headers values
func NewGetInstallationTooManyRequests() *GetInstallationTooManyRequests {
	return &GetInstallationTooManyRequests{}
}

/*GetInstallationTooManyRequests handles this case with default header values.

Too many requests
*/
type GetInstallationTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetInstallationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}][%d] getInstallationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetInstallationTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInstallationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallationDefault creates a GetInstallationDefault with default headers values
func NewGetInstallationDefault(code int) *GetInstallationDefault {
	return &GetInstallationDefault{
		_statusCode: code,
	}
}

/*GetInstallationDefault handles this case with default header values.

Unexpected error
*/
type GetInstallationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get installation default response
func (o *GetInstallationDefault) Code() int {
	return o._statusCode
}

func (o *GetInstallationDefault) Error() string {
	return fmt.Sprintf("[GET /installedapps/{installedAppId}][%d] getInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *GetInstallationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
