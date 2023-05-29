// Code generated by go-swagger; DO NOT EDIT.

package issues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewIssueNoteDeleteParams creates a new IssueNoteDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueNoteDeleteParams() *IssueNoteDeleteParams {
	return &IssueNoteDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueNoteDeleteParamsWithTimeout creates a new IssueNoteDeleteParams object
// with the ability to set a timeout on a request.
func NewIssueNoteDeleteParamsWithTimeout(timeout time.Duration) *IssueNoteDeleteParams {
	return &IssueNoteDeleteParams{
		timeout: timeout,
	}
}

// NewIssueNoteDeleteParamsWithContext creates a new IssueNoteDeleteParams object
// with the ability to set a context for a request.
func NewIssueNoteDeleteParamsWithContext(ctx context.Context) *IssueNoteDeleteParams {
	return &IssueNoteDeleteParams{
		Context: ctx,
	}
}

// NewIssueNoteDeleteParamsWithHTTPClient creates a new IssueNoteDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueNoteDeleteParamsWithHTTPClient(client *http.Client) *IssueNoteDeleteParams {
	return &IssueNoteDeleteParams{
		HTTPClient: client,
	}
}

/*
IssueNoteDeleteParams contains all the parameters to send to the API endpoint

	for the issue note delete operation.

	Typically these are written to a http.Request.
*/
type IssueNoteDeleteParams struct {

	/* ID.

	   The issue id.

	   Format: int64
	*/
	ID int64

	/* NoteID.

	   The note id.

	   Format: int64
	*/
	NoteID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the issue note delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueNoteDeleteParams) WithDefaults() *IssueNoteDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue note delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueNoteDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue note delete params
func (o *IssueNoteDeleteParams) WithTimeout(timeout time.Duration) *IssueNoteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue note delete params
func (o *IssueNoteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue note delete params
func (o *IssueNoteDeleteParams) WithContext(ctx context.Context) *IssueNoteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue note delete params
func (o *IssueNoteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue note delete params
func (o *IssueNoteDeleteParams) WithHTTPClient(client *http.Client) *IssueNoteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue note delete params
func (o *IssueNoteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the issue note delete params
func (o *IssueNoteDeleteParams) WithID(id int64) *IssueNoteDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue note delete params
func (o *IssueNoteDeleteParams) SetID(id int64) {
	o.ID = id
}

// WithNoteID adds the noteID to the issue note delete params
func (o *IssueNoteDeleteParams) WithNoteID(noteID int64) *IssueNoteDeleteParams {
	o.SetNoteID(noteID)
	return o
}

// SetNoteID adds the noteId to the issue note delete params
func (o *IssueNoteDeleteParams) SetNoteID(noteID int64) {
	o.NoteID = noteID
}

// WriteToRequest writes these params to a swagger request
func (o *IssueNoteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := swag.FormatInt64(qrID)
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	// query param note_id
	qrNoteID := o.NoteID
	qNoteID := swag.FormatInt64(qrNoteID)
	if qNoteID != "" {

		if err := r.SetQueryParam("note_id", qNoteID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
