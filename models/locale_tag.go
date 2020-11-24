// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
)

// LocaleTag The tag of the locale as defined in [RFC bcp47](http://www.rfc-editor.org/rfc/bcp/bcp47.txt).
//
// swagger:model LocaleTag
type LocaleTag string

// Validate validates this locale tag
func (m LocaleTag) Validate(formats strfmt.Registry) error {
	return nil
}
