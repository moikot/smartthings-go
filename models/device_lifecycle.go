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

// DeviceLifecycle The device lifecycle. The lifecycle will be one of:
// * CREATE - Invoked when a device is created.
// * DELETE - Invoked when a device is deleted.
// * UPDATE - Invoked when a device is updated.
// * MOVE_FROM - Invoked when a device is moved from a location.
// * MOVE_TO - Invoked when a device is moved to a location.
//
//
// swagger:model DeviceLifecycle
type DeviceLifecycle string

const (

	// DeviceLifecycleCREATE captures enum value "CREATE"
	DeviceLifecycleCREATE DeviceLifecycle = "CREATE"

	// DeviceLifecycleDELETE captures enum value "DELETE"
	DeviceLifecycleDELETE DeviceLifecycle = "DELETE"

	// DeviceLifecycleUPDATE captures enum value "UPDATE"
	DeviceLifecycleUPDATE DeviceLifecycle = "UPDATE"

	// DeviceLifecycleMOVEFROM captures enum value "MOVE_FROM"
	DeviceLifecycleMOVEFROM DeviceLifecycle = "MOVE_FROM"

	// DeviceLifecycleMOVETO captures enum value "MOVE_TO"
	DeviceLifecycleMOVETO DeviceLifecycle = "MOVE_TO"
)

// for schema
var deviceLifecycleEnum []interface{}

func init() {
	var res []DeviceLifecycle
	if err := json.Unmarshal([]byte(`["CREATE","DELETE","UPDATE","MOVE_FROM","MOVE_TO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deviceLifecycleEnum = append(deviceLifecycleEnum, v)
	}
}

func (m DeviceLifecycle) validateDeviceLifecycleEnum(path, location string, value DeviceLifecycle) error {
	if err := validate.EnumCase(path, location, value, deviceLifecycleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this device lifecycle
func (m DeviceLifecycle) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDeviceLifecycleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
