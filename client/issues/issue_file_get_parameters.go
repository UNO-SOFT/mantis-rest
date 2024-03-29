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

// NewIssueFileGetParams creates a new IssueFileGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssueFileGetParams() *IssueFileGetParams {
	return &IssueFileGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssueFileGetParamsWithTimeout creates a new IssueFileGetParams object
// with the ability to set a timeout on a request.
func NewIssueFileGetParamsWithTimeout(timeout time.Duration) *IssueFileGetParams {
	return &IssueFileGetParams{
		timeout: timeout,
	}
}

// NewIssueFileGetParamsWithContext creates a new IssueFileGetParams object
// with the ability to set a context for a request.
func NewIssueFileGetParamsWithContext(ctx context.Context) *IssueFileGetParams {
	return &IssueFileGetParams{
		Context: ctx,
	}
}

// NewIssueFileGetParamsWithHTTPClient creates a new IssueFileGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssueFileGetParamsWithHTTPClient(client *http.Client) *IssueFileGetParams {
	return &IssueFileGetParams{
		HTTPClient: client,
	}
}

/*
IssueFileGetParams contains all the parameters to send to the API endpoint

	for the issue file get operation.

	Typically these are written to a http.Request.
*/
type IssueFileGetParams struct {

	/* FileID.

	   The file id.

	   Format: int64
	*/
	FileID int64

	/* ID.

	   The issue id.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the issue file get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueFileGetParams) WithDefaults() *IssueFileGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issue file get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssueFileGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issue file get params
func (o *IssueFileGetParams) WithTimeout(timeout time.Duration) *IssueFileGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue file get params
func (o *IssueFileGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue file get params
func (o *IssueFileGetParams) WithContext(ctx context.Context) *IssueFileGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue file get params
func (o *IssueFileGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue file get params
func (o *IssueFileGetParams) WithHTTPClient(client *http.Client) *IssueFileGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue file get params
func (o *IssueFileGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileID adds the fileID to the issue file get params
func (o *IssueFileGetParams) WithFileID(fileID int64) *IssueFileGetParams {
	o.SetFileID(fileID)
	return o
}

// SetFileID adds the fileId to the issue file get params
func (o *IssueFileGetParams) SetFileID(fileID int64) {
	o.FileID = fileID
}

// WithID adds the id to the issue file get params
func (o *IssueFileGetParams) WithID(id int64) *IssueFileGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue file get params
func (o *IssueFileGetParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IssueFileGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param file_id
	qrFileID := o.FileID
	qFileID := swag.FormatInt64(qrFileID)
	if qFileID != "" {

		if err := r.SetQueryParam("file_id", qFileID); err != nil {
			return err
		}
	}

	// query param id
	qrID := o.ID
	qID := swag.FormatInt64(qrID)
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
