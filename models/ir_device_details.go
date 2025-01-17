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
)

// IrDeviceDetails ir device details
//
// swagger:model IrDeviceDetails
type IrDeviceDetails struct {

	// list of child devices
	ChildDevices []*IrDeviceDetails `json:"childDevices"`

	// function codes
	FunctionCodes *IrDeviceDetailsFunctionCodes `json:"functionCodes,omitempty"`

	// The id of the ircode
	IrCode string `json:"irCode,omitempty"`

	// Key value pairs stored in the conductor for the device.
	// UI Metadata information
	//
	Metadata interface{} `json:"metadata,omitempty"`

	// The OCF Device Type
	OcfDeviceType string `json:"ocfDeviceType,omitempty"`

	// The id of the Parent device.
	ParentDeviceID string `json:"parentDeviceId,omitempty"`

	// The id of the profile that describes the device components and capabilities.
	ProfileID string `json:"profileId,omitempty"`
}

// Validate validates this ir device details
func (m *IrDeviceDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IrDeviceDetails) validateChildDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.ChildDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.ChildDevices); i++ {
		if swag.IsZero(m.ChildDevices[i]) { // not required
			continue
		}

		if m.ChildDevices[i] != nil {
			if err := m.ChildDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IrDeviceDetails) validateFunctionCodes(formats strfmt.Registry) error {

	if swag.IsZero(m.FunctionCodes) { // not required
		return nil
	}

	if m.FunctionCodes != nil {
		if err := m.FunctionCodes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("functionCodes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IrDeviceDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IrDeviceDetails) UnmarshalBinary(b []byte) error {
	var res IrDeviceDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IrDeviceDetailsFunctionCodes List of IR function codes
//
// swagger:model IrDeviceDetailsFunctionCodes
type IrDeviceDetailsFunctionCodes struct {

	// default
	Default string `json:"default,omitempty"`

	// ir device details function codes additional properties
	IrDeviceDetailsFunctionCodesAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *IrDeviceDetailsFunctionCodes) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// default
		Default string `json:"default,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv IrDeviceDetailsFunctionCodes

	rcv.Default = stage1.Default
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "default")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.IrDeviceDetailsFunctionCodesAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m IrDeviceDetailsFunctionCodes) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// default
		Default string `json:"default,omitempty"`
	}

	stage1.Default = m.Default

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.IrDeviceDetailsFunctionCodesAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.IrDeviceDetailsFunctionCodesAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this ir device details function codes
func (m *IrDeviceDetailsFunctionCodes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IrDeviceDetailsFunctionCodes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IrDeviceDetailsFunctionCodes) UnmarshalBinary(b []byte) error {
	var res IrDeviceDetailsFunctionCodes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
