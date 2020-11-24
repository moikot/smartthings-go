// Code generated by go-swagger; DO NOT EDIT.

package rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/moikot/smartthings-go/models"
)

// DeleteRuleReader is a Reader for the DeleteRule structure.
type DeleteRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRuleOK creates a DeleteRuleOK with default headers values
func NewDeleteRuleOK() *DeleteRuleOK {
	return &DeleteRuleOK{}
}

/*DeleteRuleOK handles this case with default header values.

Successfully deleted
*/
type DeleteRuleOK struct {
	Payload *models.Rule
}

func (o *DeleteRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /rules/{ruleId}][%d] deleteRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteRuleOK) GetPayload() *models.Rule {
	return o.Payload
}

func (o *DeleteRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRuleUnauthorized creates a DeleteRuleUnauthorized with default headers values
func NewDeleteRuleUnauthorized() *DeleteRuleUnauthorized {
	return &DeleteRuleUnauthorized{}
}

/*DeleteRuleUnauthorized handles this case with default header values.

Not authenticated
*/
type DeleteRuleUnauthorized struct {
}

func (o *DeleteRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /rules/{ruleId}][%d] deleteRuleUnauthorized ", 401)
}

func (o *DeleteRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRuleForbidden creates a DeleteRuleForbidden with default headers values
func NewDeleteRuleForbidden() *DeleteRuleForbidden {
	return &DeleteRuleForbidden{}
}

/*DeleteRuleForbidden handles this case with default header values.

Not authorized or not found
*/
type DeleteRuleForbidden struct {
}

func (o *DeleteRuleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /rules/{ruleId}][%d] deleteRuleForbidden ", 403)
}

func (o *DeleteRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRuleDefault creates a DeleteRuleDefault with default headers values
func NewDeleteRuleDefault(code int) *DeleteRuleDefault {
	return &DeleteRuleDefault{
		_statusCode: code,
	}
}

/*DeleteRuleDefault handles this case with default header values.

Unexpected error
*/
type DeleteRuleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete rule default response
func (o *DeleteRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /rules/{ruleId}][%d] deleteRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
