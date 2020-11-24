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

// AttributeSchema [JSON schema](http://json-schema.org/specification-links.html#draft-4) for the attribute. The API implements JSON schema version 4. For more info regarding JSON schema, please read [Understanding JSON Schema](https://json-schema.org/understanding-json-schema/index.html).
//
//
// swagger:model AttributeSchema
type AttributeSchema struct {

	// additional properties
	// Enum: [false]
	AdditionalProperties bool `json:"additionalProperties,omitempty"`

	// properties
	// Required: true
	Properties *AttributeProperties `json:"properties"`

	// Provide requirement for `value`, `unit`, and `data` fields.
	Required []string `json:"required"`

	// The capability's name, no spaces.
	Title string `json:"title,omitempty"`

	// For schema this will always be object and is the only allowed value.
	// Enum: [object]
	Type string `json:"type,omitempty"`
}

// Validate validates this attribute schema
func (m *AttributeSchema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var attributeSchemaTypeAdditionalPropertiesPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attributeSchemaTypeAdditionalPropertiesPropEnum = append(attributeSchemaTypeAdditionalPropertiesPropEnum, v)
	}
}

// prop value enum
func (m *AttributeSchema) validateAdditionalPropertiesEnum(path, location string, value bool) error {
	if err := validate.EnumCase(path, location, value, attributeSchemaTypeAdditionalPropertiesPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttributeSchema) validateAdditionalProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalProperties) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdditionalPropertiesEnum("additionalProperties", "body", m.AdditionalProperties); err != nil {
		return err
	}

	return nil
}

func (m *AttributeSchema) validateProperties(formats strfmt.Registry) error {

	if err := validate.Required("properties", "body", m.Properties); err != nil {
		return err
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

var attributeSchemaRequiredItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["value","unit","data"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attributeSchemaRequiredItemsEnum = append(attributeSchemaRequiredItemsEnum, v)
	}
}

func (m *AttributeSchema) validateRequiredItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attributeSchemaRequiredItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttributeSchema) validateRequired(formats strfmt.Registry) error {

	if swag.IsZero(m.Required) { // not required
		return nil
	}

	for i := 0; i < len(m.Required); i++ {

		// value enum
		if err := m.validateRequiredItemsEnum("required"+"."+strconv.Itoa(i), "body", m.Required[i]); err != nil {
			return err
		}

	}

	return nil
}

var attributeSchemaTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["object"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		attributeSchemaTypeTypePropEnum = append(attributeSchemaTypeTypePropEnum, v)
	}
}

const (

	// AttributeSchemaTypeObject captures enum value "object"
	AttributeSchemaTypeObject string = "object"
)

// prop value enum
func (m *AttributeSchema) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, attributeSchemaTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AttributeSchema) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AttributeSchema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttributeSchema) UnmarshalBinary(b []byte) error {
	var res AttributeSchema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
