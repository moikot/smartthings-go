// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AirQualityData Air quality data
//
//
// swagger:model AirQualityData
type AirQualityData struct {

	// air quality index
	AirQualityIndex int64 `json:"airQualityIndex,omitempty"`

	// co amount in ugm3
	CoAmountInUgm3 float64 `json:"coAmountInUgm3,omitempty"`

	// co index
	CoIndex int64 `json:"coIndex,omitempty"`

	// no2 amount in ugm3
	No2AmountInUgm3 float64 `json:"no2AmountInUgm3,omitempty"`

	// no2 index
	No2Index int64 `json:"no2Index,omitempty"`

	// o3 amount in ugm3
	O3AmountInUgm3 float64 `json:"o3AmountInUgm3,omitempty"`

	// o3 index
	O3Index int64 `json:"o3Index,omitempty"`

	// pm10 amount in ugm3
	Pm10AmountInUgm3 float64 `json:"pm10AmountInUgm3,omitempty"`

	// pm10 index
	Pm10Index int64 `json:"pm10Index,omitempty"`

	// pm25 amount in ugm3
	Pm25AmountInUgm3 float64 `json:"pm25AmountInUgm3,omitempty"`

	// pm25 index
	Pm25Index int64 `json:"pm25Index,omitempty"`

	// so2 amount in ugm3
	So2AmountInUgm3 float64 `json:"so2AmountInUgm3,omitempty"`

	// so2 index
	So2Index int64 `json:"so2Index,omitempty"`
}

// Validate validates this air quality data
func (m *AirQualityData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AirQualityData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AirQualityData) UnmarshalBinary(b []byte) error {
	var res AirQualityData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
