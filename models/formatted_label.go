// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// FormattedLabel This displays a string. This can be a formatted string with variables.
// Example: `{{attribute.value}} {{attribute.unit}}` (where `attribute` is the name of an attribute in your capability)
//
//
// swagger:model formattedLabel
type FormattedLabel string

// Validate validates this formatted label
func (m FormattedLabel) Validate(formats strfmt.Registry) error {
	return nil
}