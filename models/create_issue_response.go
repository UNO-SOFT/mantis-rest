// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateIssueResponse create issue response
//
// swagger:model CreateIssueResponse
type CreateIssueResponse struct {

	// issues
	Issues *Issue `json:"issues,omitempty"`
}

// Validate validates this create issue response
func (m *CreateIssueResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateIssueResponse) validateIssues(formats strfmt.Registry) error {
	if swag.IsZero(m.Issues) { // not required
		return nil
	}

	if m.Issues != nil {
		if err := m.Issues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issues")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issues")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create issue response based on the context it is used
func (m *CreateIssueResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIssues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateIssueResponse) contextValidateIssues(ctx context.Context, formats strfmt.Registry) error {

	if m.Issues != nil {

		if swag.IsZero(m.Issues) { // not required
			return nil
		}

		if err := m.Issues.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issues")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issues")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateIssueResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateIssueResponse) UnmarshalBinary(b []byte) error {
	var res CreateIssueResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
