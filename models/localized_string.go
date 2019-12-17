// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LocalizedString localized string
// swagger:model LocalizedString
type LocalizedString struct {

	// localized
	Localized string `json:"localized,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this localized string
func (m *LocalizedString) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocalizedString) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocalizedString) UnmarshalBinary(b []byte) error {
	var res LocalizedString
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
