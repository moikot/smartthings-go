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

// DetailViewForCapability detail view for capability
//
// swagger:model detailViewForCapability
type DetailViewForCapability []*DetailViewForCapabilityItems0

// Validate validates this detail view for capability
func (m DetailViewForCapability) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// DetailViewForCapabilityItems0 detail view for capability items0
//
// swagger:model DetailViewForCapabilityItems0
type DetailViewForCapabilityItems0 struct {

	// Specify the type of UI component to use to display this action or state. The corresponding field must also be included. For example, if you specify "switch" here, you must also include the "switch" key and its object definition for this action or state.
	// Required: true
	// Enum: [toggleSwitch standbyPowerSwitch switch slider pushButton playPause playStop list textField numberField stepper state]
	DisplayType *string `json:"displayType"`

	// label
	// Required: true
	Label *string `json:"label"`

	// list
	List *ListForDetailView `json:"list,omitempty"`

	// number field
	NumberField *NumberField `json:"numberField,omitempty"`

	// play pause
	PlayPause *PlayPause `json:"playPause,omitempty"`

	// play stop
	PlayStop *PlayStop `json:"playStop,omitempty"`

	// push button
	PushButton *PushButton `json:"pushButton,omitempty"`

	// slider
	Slider *Slider `json:"slider,omitempty"`

	// standby power switch
	StandbyPowerSwitch *StandbyPowerSwitch `json:"standbyPowerSwitch,omitempty"`

	// state
	State *State `json:"state,omitempty"`

	// stepper
	Stepper *Stepper `json:"stepper,omitempty"`

	// switch
	Switch *Switch `json:"switch,omitempty"`

	// text field
	TextField *TextField `json:"textField,omitempty"`

	// toggle switch
	ToggleSwitch *ToggleSwitch `json:"toggleSwitch,omitempty"`
}

// Validate validates this detail view for capability items0
func (m *DetailViewForCapabilityItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayPause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayStop(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushButton(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandbyPowerSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepper(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextField(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToggleSwitch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var detailViewForCapabilityItems0TypeDisplayTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["toggleSwitch","standbyPowerSwitch","switch","slider","pushButton","playPause","playStop","list","textField","numberField","stepper","state"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		detailViewForCapabilityItems0TypeDisplayTypePropEnum = append(detailViewForCapabilityItems0TypeDisplayTypePropEnum, v)
	}
}

const (

	// DetailViewForCapabilityItems0DisplayTypeToggleSwitch captures enum value "toggleSwitch"
	DetailViewForCapabilityItems0DisplayTypeToggleSwitch string = "toggleSwitch"

	// DetailViewForCapabilityItems0DisplayTypeStandbyPowerSwitch captures enum value "standbyPowerSwitch"
	DetailViewForCapabilityItems0DisplayTypeStandbyPowerSwitch string = "standbyPowerSwitch"

	// DetailViewForCapabilityItems0DisplayTypeSwitch captures enum value "switch"
	DetailViewForCapabilityItems0DisplayTypeSwitch string = "switch"

	// DetailViewForCapabilityItems0DisplayTypeSlider captures enum value "slider"
	DetailViewForCapabilityItems0DisplayTypeSlider string = "slider"

	// DetailViewForCapabilityItems0DisplayTypePushButton captures enum value "pushButton"
	DetailViewForCapabilityItems0DisplayTypePushButton string = "pushButton"

	// DetailViewForCapabilityItems0DisplayTypePlayPause captures enum value "playPause"
	DetailViewForCapabilityItems0DisplayTypePlayPause string = "playPause"

	// DetailViewForCapabilityItems0DisplayTypePlayStop captures enum value "playStop"
	DetailViewForCapabilityItems0DisplayTypePlayStop string = "playStop"

	// DetailViewForCapabilityItems0DisplayTypeList captures enum value "list"
	DetailViewForCapabilityItems0DisplayTypeList string = "list"

	// DetailViewForCapabilityItems0DisplayTypeTextField captures enum value "textField"
	DetailViewForCapabilityItems0DisplayTypeTextField string = "textField"

	// DetailViewForCapabilityItems0DisplayTypeNumberField captures enum value "numberField"
	DetailViewForCapabilityItems0DisplayTypeNumberField string = "numberField"

	// DetailViewForCapabilityItems0DisplayTypeStepper captures enum value "stepper"
	DetailViewForCapabilityItems0DisplayTypeStepper string = "stepper"

	// DetailViewForCapabilityItems0DisplayTypeState captures enum value "state"
	DetailViewForCapabilityItems0DisplayTypeState string = "state"
)

// prop value enum
func (m *DetailViewForCapabilityItems0) validateDisplayTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, detailViewForCapabilityItems0TypeDisplayTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DetailViewForCapabilityItems0) validateDisplayType(formats strfmt.Registry) error {

	if err := validate.Required("displayType", "body", m.DisplayType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDisplayTypeEnum("displayType", "body", *m.DisplayType); err != nil {
		return err
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validateList(formats strfmt.Registry) error {

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

func (m *DetailViewForCapabilityItems0) validateNumberField(formats strfmt.Registry) error {

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

func (m *DetailViewForCapabilityItems0) validatePlayPause(formats strfmt.Registry) error {

	if swag.IsZero(m.PlayPause) { // not required
		return nil
	}

	if m.PlayPause != nil {
		if err := m.PlayPause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("playPause")
			}
			return err
		}
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validatePlayStop(formats strfmt.Registry) error {

	if swag.IsZero(m.PlayStop) { // not required
		return nil
	}

	if m.PlayStop != nil {
		if err := m.PlayStop.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("playStop")
			}
			return err
		}
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validatePushButton(formats strfmt.Registry) error {

	if swag.IsZero(m.PushButton) { // not required
		return nil
	}

	if m.PushButton != nil {
		if err := m.PushButton.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pushButton")
			}
			return err
		}
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validateSlider(formats strfmt.Registry) error {

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

func (m *DetailViewForCapabilityItems0) validateStandbyPowerSwitch(formats strfmt.Registry) error {

	if swag.IsZero(m.StandbyPowerSwitch) { // not required
		return nil
	}

	if m.StandbyPowerSwitch != nil {
		if err := m.StandbyPowerSwitch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standbyPowerSwitch")
			}
			return err
		}
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validateStepper(formats strfmt.Registry) error {

	if swag.IsZero(m.Stepper) { // not required
		return nil
	}

	if m.Stepper != nil {
		if err := m.Stepper.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stepper")
			}
			return err
		}
	}

	return nil
}

func (m *DetailViewForCapabilityItems0) validateSwitch(formats strfmt.Registry) error {

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

func (m *DetailViewForCapabilityItems0) validateTextField(formats strfmt.Registry) error {

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

func (m *DetailViewForCapabilityItems0) validateToggleSwitch(formats strfmt.Registry) error {

	if swag.IsZero(m.ToggleSwitch) { // not required
		return nil
	}

	if m.ToggleSwitch != nil {
		if err := m.ToggleSwitch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toggleSwitch")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailViewForCapabilityItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailViewForCapabilityItems0) UnmarshalBinary(b []byte) error {
	var res DetailViewForCapabilityItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}