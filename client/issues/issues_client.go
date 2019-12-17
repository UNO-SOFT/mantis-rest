// Code generated by go-swagger; DO NOT EDIT.

package issues

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new issues API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for issues API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
IssueAdd creates an issue
*/
func (a *Client) IssueAdd(params *IssueAddParams, authInfo runtime.ClientAuthInfoWriter) (*IssueAddCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "issueAdd",
		Method:             "POST",
		PathPattern:        "/issues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IssueAddCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueAdd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueDelete deletes an issue
*/
func (a *Client) IssueDelete(params *IssueDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*IssueDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "issueDelete",
		Method:             "DELETE",
		PathPattern:        "/issues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IssueDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
IssueGet gets issue details
*/
func (a *Client) IssueGet(params *IssueGetParams, authInfo runtime.ClientAuthInfoWriter) (*IssueGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIssueGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "issueGet",
		Method:             "GET",
		PathPattern:        "/issues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IssueGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IssueGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for issueGet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}