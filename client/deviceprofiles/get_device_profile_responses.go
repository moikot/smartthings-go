// Code generated by go-swagger; DO NOT EDIT.

package deviceprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// GetDeviceProfileReader is a Reader for the GetDeviceProfile structure.
type GetDeviceProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeviceProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDeviceProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDeviceProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDeviceProfileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetDeviceProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceProfileOK creates a GetDeviceProfileOK with default headers values
func NewGetDeviceProfileOK() *GetDeviceProfileOK {
	return &GetDeviceProfileOK{}
}

/*GetDeviceProfileOK handles this case with default header values.

A Device Profile
*/
type GetDeviceProfileOK struct {
	/*This header field describes the natural language(s) of the intended audience for the representation. This can have multiple values as per [RFC 7231, section 3.1.3.2](https://tools.ietf.org/html/rfc7231#section-3.1.3.2)
	 */
	ContentLanguage string

	Payload *models.DeviceProfileResponse
}

func (o *GetDeviceProfileOK) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}][%d] getDeviceProfileOK  %+v", 200, o.Payload)
}

func (o *GetDeviceProfileOK) GetPayload() *models.DeviceProfileResponse {
	return o.Payload
}

func (o *GetDeviceProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

	o.Payload = new(models.DeviceProfileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceProfileBadRequest creates a GetDeviceProfileBadRequest with default headers values
func NewGetDeviceProfileBadRequest() *GetDeviceProfileBadRequest {
	return &GetDeviceProfileBadRequest{}
}

/*GetDeviceProfileBadRequest handles this case with default header values.

Bad request
*/
type GetDeviceProfileBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceProfileBadRequest) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}][%d] getDeviceProfileBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeviceProfileBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceProfileUnauthorized creates a GetDeviceProfileUnauthorized with default headers values
func NewGetDeviceProfileUnauthorized() *GetDeviceProfileUnauthorized {
	return &GetDeviceProfileUnauthorized{}
}

/*GetDeviceProfileUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeviceProfileUnauthorized struct {
}

func (o *GetDeviceProfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}][%d] getDeviceProfileUnauthorized ", 401)
}

func (o *GetDeviceProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceProfileForbidden creates a GetDeviceProfileForbidden with default headers values
func NewGetDeviceProfileForbidden() *GetDeviceProfileForbidden {
	return &GetDeviceProfileForbidden{}
}

/*GetDeviceProfileForbidden handles this case with default header values.

Forbidden
*/
type GetDeviceProfileForbidden struct {
}

func (o *GetDeviceProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}][%d] getDeviceProfileForbidden ", 403)
}

func (o *GetDeviceProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeviceProfileTooManyRequests creates a GetDeviceProfileTooManyRequests with default headers values
func NewGetDeviceProfileTooManyRequests() *GetDeviceProfileTooManyRequests {
	return &GetDeviceProfileTooManyRequests{}
}

/*GetDeviceProfileTooManyRequests handles this case with default header values.

Too many requests
*/
type GetDeviceProfileTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetDeviceProfileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}][%d] getDeviceProfileTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDeviceProfileTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceProfileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceProfileDefault creates a GetDeviceProfileDefault with default headers values
func NewGetDeviceProfileDefault(code int) *GetDeviceProfileDefault {
	return &GetDeviceProfileDefault{
		_statusCode: code,
	}
}

/*GetDeviceProfileDefault handles this case with default header values.

Unexpected error
*/
type GetDeviceProfileDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device profile default response
func (o *GetDeviceProfileDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceProfileDefault) Error() string {
	return fmt.Sprintf("[GET /deviceprofiles/{deviceProfileId}][%d] getDeviceProfile default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceProfileDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetDeviceProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
