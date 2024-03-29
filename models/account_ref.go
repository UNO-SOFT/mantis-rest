// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountRef account ref
//
// swagger:model AccountRef
type AccountRef struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// realname
	Realname string `json:"realname,omitempty"`
}

// Validate validates this account ref
func (m *AccountRef) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account ref based on context it is used
func (m *AccountRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRef) UnmarshalBinary(b []byte) error {
	var res AccountRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
