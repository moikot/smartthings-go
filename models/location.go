// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Location location
//
// swagger:model Location
type Location struct {

	// Additional information about the location that allows SmartThings to further define your location.
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`

	// Not currently in use.
	BackgroundImage string `json:"backgroundImage,omitempty"`

	// An ISO Alpha-3 country code.  (i.e. GBR, USA)
	CountryCode string `json:"countryCode,omitempty"`

	// A geographical latitude.
	Latitude float64 `json:"latitude,omitempty"`

	// An IETF BCP 47 language tag representing the chosen locale for this location.
	Locale string `json:"locale,omitempty"`

	// The ID of the location.
	// Format: uuid
	LocationID strfmt.UUID `json:"locationId,omitempty"`

	// A geographical longitude.
	Longitude float64 `json:"longitude,omitempty"`

	// A name given for the location (eg. Home)
	Name string `json:"name,omitempty"`

	// parent
	Parent *LocationParent `json:"parent,omitempty"`

	// The radius in meters around latitude and longitude which defines this location.
	RegionRadius int32 `json:"regionRadius,omitempty"`

	// The desired temperature scale used within location. Value can be F or C.
	TemperatureScale string `json:"temperatureScale,omitempty"`

	// An ID matching the Java Time Zone ID of the location. Automatically generated if latitude and longitude have been
	// provided.
	//
	TimeZoneID string `json:"timeZoneId,omitempty"`
}

// Validate validates this location
func (m *Location) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Location) validateLocationID(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationID) { // not required
		return nil
	}

	if err := validate.FormatOf("locationId", "body", "uuid", m.LocationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Location) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Location) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Location) UnmarshalBinary(b []byte) error {
	var res Location
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
