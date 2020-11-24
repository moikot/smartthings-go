// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceLifecycleEvent A device lifecycle event.
//
// swagger:model DeviceLifecycleEvent
type DeviceLifecycleEvent struct {

	// create
	Create *DeviceLifecycleCreate `json:"create,omitempty"`

	// delete
	Delete DeviceLifecycleDelete `json:"delete,omitempty"`

	// The id of the device.
	DeviceID string `json:"deviceId,omitempty"`

	// The name of the device
	DeviceName string `json:"deviceName,omitempty"`

	// The id of the event.
	EventID string `json:"eventId,omitempty"`

	// lifecycle
	Lifecycle DeviceLifecycle `json:"lifecycle,omitempty"`

	// The id of the location in which the event was triggered.
	LocationID string `json:"locationId,omitempty"`

	// move from
	MoveFrom *DeviceLifecycleMove `json:"moveFrom,omitempty"`

	// move to
	MoveTo *DeviceLifecycleMove `json:"moveTo,omitempty"`

	// ID for what owns the device lifecyle event. Works in tandem with `ownerType` as a composite identifier.
	OwnerID string `json:"ownerId,omitempty"`

	// owner type
	OwnerType EventOwnerType `json:"ownerType,omitempty"`

	// The principal that made the change
	Principal string `json:"principal,omitempty"`

	// update
	Update *DeviceLifecycleUpdate `json:"update,omitempty"`
}

// Validate validates this device lifecycle event
func (m *DeviceLifecycleEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoveTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceLifecycleEvent) validateCreate(formats strfmt.Registry) error {

	if swag.IsZero(m.Create) { // not required
		return nil
	}

	if m.Create != nil {
		if err := m.Create.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("create")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceLifecycleEvent) validateLifecycle(formats strfmt.Registry) error {

	if swag.IsZero(m.Lifecycle) { // not required
		return nil
	}

	if err := m.Lifecycle.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lifecycle")
		}
		return err
	}

	return nil
}

func (m *DeviceLifecycleEvent) validateMoveFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveFrom) { // not required
		return nil
	}

	if m.MoveFrom != nil {
		if err := m.MoveFrom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveFrom")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceLifecycleEvent) validateMoveTo(formats strfmt.Registry) error {

	if swag.IsZero(m.MoveTo) { // not required
		return nil
	}

	if m.MoveTo != nil {
		if err := m.MoveTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("moveTo")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceLifecycleEvent) validateOwnerType(formats strfmt.Registry) error {

	if swag.IsZero(m.OwnerType) { // not required
		return nil
	}

	if err := m.OwnerType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ownerType")
		}
		return err
	}

	return nil
}

func (m *DeviceLifecycleEvent) validateUpdate(formats strfmt.Registry) error {

	if swag.IsZero(m.Update) { // not required
		return nil
	}

	if m.Update != nil {
		if err := m.Update.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceLifecycleEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceLifecycleEvent) UnmarshalBinary(b []byte) error {
	var res DeviceLifecycleEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}