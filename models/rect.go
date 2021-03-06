// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Rect rect
//
// swagger:model Rect
type Rect struct {

	// height
	Height int64 `json:"height,omitempty"`

	// width
	Width int64 `json:"width,omitempty"`

	// x
	X int64 `json:"x,omitempty"`

	// y
	Y int64 `json:"y,omitempty"`
}

// Validate validates this rect
func (m *Rect) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Rect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rect) UnmarshalBinary(b []byte) error {
	var res Rect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
