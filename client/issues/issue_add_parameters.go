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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/UNO-SOFT/mantis-rest/models"
)

// NewIssueAddParams creates a new IssueAddParams object
// with the default values initialized.
func NewIssueAddParams() *IssueAddParams {
	var ()
	return &IssueAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIssueAddParamsWithTimeout creates a new IssueAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIssueAddParamsWithTimeout(timeout time.Duration) *IssueAddParams {
	var ()
	return &IssueAddParams{

		timeout: timeout,
	}
}

// NewIssueAddParamsWithContext creates a new IssueAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewIssueAddParamsWithContext(ctx context.Context) *IssueAddParams {
	var ()
	return &IssueAddParams{

		Context: ctx,
	}
}

// NewIssueAddParamsWithHTTPClient creates a new IssueAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIssueAddParamsWithHTTPClient(client *http.Client) *IssueAddParams {
	var ()
	return &IssueAddParams{
		HTTPClient: client,
	}
}

/*IssueAddParams contains all the parameters to send to the API endpoint
for the issue add operation typically these are written to a http.Request
*/
type IssueAddParams struct {

	/*Body
	  The issue to add.

	*/
	Body *models.Issue

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the issue add params
func (o *IssueAddParams) WithTimeout(timeout time.Duration) *IssueAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issue add params
func (o *IssueAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issue add params
func (o *IssueAddParams) WithContext(ctx context.Context) *IssueAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issue add params
func (o *IssueAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issue add params
func (o *IssueAddParams) WithHTTPClient(client *http.Client) *IssueAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issue add params
func (o *IssueAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issue add params
func (o *IssueAddParams) WithBody(body *models.Issue) *IssueAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issue add params
func (o *IssueAddParams) SetBody(body *models.Issue) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IssueAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
