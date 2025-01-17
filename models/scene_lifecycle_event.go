// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SceneLifecycleEvent A scene lifecycle event.
//
// swagger:model SceneLifecycleEvent
type SceneLifecycleEvent struct {

	// create
	Create SceneLifecycleCreate `json:"create,omitempty"`

	// delete
	Delete SceneLifecycleDelete `json:"delete,omitempty"`

	// The id of the event.
	EventID string `json:"eventId,omitempty"`

	// lifecycle
	Lifecycle SceneLifecycle `json:"lifecycle,omitempty"`

	// The id of the location in which the event was triggered.
	LocationID string `json:"locationId,omitempty"`

	// The id of the scene.
	SceneID string `json:"sceneId,omitempty"`

	// update
	Update SceneLifecycleUpdate `json:"update,omitempty"`
}

// Validate validates this scene lifecycle event
func (m *SceneLifecycleEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLifecycle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SceneLifecycleEvent) validateLifecycle(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SceneLifecycleEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SceneLifecycleEvent) UnmarshalBinary(b []byte) error {
	var res SceneLifecycleEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
