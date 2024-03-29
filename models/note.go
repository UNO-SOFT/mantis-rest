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
	"github.com/go-openapi/validate"
)

// Note note
//
// swagger:model Note
type Note struct {

	// attachments
	Attachments []*FileRef `json:"attachments"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// reporter
	Reporter *AccountRef `json:"reporter,omitempty"`

	// text
	// Required: true
	Text *string `json:"text"`

	// time tracking
	TimeTracking *TimeTrackingRef `json:"time_tracking,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// view state
	ViewState *ViewStateRef `json:"view_state,omitempty"`
}

// Validate validates this note
func (m *Note) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReporter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeTracking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Note) validateAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Note) validateReporter(formats strfmt.Registry) error {
	if swag.IsZero(m.Reporter) { // not required
		return nil
	}

	if m.Reporter != nil {
		if err := m.Reporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reporter")
			}
			return err
		}
	}

	return nil
}

func (m *Note) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *Note) validateTimeTracking(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeTracking) { // not required
		return nil
	}

	if m.TimeTracking != nil {
		if err := m.TimeTracking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_tracking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *Note) validateViewState(formats strfmt.Registry) error {
	if swag.IsZero(m.ViewState) { // not required
		return nil
	}

	if m.ViewState != nil {
		if err := m.ViewState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("view_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("view_state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this note based on the context it is used
func (m *Note) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeTracking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Note) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {

			if swag.IsZero(m.Attachments[i]) { // not required
				return nil
			}

			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attachments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attachments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Note) contextValidateReporter(ctx context.Context, formats strfmt.Registry) error {

	if m.Reporter != nil {

		if swag.IsZero(m.Reporter) { // not required
			return nil
		}

		if err := m.Reporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reporter")
			}
			return err
		}
	}

	return nil
}

func (m *Note) contextValidateTimeTracking(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeTracking != nil {

		if swag.IsZero(m.TimeTracking) { // not required
			return nil
		}

		if err := m.TimeTracking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("time_tracking")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("time_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *Note) contextValidateViewState(ctx context.Context, formats strfmt.Registry) error {

	if m.ViewState != nil {

		if swag.IsZero(m.ViewState) { // not required
			return nil
		}

		if err := m.ViewState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("view_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("view_state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Note) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Note) UnmarshalBinary(b []byte) error {
	var res Note
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
