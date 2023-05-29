// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewConfigGetParams creates a new ConfigGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfigGetParams() *ConfigGetParams {
	return &ConfigGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfigGetParamsWithTimeout creates a new ConfigGetParams object
// with the ability to set a timeout on a request.
func NewConfigGetParamsWithTimeout(timeout time.Duration) *ConfigGetParams {
	return &ConfigGetParams{
		timeout: timeout,
	}
}

// NewConfigGetParamsWithContext creates a new ConfigGetParams object
// with the ability to set a context for a request.
func NewConfigGetParamsWithContext(ctx context.Context) *ConfigGetParams {
	return &ConfigGetParams{
		Context: ctx,
	}
}

// NewConfigGetParamsWithHTTPClient creates a new ConfigGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfigGetParamsWithHTTPClient(client *http.Client) *ConfigGetParams {
	return &ConfigGetParams{
		HTTPClient: client,
	}
}

/*
ConfigGetParams contains all the parameters to send to the API endpoint

	for the config get operation.

	Typically these are written to a http.Request.
*/
type ConfigGetParams struct {

	/* Option.

	   An array of configuration options.
	*/
	Option []string

	/* ProjectID.

	   The project id (default All Projects).

	   Format: int64
	*/
	ProjectID *int64

	/* UserID.

	   The user id (default is logged in user).  This can only be set by users with access level ADMINISTRATOR.

	   Format: int64
	*/
	UserID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the config get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigGetParams) WithDefaults() *ConfigGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the config get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the config get params
func (o *ConfigGetParams) WithTimeout(timeout time.Duration) *ConfigGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config get params
func (o *ConfigGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config get params
func (o *ConfigGetParams) WithContext(ctx context.Context) *ConfigGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config get params
func (o *ConfigGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config get params
func (o *ConfigGetParams) WithHTTPClient(client *http.Client) *ConfigGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config get params
func (o *ConfigGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOption adds the option to the config get params
func (o *ConfigGetParams) WithOption(option []string) *ConfigGetParams {
	o.SetOption(option)
	return o
}

// SetOption adds the option to the config get params
func (o *ConfigGetParams) SetOption(option []string) {
	o.Option = option
}

// WithProjectID adds the projectID to the config get params
func (o *ConfigGetParams) WithProjectID(projectID *int64) *ConfigGetParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the config get params
func (o *ConfigGetParams) SetProjectID(projectID *int64) {
	o.ProjectID = projectID
}

// WithUserID adds the userID to the config get params
func (o *ConfigGetParams) WithUserID(userID *int64) *ConfigGetParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the config get params
func (o *ConfigGetParams) SetUserID(userID *int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Option != nil {

		// binding items for option
		joinedOption := o.bindParamOption(reg)

		// query array param option
		if err := r.SetQueryParam("option", joinedOption...); err != nil {
			return err
		}
	}

	if o.ProjectID != nil {

		// query param project_id
		var qrProjectID int64

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt64(qrProjectID)
		if qProjectID != "" {

			if err := r.SetQueryParam("project_id", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID int64

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := swag.FormatInt64(qrUserID)
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamConfigGet binds the parameter option
func (o *ConfigGetParams) bindParamOption(formats strfmt.Registry) []string {
	optionIR := o.Option

	var optionIC []string
	for _, optionIIR := range optionIR { // explode []string

		optionIIV := optionIIR // string as string
		optionIC = append(optionIC, optionIIV)
	}

	// items.CollectionFormat: ""
	optionIS := swag.JoinByFormat(optionIC, "")

	return optionIS
}
