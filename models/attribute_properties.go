// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AttributeProperties Data type and value properties used to describe the attribute.
//
// swagger:model AttributeProperties
type AttributeProperties struct {

	// data
	Data *AttributePropertiesData `json:"data,omitempty"`

	// unit
	Unit *AttributePropertiesUnit `json:"unit,omitempty"`

	// value
	// Required: true
	Value *AttributePropertiesValue `json:"value"`
}

// Validate validates this attribute properties
func (m *AttributeProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AttributeProperties) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *AttributeProperties) validateUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

func (m *AttributeProperties) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttributeProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributeProperties) UnmarshalBinary(b []byte) error {
	var res AttributeProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AttributePropertiesData Special case details may be desired to describe the attribute that may not be used the same as a value. Some examples are a timeout or a lock code.
//
// swagger:model AttributePropertiesData
type AttributePropertiesData struct {

	// Indicates if properties not explicitly defined by the schema are allowed. Default is false.
	AdditionalProperties *bool `json:"additionalProperties,omitempty"`

	// Objects can be described in JSON schema for the `data` that should be stored on this attribute.
	Properties interface{} `json:"properties,omitempty"`

	// An array of the properties the `data` schema requires.
	Required []string `json:"required"`

	// The data type for the `data` schema is object.
	// Required: true
	// Enum: [object]
	Type *string `json:"type"`
}

// Validate validates this attribute properties data
func (m *AttributePropertiesData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var attributePropertiesDataTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["object"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attributePropertiesDataTypeTypePropEnum = append(attributePropertiesDataTypeTypePropEnum, v)
	}
}

const (

	// AttributePropertiesDataTypeObject captures enum value "object"
	AttributePropertiesDataTypeObject string = "object"
)

// prop value enum
func (m *AttributePropertiesData) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attributePropertiesDataTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttributePropertiesData) validateType(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("data"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttributePropertiesData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributePropertiesData) UnmarshalBinary(b []byte) error {
	var res AttributePropertiesData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AttributePropertiesUnit The unit which describes the value of the attribute.
//
// swagger:model AttributePropertiesUnit
type AttributePropertiesUnit struct {

	// The default unit to be used for the attribute value.
	Default string `json:"default,omitempty"`

	// Array of all possible units for the attribute value.
	Enum []string `json:"enum"`

	// Data type for the unit. This is defined by capability schema as a `string`.
	// Enum: [string]
	Type *string `json:"type,omitempty"`
}

// Validate validates this attribute properties unit
func (m *AttributePropertiesUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var attributePropertiesUnitTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attributePropertiesUnitTypeTypePropEnum = append(attributePropertiesUnitTypeTypePropEnum, v)
	}
}

const (

	// AttributePropertiesUnitTypeString captures enum value "string"
	AttributePropertiesUnitTypeString string = "string"
)

// prop value enum
func (m *AttributePropertiesUnit) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attributePropertiesUnitTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttributePropertiesUnit) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("unit"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttributePropertiesUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributePropertiesUnit) UnmarshalBinary(b []byte) error {
	var res AttributePropertiesUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AttributePropertiesValue Properties that describe the value of an attribute.
//
// swagger:model AttributePropertiesValue
type AttributePropertiesValue struct {

	// Array of possible values or a `minimum:` `maximum:` can be defined for `integer` type.
	Enum []string `json:"enum"`

	// Data type of the attribute's value.
	Type string `json:"type,omitempty"`

	// attribute properties value
	AttributePropertiesValue map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *AttributePropertiesValue) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// Array of possible values or a `minimum:` `maximum:` can be defined for `integer` type.
		Enum []string `json:"enum"`

		// Data type of the attribute's value.
		Type string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv AttributePropertiesValue

	rcv.Enum = stage1.Enum
	rcv.Type = stage1.Type
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "enum")
	delete(stage2, "type")
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
		m.AttributePropertiesValue = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m AttributePropertiesValue) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// Array of possible values or a `minimum:` `maximum:` can be defined for `integer` type.
		Enum []string `json:"enum"`

		// Data type of the attribute's value.
		Type string `json:"type,omitempty"`
	}

	stage1.Enum = m.Enum
	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.AttributePropertiesValue) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.AttributePropertiesValue)
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

// Validate validates this attribute properties value
func (m *AttributePropertiesValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttributePropertiesValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributePropertiesValue) UnmarshalBinary(b []byte) error {
	var res AttributePropertiesValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
