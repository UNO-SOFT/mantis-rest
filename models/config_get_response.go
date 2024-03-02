// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigGetResponse config get response
//
// swagger:model ConfigGetResponse
type ConfigGetResponse struct {

	// configs
	Configs []*ConfigOption `json:"configs"`
}

// Validate validates this config get response
func (m *ConfigGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigGetResponse) validateConfigs(formats strfmt.Registry) error {
	if swag.IsZero(m.Configs) { // not required
		return nil
	}

	for i := 0; i < len(m.Configs); i++ {
		if swag.IsZero(m.Configs[i]) { // not required
			continue
		}

		if m.Configs[i] != nil {
			if err := m.Configs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config get response based on the context it is used
func (m *ConfigGetResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigGetResponse) contextValidateConfigs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Configs); i++ {

		if m.Configs[i] != nil {

			if swag.IsZero(m.Configs[i]) { // not required
				return nil
			}

			if err := m.Configs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigGetResponse) UnmarshalBinary(b []byte) error {
	var res ConfigGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
