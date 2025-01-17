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

// GenerateAppOauthReader is a Reader for the GenerateAppOauth structure.
type GenerateAppOauthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateAppOauthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateAppOauthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenerateAppOauthBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGenerateAppOauthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGenerateAppOauthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGenerateAppOauthTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGenerateAppOauthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateAppOauthOK creates a GenerateAppOauthOK with default headers values
func NewGenerateAppOauthOK() *GenerateAppOauthOK {
	return &GenerateAppOauthOK{}
}

/*GenerateAppOauthOK handles this case with default header values.

An a response object containing the newly create OAuth Client ID / Secret and relevant details
pertaining to the OAuth client.

*/
type GenerateAppOauthOK struct {
	Payload *models.GenerateAppOAuthResponse
}

func (o *GenerateAppOauthOK) Error() string {
	return fmt.Sprintf("[POST /apps/{appNameOrId}/oauth/generate][%d] generateAppOauthOK  %+v", 200, o.Payload)
}

func (o *GenerateAppOauthOK) GetPayload() *models.GenerateAppOAuthResponse {
	return o.Payload
}

func (o *GenerateAppOauthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenerateAppOAuthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateAppOauthBadRequest creates a GenerateAppOauthBadRequest with default headers values
func NewGenerateAppOauthBadRequest() *GenerateAppOauthBadRequest {
	return &GenerateAppOauthBadRequest{}
}

/*GenerateAppOauthBadRequest handles this case with default header values.

Bad request
*/
type GenerateAppOauthBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GenerateAppOauthBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{appNameOrId}/oauth/generate][%d] generateAppOauthBadRequest  %+v", 400, o.Payload)
}

func (o *GenerateAppOauthBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GenerateAppOauthBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateAppOauthUnauthorized creates a GenerateAppOauthUnauthorized with default headers values
func NewGenerateAppOauthUnauthorized() *GenerateAppOauthUnauthorized {
	return &GenerateAppOauthUnauthorized{}
}

/*GenerateAppOauthUnauthorized handles this case with default header values.

Unauthorized
*/
type GenerateAppOauthUnauthorized struct {
}

func (o *GenerateAppOauthUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{appNameOrId}/oauth/generate][%d] generateAppOauthUnauthorized ", 401)
}

func (o *GenerateAppOauthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateAppOauthForbidden creates a GenerateAppOauthForbidden with default headers values
func NewGenerateAppOauthForbidden() *GenerateAppOauthForbidden {
	return &GenerateAppOauthForbidden{}
}

/*GenerateAppOauthForbidden handles this case with default header values.

Forbidden
*/
type GenerateAppOauthForbidden struct {
}

func (o *GenerateAppOauthForbidden) Error() string {
	return fmt.Sprintf("[POST /apps/{appNameOrId}/oauth/generate][%d] generateAppOauthForbidden ", 403)
}

func (o *GenerateAppOauthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGenerateAppOauthTooManyRequests creates a GenerateAppOauthTooManyRequests with default headers values
func NewGenerateAppOauthTooManyRequests() *GenerateAppOauthTooManyRequests {
	return &GenerateAppOauthTooManyRequests{}
}

/*GenerateAppOauthTooManyRequests handles this case with default header values.

Too many requests
*/
type GenerateAppOauthTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GenerateAppOauthTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /apps/{appNameOrId}/oauth/generate][%d] generateAppOauthTooManyRequests  %+v", 429, o.Payload)
}

func (o *GenerateAppOauthTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GenerateAppOauthTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateAppOauthDefault creates a GenerateAppOauthDefault with default headers values
func NewGenerateAppOauthDefault(code int) *GenerateAppOauthDefault {
	return &GenerateAppOauthDefault{
		_statusCode: code,
	}
}

/*GenerateAppOauthDefault handles this case with default header values.

Unexpected error
*/
type GenerateAppOauthDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the generate app oauth default response
func (o *GenerateAppOauthDefault) Code() int {
	return o._statusCode
}

func (o *GenerateAppOauthDefault) Error() string {
	return fmt.Sprintf("[POST /apps/{appNameOrId}/oauth/generate][%d] generateAppOauth default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateAppOauthDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GenerateAppOauthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
