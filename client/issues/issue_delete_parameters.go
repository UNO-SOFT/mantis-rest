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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIssueDeleteParams creates a new IssueDeleteParams object
// with the default values initialized.
func NewIssueDeleteParams() *IssueDeleteParams {
	var ()
	return &IssueDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIssueDeleteParamsWithTimeout creates a new IssueDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIssueDeleteParamsWithTimeout(timeout time.Duration) *IssueDeleteParams {
	var ()
	return &IssueDeleteParams{

		timeout: timeout,
	}
}

// NewIssueDeleteParamsWithContext creates a new IssueDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewIssueDeleteParamsWithContext(ctx context.Context) *IssueDeleteParams {
	var ()
	return &IssueDeleteParams{

		Context: ctx,
	}
}

// NewIssueDeleteParamsWithHTTPClient creates a new IssueDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIssueDeleteParamsWithHTTPClient(client *http.Client) *IssueDeleteParams {
	var ()
	return &IssueDeleteParams{
		HTTPClient: client,
	}
}

/*IssueDeleteParams contains all the parameters to send to the API endpoint
for the issue delete operation typically these are written to a http.Request
*/
type IssueDeleteParams struct {

	/*ID
	  The issue id.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the issue delete params
func (o *IssueDeleteParams) WithTimeout(timeout time.Duration) *IssueDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue delete params
func (o *IssueDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue delete params
func (o *IssueDeleteParams) WithContext(ctx context.Context) *IssueDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue delete params
func (o *IssueDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue delete params
func (o *IssueDeleteParams) WithHTTPClient(client *http.Client) *IssueDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue delete params
func (o *IssueDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the issue delete params
func (o *IssueDeleteParams) WithID(id int64) *IssueDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the issue delete params
func (o *IssueDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IssueDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
