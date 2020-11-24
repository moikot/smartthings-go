// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DeviceNetworkSecurityLevel For hub-connected devices on protocols with multiple security levels, the security level the device was added with.
// Values ZWAVE_S2_FAILED, ZWAVE_S0_FAILED, ZWAVE_S2_DOWNGRADE, and ZWAVE_S0_DOWNGRADE should be treated as failures
// and prompt the user to remove and re-add the device. Other Z-Wave values are success cases.
//
//
// swagger:model DeviceNetworkSecurityLevel
type DeviceNetworkSecurityLevel string

const (

	// DeviceNetworkSecurityLevelUNKNOWN captures enum value "UNKNOWN"
	DeviceNetworkSecurityLevelUNKNOWN DeviceNetworkSecurityLevel = "UNKNOWN"

	// DeviceNetworkSecurityLevelZWAVELEGACYNONSECURE captures enum value "ZWAVE_LEGACY_NON_SECURE"
	DeviceNetworkSecurityLevelZWAVELEGACYNONSECURE DeviceNetworkSecurityLevel = "ZWAVE_LEGACY_NON_SECURE"

	// DeviceNetworkSecurityLevelZWAVES0LEGACY captures enum value "ZWAVE_S0_LEGACY"
	DeviceNetworkSecurityLevelZWAVES0LEGACY DeviceNetworkSecurityLevel = "ZWAVE_S0_LEGACY"

	// DeviceNetworkSecurityLevelZWAVES0FALLBACK captures enum value "ZWAVE_S0_FALLBACK"
	DeviceNetworkSecurityLevelZWAVES0FALLBACK DeviceNetworkSecurityLevel = "ZWAVE_S0_FALLBACK"

	// DeviceNetworkSecurityLevelZWAVES2UNAUTHENTICATED captures enum value "ZWAVE_S2_UNAUTHENTICATED"
	DeviceNetworkSecurityLevelZWAVES2UNAUTHENTICATED DeviceNetworkSecurityLevel = "ZWAVE_S2_UNAUTHENTICATED"

	// DeviceNetworkSecurityLevelZWAVES2AUTHENTICATED captures enum value "ZWAVE_S2_AUTHENTICATED"
	DeviceNetworkSecurityLevelZWAVES2AUTHENTICATED DeviceNetworkSecurityLevel = "ZWAVE_S2_AUTHENTICATED"

	// DeviceNetworkSecurityLevelZWAVES2ACCESSCONTROL captures enum value "ZWAVE_S2_ACCESS_CONTROL"
	DeviceNetworkSecurityLevelZWAVES2ACCESSCONTROL DeviceNetworkSecurityLevel = "ZWAVE_S2_ACCESS_CONTROL"

	// DeviceNetworkSecurityLevelZWAVES2FAILED captures enum value "ZWAVE_S2_FAILED"
	DeviceNetworkSecurityLevelZWAVES2FAILED DeviceNetworkSecurityLevel = "ZWAVE_S2_FAILED"

	// DeviceNetworkSecurityLevelZWAVES0FAILED captures enum value "ZWAVE_S0_FAILED"
	DeviceNetworkSecurityLevelZWAVES0FAILED DeviceNetworkSecurityLevel = "ZWAVE_S0_FAILED"

	// DeviceNetworkSecurityLevelZWAVES2DOWNGRADE captures enum value "ZWAVE_S2_DOWNGRADE"
	DeviceNetworkSecurityLevelZWAVES2DOWNGRADE DeviceNetworkSecurityLevel = "ZWAVE_S2_DOWNGRADE"

	// DeviceNetworkSecurityLevelZWAVES0DOWNGRADE captures enum value "ZWAVE_S0_DOWNGRADE"
	DeviceNetworkSecurityLevelZWAVES0DOWNGRADE DeviceNetworkSecurityLevel = "ZWAVE_S0_DOWNGRADE"
)

// for schema
var deviceNetworkSecurityLevelEnum []interface{}

func init() {
	var res []DeviceNetworkSecurityLevel
	if err := json.Unmarshal([]byte(`["UNKNOWN","ZWAVE_LEGACY_NON_SECURE","ZWAVE_S0_LEGACY","ZWAVE_S0_FALLBACK","ZWAVE_S2_UNAUTHENTICATED","ZWAVE_S2_AUTHENTICATED","ZWAVE_S2_ACCESS_CONTROL","ZWAVE_S2_FAILED","ZWAVE_S0_FAILED","ZWAVE_S2_DOWNGRADE","ZWAVE_S0_DOWNGRADE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceNetworkSecurityLevelEnum = append(deviceNetworkSecurityLevelEnum, v)
	}
}

func (m DeviceNetworkSecurityLevel) validateDeviceNetworkSecurityLevelEnum(path, location string, value DeviceNetworkSecurityLevel) error {
	if err := validate.EnumCase(path, location, value, deviceNetworkSecurityLevelEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this device network security level
func (m DeviceNetworkSecurityLevel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeviceNetworkSecurityLevelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}