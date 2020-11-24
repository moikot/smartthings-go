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

// GetAppSettingsReader is a Reader for the GetAppSettings structure.
type GetAppSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAppSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAppSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAppSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAppSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAppSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetAppSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAppSettingsOK creates a GetAppSettingsOK with default headers values
func NewGetAppSettingsOK() *GetAppSettingsOK {
	return &GetAppSettingsOK{}
}

/*GetAppSettingsOK handles this case with default header values.

An app settings model.
*/
type GetAppSettingsOK struct {
	Payload *models.GetAppSettingsResponse
}

func (o *GetAppSettingsOK) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/settings][%d] getAppSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAppSettingsOK) GetPayload() *models.GetAppSettingsResponse {
	return o.Payload
}

func (o *GetAppSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAppSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSettingsBadRequest creates a GetAppSettingsBadRequest with default headers values
func NewGetAppSettingsBadRequest() *GetAppSettingsBadRequest {
	return &GetAppSettingsBadRequest{}
}

/*GetAppSettingsBadRequest handles this case with default header values.

Bad request
*/
type GetAppSettingsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetAppSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/settings][%d] getAppSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAppSettingsBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAppSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSettingsUnauthorized creates a GetAppSettingsUnauthorized with default headers values
func NewGetAppSettingsUnauthorized() *GetAppSettingsUnauthorized {
	return &GetAppSettingsUnauthorized{}
}

/*GetAppSettingsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAppSettingsUnauthorized struct {
}

func (o *GetAppSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/settings][%d] getAppSettingsUnauthorized ", 401)
}

func (o *GetAppSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppSettingsForbidden creates a GetAppSettingsForbidden with default headers values
func NewGetAppSettingsForbidden() *GetAppSettingsForbidden {
	return &GetAppSettingsForbidden{}
}

/*GetAppSettingsForbidden handles this case with default header values.

Forbidden
*/
type GetAppSettingsForbidden struct {
}

func (o *GetAppSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/settings][%d] getAppSettingsForbidden ", 403)
}

func (o *GetAppSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppSettingsTooManyRequests creates a GetAppSettingsTooManyRequests with default headers values
func NewGetAppSettingsTooManyRequests() *GetAppSettingsTooManyRequests {
	return &GetAppSettingsTooManyRequests{}
}

/*GetAppSettingsTooManyRequests handles this case with default header values.

Too many requests
*/
type GetAppSettingsTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *GetAppSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/settings][%d] getAppSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAppSettingsTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAppSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppSettingsDefault creates a GetAppSettingsDefault with default headers values
func NewGetAppSettingsDefault(code int) *GetAppSettingsDefault {
	return &GetAppSettingsDefault{
		_statusCode: code,
	}
}

/*GetAppSettingsDefault handles this case with default header values.

Unexpected error
*/
type GetAppSettingsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get app settings default response
func (o *GetAppSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetAppSettingsDefault) Error() string {
	return fmt.Sprintf("[GET /apps/{appNameOrId}/settings][%d] getAppSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetAppSettingsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetAppSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}