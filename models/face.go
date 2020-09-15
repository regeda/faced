// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Face face
//
// swagger:model Face
type Face struct {

	// bounds
	Bounds *Rect `json:"bounds,omitempty"`

	// left eye
	LeftEye *Point `json:"left_eye,omitempty"`

	// mouth
	Mouth *Point `json:"mouth,omitempty"`

	// right eye
	RightEye *Point `json:"right_eye,omitempty"`
}

// Validate validates this face
func (m *Face) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBounds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLeftEye(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMouth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRightEye(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Face) validateBounds(formats strfmt.Registry) error {

	if swag.IsZero(m.Bounds) { // not required
		return nil
	}

	if m.Bounds != nil {
		if err := m.Bounds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bounds")
			}
			return err
		}
	}

	return nil
}

func (m *Face) validateLeftEye(formats strfmt.Registry) error {

	if swag.IsZero(m.LeftEye) { // not required
		return nil
	}

	if m.LeftEye != nil {
		if err := m.LeftEye.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("left_eye")
			}
			return err
		}
	}

	return nil
}

func (m *Face) validateMouth(formats strfmt.Registry) error {

	if swag.IsZero(m.Mouth) { // not required
		return nil
	}

	if m.Mouth != nil {
		if err := m.Mouth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mouth")
			}
			return err
		}
	}

	return nil
}

func (m *Face) validateRightEye(formats strfmt.Registry) error {

	if swag.IsZero(m.RightEye) { // not required
		return nil
	}

	if m.RightEye != nil {
		if err := m.RightEye.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("right_eye")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Face) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Face) UnmarshalBinary(b []byte) error {
	var res Face
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}