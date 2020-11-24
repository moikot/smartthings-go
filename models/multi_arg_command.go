// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MultiArgCommand To specify multiple arguments command and displayType of each argument.
//
// swagger:model multiArgCommand
type MultiArgCommand struct {

	// arguments
	// Required: true
	Arguments []*MultiArgCommandArgumentsItems0 `json:"arguments"`

	// Command name to trigger
	// Required: true
	Command *string `json:"command"`
}

// Validate validates this multi arg command
func (m *MultiArgCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultiArgCommand) validateArguments(formats strfmt.Registry) error {

	if err := validate.Required("arguments", "body", m.Arguments); err != nil {
		return err
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MultiArgCommand) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiArgCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiArgCommand) UnmarshalBinary(b []byte) error {
	var res MultiArgCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MultiArgCommandArgumentsItems0 multi arg command arguments items0
//
// swagger:model MultiArgCommandArgumentsItems0
type MultiArgCommandArgumentsItems0 struct {

	// Specify the type of UI component to use to display this action or state. The corresponding field must also be included. For example, if you specify "switch" here, you must also include the "switch" key and its object definition for this action or state.
	// Required: true
	// Enum: [switch slider list textField numberField]
	DisplayType *string `json:"displayType"`

	// list
	List *ListForArgument `json:"list,omitempty"`

	// number field
	NumberField *NumberFieldForArgument `json:"numberField,omitempty"`

	// slider
	Slider *SliderForArgument `json:"slider,omitempty"`

	// switch
	Switch *SwitchForArgument `json:"switch,omitempty"`

	// text field
	TextField *TextFieldForArgument `json:"textField,omitempty"`
}

// Validate validates this multi arg command arguments items0
func (m *MultiArgCommandArgumentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var multiArgCommandArgumentsItems0TypeDisplayTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["switch","slider","list","textField","numberField"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		multiArgCommandArgumentsItems0TypeDisplayTypePropEnum = append(multiArgCommandArgumentsItems0TypeDisplayTypePropEnum, v)
	}
}

const (

	// MultiArgCommandArgumentsItems0DisplayTypeSwitch captures enum value "switch"
	MultiArgCommandArgumentsItems0DisplayTypeSwitch string = "switch"

	// MultiArgCommandArgumentsItems0DisplayTypeSlider captures enum value "slider"
	MultiArgCommandArgumentsItems0DisplayTypeSlider string = "slider"

	// MultiArgCommandArgumentsItems0DisplayTypeList captures enum value "list"
	MultiArgCommandArgumentsItems0DisplayTypeList string = "list"

	// MultiArgCommandArgumentsItems0DisplayTypeTextField captures enum value "textField"
	MultiArgCommandArgumentsItems0DisplayTypeTextField string = "textField"

	// MultiArgCommandArgumentsItems0DisplayTypeNumberField captures enum value "numberField"
	MultiArgCommandArgumentsItems0DisplayTypeNumberField string = "numberField"
)

// prop value enum
func (m *MultiArgCommandArgumentsItems0) validateDisplayTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, multiArgCommandArgumentsItems0TypeDisplayTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MultiArgCommandArgumentsItems0) validateDisplayType(formats strfmt.Registry) error {

	if err := validate.Required("displayType", "body", m.DisplayType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDisplayTypeEnum("displayType", "body", *m.DisplayType); err != nil {
		return err
	}

	return nil
}

func (m *MultiArgCommandArgumentsItems0) validateList(formats strfmt.Registry) error {

	if swag.IsZero(m.List) { // not required
		return nil
	}

	if m.List != nil {
		if err := m.List.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("list")
			}
			return err
		}
	}

	return nil
}

func (m *MultiArgCommandArgumentsItems0) validateNumberField(formats strfmt.Registry) error {

	if swag.IsZero(m.NumberField) { // not required
		return nil
	}

	if m.NumberField != nil {
		if err := m.NumberField.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("numberField")
			}
			return err
		}
	}

	return nil
}

func (m *MultiArgCommandArgumentsItems0) validateSlider(formats strfmt.Registry) error {

	if swag.IsZero(m.Slider) { // not required
		return nil
	}

	if m.Slider != nil {
		if err := m.Slider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slider")
			}
			return err
		}
	}

	return nil
}

func (m *MultiArgCommandArgumentsItems0) validateSwitch(formats strfmt.Registry) error {

	if swag.IsZero(m.Switch) { // not required
		return nil
	}

	if m.Switch != nil {
		if err := m.Switch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("switch")
			}
			return err
		}
	}

	return nil
}

func (m *MultiArgCommandArgumentsItems0) validateTextField(formats strfmt.Registry) error {

	if swag.IsZero(m.TextField) { // not required
		return nil
	}

	if m.TextField != nil {
		if err := m.TextField.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("textField")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultiArgCommandArgumentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultiArgCommandArgumentsItems0) UnmarshalBinary(b []byte) error {
	var res MultiArgCommandArgumentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}