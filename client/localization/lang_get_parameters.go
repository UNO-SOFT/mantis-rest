// Code generated by go-swagger; DO NOT EDIT.

package localization

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

// NewLangGetParams creates a new LangGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLangGetParams() *LangGetParams {
	return &LangGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLangGetParamsWithTimeout creates a new LangGetParams object
// with the ability to set a timeout on a request.
func NewLangGetParamsWithTimeout(timeout time.Duration) *LangGetParams {
	return &LangGetParams{
		timeout: timeout,
	}
}

// NewLangGetParamsWithContext creates a new LangGetParams object
// with the ability to set a context for a request.
func NewLangGetParamsWithContext(ctx context.Context) *LangGetParams {
	return &LangGetParams{
		Context: ctx,
	}
}

// NewLangGetParamsWithHTTPClient creates a new LangGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLangGetParamsWithHTTPClient(client *http.Client) *LangGetParams {
	return &LangGetParams{
		HTTPClient: client,
	}
}

/*
LangGetParams contains all the parameters to send to the API endpoint

	for the lang get operation.

	Typically these are written to a http.Request.
*/
type LangGetParams struct {

	/* String.

	   An array of localized labels given their name string lang/strings_english.txt folder in MantisBT.  The name doesn't include $s_ prefix.
	*/
	String []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lang get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LangGetParams) WithDefaults() *LangGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lang get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LangGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lang get params
func (o *LangGetParams) WithTimeout(timeout time.Duration) *LangGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lang get params
func (o *LangGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lang get params
func (o *LangGetParams) WithContext(ctx context.Context) *LangGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lang get params
func (o *LangGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lang get params
func (o *LangGetParams) WithHTTPClient(client *http.Client) *LangGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lang get params
func (o *LangGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithString adds the string to the lang get params
func (o *LangGetParams) WithString(string []string) *LangGetParams {
	o.SetString(string)
	return o
}

// SetString adds the string to the lang get params
func (o *LangGetParams) SetString(string []string) {
	o.String = string
}

// WriteToRequest writes these params to a swagger request
func (o *LangGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.String != nil {

		// binding items for string
		joinedString := o.bindParamString(reg)

		// query array param string
		if err := r.SetQueryParam("string", joinedString...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLangGet binds the parameter string
func (o *LangGetParams) bindParamString(formats strfmt.Registry) []string {
	stringIR := o.String

	var stringIC []string
	for _, stringIIR := range stringIR { // explode []string

		stringIIV := stringIIR // string as string
		stringIC = append(stringIC, stringIIV)
	}

	// items.CollectionFormat: ""
	stringIS := swag.JoinByFormat(stringIC, "")

	return stringIS
}
