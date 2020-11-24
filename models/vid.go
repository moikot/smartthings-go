// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// Vid A unique identifier for the presentation of a device. This can be a model number on legacy device integrations, but also may be a system generated UUID based on a device's structure and display configuration.
//
// swagger:model vid
type Vid string

// Validate validates this vid
func (m Vid) Validate(formats strfmt.Registry) error {
	return nil
}
