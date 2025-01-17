// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// GetAppReader is a Reader for the GetApp structure.
type GetAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAppTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppOK creates a GetAppOK with default headers values
func NewGetAppOK() *GetAppOK {
	return &GetAppOK{}
}

/*GetAppOK handles this case with default header values.

An app.
*/
type GetAppOK struct {
	Payload *models.App
}

func (o *GetAppOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}][%d] getAppOK  %+v", 200, o.Payload)
}

func (o *GetAppOK) GetPayload() *models.App {
	return o.Payload
}

func (o *GetAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppBadRequest creates a GetAppBadRequest with default headers values
func NewGetAppBadRequest() *GetAppBadRequest {
	return &GetAppBadRequest{}
}

/*GetAppBadRequest handles this case with default header values.

Bad request
*/
type GetAppBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetAppBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}][%d] getAppBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppUnauthorized creates a GetAppUnauthorized with default headers values
func NewGetAppUnauthorized() *GetAppUnauthorized {
	return &GetAppUnauthorized{}
}

/*GetAppUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAppUnauthorized struct {
}

func (o *GetAppUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}][%d] getAppUnauthorized ", 401)
}

func (o *GetAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppForbidden creates a GetAppForbidden with default headers values
func NewGetAppForbidden() *GetAppForbidden {
	return &GetAppForbidden{}
}

/*GetAppForbidden handles this case with default header values.

Forbidden
*/
type GetAppForbidden struct {
}

func (o *GetAppForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}][%d] getAppForbidden ", 403)
}

func (o *GetAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppTooManyRequests creates a GetAppTooManyRequests with default headers values
func NewGetAppTooManyRequests() *GetAppTooManyRequests {
	return &GetAppTooManyRequests{}
}

/*GetAppTooManyRequests handles this case with default header values.

Too many requests
*/
type GetAppTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetAppTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}][%d] getAppTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAppTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAppTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppDefault creates a GetAppDefault with default headers values
func NewGetAppDefault(code int) *GetAppDefault {
	return &GetAppDefault{
		_statusCode: code,
	}
}

/*GetAppDefault handles this case with default header values.

Unexpected error
*/
type GetAppDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get app default response
func (o *GetAppDefault) Code() int {
	return o._statusCode
}

func (o *GetAppDefault) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}][%d] getApp default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
