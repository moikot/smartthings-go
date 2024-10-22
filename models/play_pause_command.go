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

// PlayPauseCommand Display type for a play/pause button.
//
// swagger:model playPauseCommand
type PlayPauseCommand struct {

	// command
	// Required: true
	Command *PlayPauseCommandCommand `json:"command"`
}

// Validate validates this play pause command
func (m *PlayPauseCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayPauseCommand) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	if m.Command != nil {
		if err := m.Command.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("command")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayPauseCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayPauseCommand) UnmarshalBinary(b []byte) error {
	var res PlayPauseCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PlayPauseCommandCommand To specify separate commands with no arguments for on and off, use the “play” and “pause” fields respectively. To specify a single command, use “name” for the command and the “play” and “pause” fields for the arguments.
//
// swagger:model PlayPauseCommandCommand
type PlayPauseCommandCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// pause
	// Required: true
	Pause *string `json:"pause"`

	// play
	// Required: true
	Play *string `json:"play"`
}

// Validate validates this play pause command command
func (m *PlayPauseCommandCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlayPauseCommandCommand) validatePause(formats strfmt.Registry) error {

	if err := validate.Required("command"+"."+"pause", "body", m.Pause); err != nil {
		return err
	}

	return nil
}

func (m *PlayPauseCommandCommand) validatePlay(formats strfmt.Registry) error {

	if err := validate.Required("command"+"."+"play", "body", m.Play); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlayPauseCommandCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlayPauseCommandCommand) UnmarshalBinary(b []byte) error {
	var res PlayPauseCommandCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
