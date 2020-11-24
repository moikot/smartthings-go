// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Device device
//
// swagger:model Device
type Device struct {

	// Device Profile information for the SmartApp. This field will be empty if device type is not ENDPOINT_APP.
	App *AppDeviceDetails `json:"app,omitempty"`

	// Bluetooth device information. This field will be empty if device type is not BLE.
	Ble BleDeviceDetails `json:"ble,omitempty"`

	// Bluetooth device to device information. This field will be empty if device type is not BLE_D2D.
	BleD2D *BleD2DDeviceDetails `json:"bleD2D,omitempty"`

	// Device details for all child devices of the device.
	ChildDevices []*Device `json:"childDevices"`

	// The IDs of all compenents on the device.
	Components []*DeviceComponent `json:"components"`

	// The identifier for the device instance.
	// Required: true
	DeviceID *string `json:"deviceId"`

	// The device manufacturer code.
	DeviceManufacturerCode string `json:"deviceManufacturerCode,omitempty"`

	// Deprecated please look under "dth".
	DeviceNetworkType string `json:"deviceNetworkType,omitempty"`

	// Deprecated please look under "dth".
	DeviceTypeID string `json:"deviceTypeId,omitempty"`

	// Deprecated please look under "dth".
	DeviceTypeName string `json:"deviceTypeName,omitempty"`

	// Device Profile information for DTH. This field will be empty if device type is not DTH.
	Dth *DthDeviceDetails `json:"dth,omitempty"`

	// IR device information. This field will be empty if device type is not IR.
	Ir *IrDeviceDetails `json:"ir,omitempty"`

	// IR_OCF device information. This field will be empty if device type is not IR_OCF.
	IrOcf *IrDeviceDetails `json:"irOcf,omitempty"`

	// The name that a user chooses for the device. This defaults to the same value as name.
	Label string `json:"label,omitempty"`

	// The ID of the Location with which the device is associated.
	LocationID string `json:"locationId,omitempty"`

	// The device manufacturer name.
	// Required: true
	ManufacturerName *string `json:"manufacturerName"`

	// The name that the device integration (Device Handler or SmartApp) defines for the device.
	Name string `json:"name,omitempty"`

	// The identifier for the owner of the device instance.
	OwnerID string `json:"ownerId,omitempty"`

	// An non-unique id that is used to help drive UI information.
	// Required: true
	PresentationID *string `json:"presentationId"`

	// profile
	Profile *DeviceProfileReference `json:"profile,omitempty"`

	// Restriction tier of the device, if any.
	// Required: true
	RestrictionTier *int64 `json:"restrictionTier"`

	// The ID of the Room with which the device is associated. If the device is not associated with any room, then this field will be null.
	RoomID string `json:"roomId,omitempty"`

	// type
	// Required: true
	Type DeviceIntegrationType `json:"type"`

	// Viper device information. This field will be empty if device type is not VIPER.
	Viper *ViperDeviceDetails `json:"viper,omitempty"`
}

// Validate validates this device
func (m *Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBleD2D(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIrOcf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturerName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresentationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictionTier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViper(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Device) validateApp(formats strfmt.Registry) error {

	if swag.IsZero(m.App) { // not required
		return nil
	}

	if m.App != nil {
		if err := m.App.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateBleD2D(formats strfmt.Registry) error {

	if swag.IsZero(m.BleD2D) { // not required
		return nil
	}

	if m.BleD2D != nil {
		if err := m.BleD2D.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bleD2D")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateChildDevices(formats strfmt.Registry) error {

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

func (m *Device) validateComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.Components) { // not required
		return nil
	}

	for i := 0; i < len(m.Components); i++ {
		if swag.IsZero(m.Components[i]) { // not required
			continue
		}

		if m.Components[i] != nil {
			if err := m.Components[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("components" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Device) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateDth(formats strfmt.Registry) error {

	if swag.IsZero(m.Dth) { // not required
		return nil
	}

	if m.Dth != nil {
		if err := m.Dth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dth")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateIr(formats strfmt.Registry) error {

	if swag.IsZero(m.Ir) { // not required
		return nil
	}

	if m.Ir != nil {
		if err := m.Ir.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ir")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateIrOcf(formats strfmt.Registry) error {

	if swag.IsZero(m.IrOcf) { // not required
		return nil
	}

	if m.IrOcf != nil {
		if err := m.IrOcf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("irOcf")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateManufacturerName(formats strfmt.Registry) error {

	if err := validate.Required("manufacturerName", "body", m.ManufacturerName); err != nil {
		return err
	}

	return nil
}

func (m *Device) validatePresentationID(formats strfmt.Registry) error {

	if err := validate.Required("presentationId", "body", m.PresentationID); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	if m.Profile != nil {
		if err := m.Profile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profile")
			}
			return err
		}
	}

	return nil
}

func (m *Device) validateRestrictionTier(formats strfmt.Registry) error {

	if err := validate.Required("restrictionTier", "body", m.RestrictionTier); err != nil {
		return err
	}

	return nil
}

func (m *Device) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *Device) validateViper(formats strfmt.Registry) error {

	if swag.IsZero(m.Viper) { // not required
		return nil
	}

	if m.Viper != nil {
		if err := m.Viper.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viper")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Device) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Device) UnmarshalBinary(b []byte) error {
	var res Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
